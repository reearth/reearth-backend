package interactor

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	htmlTmpl "html/template"
	"net/mail"
	textTmpl "text/template"

	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/log"
	"github.com/reearth/reearth-backend/pkg/project"
	"github.com/reearth/reearth-backend/pkg/rerror"
	"github.com/reearth/reearth-backend/pkg/user"
)

type User struct {
	common
	userRepo          repo.User
	teamRepo          repo.Team
	projectRepo       repo.Project
	sceneRepo         repo.Scene
	sceneLockRepo     repo.SceneLock
	layerRepo         repo.Layer
	propertyRepo      repo.Property
	datasetRepo       repo.Dataset
	datasetSchemaRepo repo.DatasetSchema
	transaction       repo.Transaction
	file              gateway.File
	authenticator     gateway.Authenticator
	mailer            gateway.Mailer
	signupSecret      string
}

var (
	//go:embed emails/password_reset_html.tmpl
	passwordResetHTMLTMPLStr string
	//go:embed emails/password_reset_text.tmpl
	passwordResetTextTMPLStr string

	passwordResetTextTMPL *textTmpl.Template
	passwordResetHTMLTMPL *htmlTmpl.Template
)

func init() {
	var err error
	passwordResetTextTMPL, err = textTmpl.New("passwordReset").Parse(passwordResetTextTMPLStr)
	if err != nil {
		log.Panicf("password reset email template parse error: %s\n", err)
	}
	passwordResetHTMLTMPL, err = htmlTmpl.New("passwordReset").Parse(passwordResetHTMLTMPLStr)
	if err != nil {
		log.Panicf("password reset email template parse error: %s\n", err)
	}
}

func NewUser(r *repo.Container, g *gateway.Container, signupSecret string) interfaces.User {
	return &User{
		userRepo:          r.User,
		teamRepo:          r.Team,
		projectRepo:       r.Project,
		sceneRepo:         r.Scene,
		sceneLockRepo:     r.SceneLock,
		layerRepo:         r.Layer,
		propertyRepo:      r.Property,
		datasetRepo:       r.Dataset,
		datasetSchemaRepo: r.DatasetSchema,
		transaction:       r.Transaction,
		file:              g.File,
		authenticator:     g.Authenticator,
		signupSecret:      signupSecret,
		mailer:            g.Mailer,
	}
}

func (i *User) Fetch(ctx context.Context, ids []id.UserID, operator *usecase.Operator) ([]*user.User, error) {
	if err := i.OnlyOperator(operator); err != nil {
		return nil, err
	}
	res, err := i.userRepo.FindByIDs(ctx, ids)
	if err != nil {
		return res, err
	}
	// filter
	for k, u := range res {
		teams, err := i.teamRepo.FindByUser(ctx, u.ID())
		if err != nil {
			return res, err
		}
		teamIDs := make([]id.TeamID, 0, len(teams))
		for _, t := range teams {
			if t != nil {
				teamIDs = append(teamIDs, t.ID())
			}
		}
		if !operator.IsReadableTeamsIncluded(teamIDs) {
			res[k] = nil
		}
	}
	return res, nil
}

func (i *User) Signup(ctx context.Context, inp interfaces.SignupParam) (u *user.User, _ *user.Team, err error) {
	var team *user.Team
	var email, name string
	var auth *user.Auth
	var tx repo.Tx

	if inp.Secret != nil && inp.Sub != nil {
		// Auth0
		if i.signupSecret != "" && *inp.Secret != i.signupSecret {
			return nil, nil, interfaces.ErrSignupInvalidSecret
		}

		if len(*inp.Sub) == 0 {
			return nil, nil, errors.New("sub is required")
		}

		tx, err = i.transaction.Begin()
		if err != nil {
			return nil, nil, err
		}
		defer func() {
			if err2 := tx.End(ctx); err == nil && err2 != nil {
				err = err2
			}
		}()

		name, email, auth, err = i.auth0Signup(ctx, inp)
		if err != nil {
			return
		}

	} else if inp.Name != nil && inp.Email != nil && inp.Password != nil {
		if *inp.Name == "" {
			return nil, nil, interfaces.ErrSignupInvalidName
		}
		if _, err := mail.ParseAddress(*inp.Email); err != nil {
			return nil, nil, interfaces.ErrInvalidUserEmail
		}
		if *inp.Password == "" {
			return nil, nil, interfaces.ErrSignupInvalidPassword
		}

		var unverifiedUser *user.User
		var unverifiedTeam *user.Team

		tx, err = i.transaction.Begin()
		if err != nil {
			return nil, nil, err
		}
		defer func() {
			if err2 := tx.End(ctx); err == nil && err2 != nil {
				err = err2
			}
		}()

		name, email, unverifiedUser, unverifiedTeam, err = i.authSystemSignup(ctx, inp)
		if err != nil {
			return
		}
		if unverifiedUser != nil && unverifiedTeam != nil {
			return unverifiedUser, unverifiedTeam, nil
		}
	}

	// Check if team already exists
	if inp.TeamID != nil {
		existed, err := i.teamRepo.FindByID(ctx, *inp.TeamID)
		if err != nil && !errors.Is(err, rerror.ErrNotFound) {
			return nil, nil, err
		}
		if existed != nil {
			return nil, nil, errors.New("existed team")
		}
	}

	// Initialize user and team
	u, team, err = user.Init(user.InitParams{
		Email:    email,
		Name:     name,
		Sub:      auth,
		Password: *inp.Password,
		Lang:     inp.Lang,
		Theme:    inp.Theme,
		UserID:   inp.UserID,
		TeamID:   inp.TeamID,
	})
	if err != nil {
		return nil, nil, err
	}
	if err := i.userRepo.Save(ctx, u); err != nil {
		return nil, nil, err
	}
	if err := i.teamRepo.Save(ctx, team); err != nil {
		return nil, nil, err
	}
	if tx != nil {
		tx.Commit()
	}

	return u, team, nil
}

func (i *User) authSystemSignup(ctx context.Context, inp interfaces.SignupParam) (string, string, *user.User, *user.Team, error) {
	// Check if user email already exists
	existed, err := i.userRepo.FindByEmail(ctx, *inp.Email)
	if err != nil && !errors.Is(err, rerror.ErrNotFound) {
		return "", "", nil, nil, err
	}

	if existed != nil {
		if existed.Verification().IsVerified() {
			return "", "", nil, nil, errors.New("existed user email")
		} else {
			//	if user exists but not verified -> create a new verification
			if err := i.CreateVerification(ctx, *inp.Email); err != nil {
				return "", "", nil, nil, err
			} else {
				team, err := i.teamRepo.FindByID(ctx, existed.Team())
				if err != nil && !errors.Is(err, rerror.ErrNotFound) {
					return "", "", nil, nil, err
				}
				return "", "", existed, team, nil
			}
		}
	}

	return *inp.Name, *inp.Email, nil, nil, nil
}

func (i *User) auth0Signup(ctx context.Context, inp interfaces.SignupParam) (string, string, *user.Auth, error) {
	// Check if user already exists
	existed, err := i.userRepo.FindByAuth0Sub(ctx, *inp.Sub)
	if err != nil && !errors.Is(err, rerror.ErrNotFound) {
		return "", "", nil, err
	}
	if existed != nil {
		return "", "", nil, errors.New("existed user")
	}

	if inp.UserID != nil {
		existed, err := i.userRepo.FindByID(ctx, *inp.UserID)
		if err != nil && !errors.Is(err, rerror.ErrNotFound) {
			return "", "", nil, err
		}
		if existed != nil {
			return "", "", nil, errors.New("existed user")
		}
	}

	// Fetch user info
	ui, err := i.authenticator.FetchUser(*inp.Sub)
	if err != nil {
		return "", "", nil, err
	}

	// Check if user and team already exists
	existed, err = i.userRepo.FindByEmail(ctx, ui.Email)
	if err != nil && !errors.Is(err, rerror.ErrNotFound) {
		return "", "", nil, err
	}
	if existed != nil {
		return "", "", nil, errors.New("existed user")
	}

	return ui.Name, ui.Email, user.AuthFromAuth0Sub(*inp.Sub).Ref(), nil
}

func (i *User) GetUserByCredentials(ctx context.Context, inp interfaces.GetUserByCredentials) (u *user.User, err error) {
	u, err = i.userRepo.FindByNameOrEmail(ctx, inp.Email)
	if err != nil && !errors.Is(rerror.ErrNotFound, err) {
		return nil, err
	} else if u == nil {
		return nil, interfaces.ErrInvalidUserEmail
	}
	matched, err := u.MatchPassword(inp.Password)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, interfaces.ErrSignupInvalidPassword
	}
	return u, nil
}

func (i *User) GetUserBySubject(ctx context.Context, sub string) (u *user.User, err error) {
	u, err = i.userRepo.FindByAuth0Sub(ctx, sub)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (i *User) StartPasswordReset(ctx context.Context, email string) error {
	tx, err := i.transaction.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	u, err := i.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	pr := user.NewPasswordReset()
	u.SetPasswordReset(pr)

	if err := i.userRepo.Save(ctx, u); err != nil {
		return err
	}

	var TextOut, HTMLOut bytes.Buffer
	link := "localhost:3000/?pwd-reset-token=" + pr.Token
	err = passwordResetTextTMPL.Execute(&TextOut, link)
	if err != nil {
		return err
	}
	err = passwordResetHTMLTMPL.Execute(&HTMLOut, link)
	if err != nil {
		return err
	}

	err = i.mailer.SendMail([]gateway.Contact{
		{
			Email: u.Email(),
			Name:  u.Name(),
		},
	}, "Password reset", TextOut.String(), HTMLOut.String())
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (i *User) PasswordReset(ctx context.Context, password, token string) error {
	tx, err := i.transaction.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	u, err := i.userRepo.FindByPasswordResetRequest(ctx, token)
	if err != nil {
		return err
	}

	passwordReset := u.PasswordReset()
	ok := passwordReset.Validate(token)

	if !ok {
		return interfaces.ErrUserInvalidPasswordReset
	}

	u.SetPasswordReset(nil)

	if err := u.SetPassword(password); err != nil {
		return err
	}

	if err := i.userRepo.Save(ctx, u); err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (i *User) UpdateMe(ctx context.Context, p interfaces.UpdateMeParam, operator *usecase.Operator) (u *user.User, err error) {
	if err := i.OnlyOperator(operator); err != nil {
		return nil, err
	}

	if p.Password != nil {
		if p.PasswordConfirmation == nil || *p.Password != *p.PasswordConfirmation {
			return nil, interfaces.ErrUserInvalidPasswordConfirmation
		}
	}

	tx, err := i.transaction.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	var team *user.Team

	u, err = i.userRepo.FindByID(ctx, operator.User)
	if err != nil {
		return nil, err
	}

	if p.Name != nil {
		oldName := u.Name()
		u.UpdateName(*p.Name)

		team, err = i.teamRepo.FindByID(ctx, u.Team())
		if err != nil && !errors.Is(err, rerror.ErrNotFound) {
			return nil, err
		}

		tn := team.Name()
		if tn == "" || tn == oldName {
			team.Rename(*p.Name)
		} else {
			team = nil
		}
	}
	if p.Email != nil {
		u.UpdateEmail(*p.Email)
	}
	if p.Lang != nil {
		u.UpdateLang(*p.Lang)
	}
	if p.Theme != nil {
		u.UpdateTheme(*p.Theme)
	}

	// Update Auth0 users
	if p.Name != nil || p.Email != nil || p.Password != nil {
		for _, a := range u.Auths() {
			if _, err := i.authenticator.UpdateUser(gateway.AuthenticatorUpdateUserParam{
				ID:       a.Sub,
				Name:     p.Name,
				Email:    p.Email,
				Password: p.Password,
			}); err != nil {
				return nil, err
			}
		}
	}

	if team != nil {
		err = i.teamRepo.Save(ctx, team)
		if err != nil {
			return nil, err
		}
	}

	err = i.userRepo.Save(ctx, u)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return u, nil
}

func (i *User) RemoveMyAuth(ctx context.Context, authProvider string, operator *usecase.Operator) (u *user.User, err error) {
	if err := i.OnlyOperator(operator); err != nil {
		return nil, err
	}

	tx, err := i.transaction.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	u, err = i.userRepo.FindByID(ctx, operator.User)
	if err != nil {
		return nil, err
	}

	u.RemoveAuthByProvider(authProvider)

	err = i.userRepo.Save(ctx, u)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return u, nil
}

func (i *User) SearchUser(ctx context.Context, nameOrEmail string, operator *usecase.Operator) (u *user.User, err error) {
	u, err = i.userRepo.FindByNameOrEmail(ctx, nameOrEmail)
	if err != nil && !errors.Is(err, rerror.ErrNotFound) {
		return nil, err
	}
	return u, nil
}

func (i *User) DeleteMe(ctx context.Context, userID id.UserID, operator *usecase.Operator) (err error) {
	if operator == nil || operator.User.IsNil() {
		return nil
	}

	if userID.IsNil() || userID != operator.User {
		return errors.New("invalid user id")
	}

	tx, err := i.transaction.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	u, err := i.userRepo.FindByID(ctx, userID)
	if err != nil && !errors.Is(err, rerror.ErrNotFound) {
		return err
	}
	if u == nil {
		return nil
	}

	teams, err := i.teamRepo.FindByUser(ctx, u.ID())
	if err != nil {
		return err
	}

	deleter := ProjectDeleter{
		SceneDeleter: SceneDeleter{
			Scene:         i.sceneRepo,
			SceneLock:     i.sceneLockRepo,
			Layer:         i.layerRepo,
			Property:      i.propertyRepo,
			Dataset:       i.datasetRepo,
			DatasetSchema: i.datasetSchemaRepo,
		},
		File:    i.file,
		Project: i.projectRepo,
	}
	updatedTeams := make([]*user.Team, 0, len(teams))
	deletedTeams := []id.TeamID{}

	for _, team := range teams {
		if !team.IsPersonal() && !team.Members().IsOnlyOwner(u.ID()) {
			_ = team.Members().Leave(u.ID())
			updatedTeams = append(updatedTeams, team)
			continue
		}

		// Delete all projects
		err := repo.IterateProjectsByTeam(i.projectRepo, ctx, team.ID(), 50, func(projects []*project.Project) error {
			for _, prj := range projects {
				if err := deleter.Delete(ctx, prj, true, operator); err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return err
		}

		deletedTeams = append(deletedTeams, team.ID())
	}

	// Save teams
	if err := i.teamRepo.SaveAll(ctx, updatedTeams); err != nil {
		return err
	}

	// Delete teams
	if err := i.teamRepo.RemoveAll(ctx, deletedTeams); err != nil {
		return err
	}

	// Delete user
	if err := i.userRepo.Remove(ctx, u.ID()); err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (i *User) CreateVerification(ctx context.Context, email string) error {
	tx, err := i.transaction.Begin()
	if err != nil {
		return err
	}
	u, err := i.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	u.SetVerification(user.NewVerification())
	err = i.userRepo.Save(ctx, u)
	if err != nil {
		return err
	}

	err = i.mailer.SendMail([]gateway.Contact{
		{
			Email: u.Email(),
			Name:  u.Name(),
		},
	}, "email verification", "", "")
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (i *User) VerifyUser(ctx context.Context, code string) (*user.User, error) {
	tx, err := i.transaction.Begin()
	if err != nil {
		return nil, err
	}
	u, err := i.userRepo.FindByVerification(ctx, code)
	if err != nil {
		return nil, err
	}
	if u.Verification().IsExpired() {
		return nil, errors.New("verification expired")
	}
	u.Verification().SetVerified(true)
	err = i.userRepo.Save(ctx, u)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return u, nil
}

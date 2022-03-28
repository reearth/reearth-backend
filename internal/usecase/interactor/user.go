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
	authSrvUIDomain   string
}

type mailContent struct {
	UserName    string
	Message     string
	Suffix      string
	ActionLabel string
	ActionURL   htmlTmpl.URL
}

var (
	//go:embed emails/auth_html.tmpl
	autHTMLTMPLStr string
	//go:embed emails/auth_text.tmpl
	authTextTMPLStr string

	authTextTMPL *textTmpl.Template
	authHTMLTMPL *htmlTmpl.Template

	passwordResetMailContent mailContent
)

func init() {
	var err error
	authTextTMPL, err = textTmpl.New("passwordReset").Parse(authTextTMPLStr)
	if err != nil {
		log.Panicf("password reset email template parse error: %s\n", err)
	}
	authHTMLTMPL, err = htmlTmpl.New("passwordReset").Parse(autHTMLTMPLStr)
	if err != nil {
		log.Panicf("password reset email template parse error: %s\n", err)
	}

	passwordResetMailContent = mailContent{
		Message:     "Thank you for using Re:Earth. We’ve received a request to reset your password. If this was you, please click the link below to confirm and change your password.",
		Suffix:      "If you did not mean to reset your password, then you can ignore this email.",
		ActionLabel: "Confirm to reset your password",
	}
}

func NewUser(r *repo.Container, g *gateway.Container, signupSecret, authSrcUIDomain string) interfaces.User {
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
		authSrvUIDomain:   authSrcUIDomain,
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
		if !operator.IsReadableTeam(teamIDs...) {
			res[k] = nil
		}
	}
	return res, nil
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
	if u.Verification() == nil || !u.Verification().IsVerified() {
		return nil, interfaces.ErrNotVerifiedUser
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
	link := i.authSrvUIDomain + "/?pwd-reset-token=" + pr.Token
	passwordResetMailContent.UserName = u.Name()
	passwordResetMailContent.ActionURL = htmlTmpl.URL(link)

	if err := authTextTMPL.Execute(&TextOut, passwordResetMailContent); err != nil {
		return err
	}
	if err := authHTMLTMPL.Execute(&HTMLOut, passwordResetMailContent); err != nil {
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

	if p.Name != nil && *p.Name != u.Name() {
		// username should not be a valid mail
		if _, err := mail.ParseAddress(*p.Name); err == nil {
			return nil, interfaces.ErrSignupInvalidName
		}
		// make sure the username is not exists
		if userByName, _ := i.userRepo.FindByName(ctx, *p.Name); userByName != nil {
			return nil, interfaces.ErrSignupInvalidName
		}
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
		if err := u.UpdateEmail(*p.Email); err != nil {
			return nil, err
		}
	}
	if p.Lang != nil {
		u.UpdateLang(*p.Lang)
	}
	if p.Theme != nil {
		u.UpdateTheme(*p.Theme)
	}

	if p.Password != nil && u.HasAuthProvider("reearth") {
		if err := u.SetPassword(*p.Password); err != nil {
			return nil, err
		}
	}

	// Update Auth0 users
	if p.Name != nil || p.Email != nil || p.Password != nil {
		for _, a := range u.Auths() {
			if a.Provider != "auth0" {
				continue
			}
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
	u, err := i.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	if u.Verification().IsVerified() {
		return nil
	}
	return i.createVerification(ctx, u)
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

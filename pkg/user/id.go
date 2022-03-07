package user

import "github.com/reearth/reearth-backend/pkg/id"

type ID = id.UserID
type TeamID = id.TeamID

var NewID = id.NewUserID
var NewTeamID = id.NewTeamID

var MustID = id.MustUserID
var MustTeamID = id.MustTeamID

var IDFrom = id.UserIDFrom
var TeamIDFrom = id.TeamIDFrom

var IDFromRef = id.UserIDFromRef
var TeamIDFromRef = id.TeamIDFromRef

var IDFromRefID = id.UserIDFromRefID
var TeamIDFromRefID = id.TeamIDFromRefID

var ErrInvalidID = id.ErrInvalidID

type TeamIDList []TeamID

func (l TeamIDList) Filter(ids ...TeamID) TeamIDList {
	res := make(TeamIDList, 0, len(l))
	for _, t := range l {
		for _, t2 := range ids {
			if t == t2 {
				res = append(res, t)
			}
		}
	}
	return res
}

func (l TeamIDList) Includes(ids ...TeamID) bool {
	for _, t := range l {
		for _, t2 := range ids {
			if t == t2 {
				return true
			}
		}
	}
	return false
}

func (k TeamIDList) Len() int {
	return len(k)
}

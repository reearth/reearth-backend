package google

import (
	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"github.com/reearth/reearth-backend/pkg/file"
)

type google struct {
}

func NewGoogle() gateway.Google {
	return &google{}
}

func (g google) FetchCSV(token string, fileId string, sheetName string) (*file.File, error) {
	return fetchCSV(token, fileId, sheetName)
}

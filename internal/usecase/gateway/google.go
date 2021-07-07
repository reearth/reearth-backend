package gateway

import (
	"github.com/reearth/reearth-backend/pkg/file"
)

type Google interface {
	FetchCSV(token string, fileId string, sheetName string) (*file.File, error)
}

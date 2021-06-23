package gateway

import (
	"github.com/reearth/reearth-backend/pkg/file"
)

type CSVDatasource interface {
	Fetch(token string, fileId string, sheetName string) (*file.File, error)
}

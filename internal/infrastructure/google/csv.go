package google

import (
	"github.com/reearth/reearth-backend/internal/usecase/gateway"
	"github.com/reearth/reearth-backend/pkg/file"
)

type csv struct {
}

func NewCSV() gateway.CSVDatasource {
	return &csv{}
}

func (c csv) Fetch(token string, fileId string, sheetName string) (*file.File, error) {
	return fetchCSV(token, fileId, sheetName)
}

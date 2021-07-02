package google

import (
	"fmt"
	"net/http"

	"github.com/reearth/reearth-backend/pkg/file"
)

func fetchCSV(token string, fileId string, sheetName string) (*file.File, error) {
	url := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s", fileId, sheetName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode=%d", res.StatusCode)
	}

	return &file.File{
		Content:     res.Body,
		Name:        sheetName,
		Fullpath:    sheetName,
		Size:        0,
		ContentType: "text/csv",
	}, nil
}

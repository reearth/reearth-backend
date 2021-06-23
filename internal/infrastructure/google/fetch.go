package google

import (
	"fmt"
	"io"
	"io/ioutil"
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
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode=%d", res.StatusCode)
	}

	out, err := ioutil.TempFile("", fileId+".*.csv")
	if err != nil {
		return nil, err
	}

	size, err := io.Copy(out, res.Body)
	if err != nil {
		return nil, err
	}

	_, err = out.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return &file.File{
		Content:     out,
		Name:        out.Name(),
		Fullpath:    "-",
		Size:        size,
		ContentType: "text/csv",
	}, nil
}

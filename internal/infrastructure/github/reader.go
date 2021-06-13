package github

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"strings"
)

var ErrReearthFileNotFound = errors.New("can't find reearth.json file")

func readZipResponse(r io.ReadCloser) (*zip.Reader, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Close()
	}()

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, err
	}
	return zipReader, nil
}

func readZipFile(zf *zip.File) (io.ReadCloser, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	return f, nil
}

func readRearthFile(zr *zip.Reader) (io.Reader, error) {
	for _, zf := range zr.File {
		if strings.HasSuffix(zf.Name, "reearth.json") {
			content, err := readZipFile(zf)

			if err != nil {
				return nil, err
			}
			return content, nil
		}
	}
	return nil, ErrReearthFileNotFound
}

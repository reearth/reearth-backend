package github

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func fetchURL(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode=%d", res.StatusCode)
	}

	return res.Body, nil
}

func fetchArchiveURL(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, err
	}
	return res.Body, nil

}

func FetchPluginContent(ctx context.Context, url string) (io.Reader, error) {
	res, err := fetchArchiveURL(ctx, url)
	if err != nil {
		return nil, err
	}

	zr, err := readZipResponse(res)
	if err != nil {
		return nil, err
	}

	content, err := readRearthFile(zr)
	if err != nil {
		return nil, err
	}
	return content, nil
}

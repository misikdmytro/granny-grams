package client

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"os"
)

type WebClient interface {
	DownloadImageToFile(context.Context, string, string) error
}

type webClient struct {
	*http.Client
}

// DownloadImageToFile implements WebClient.
func (w *webClient) DownloadImageToFile(ctx context.Context, url, path string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := w.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		return err
	}

	return writer.Flush()
}

func NewWebClient() WebClient {
	client := &http.Client{}
	return &webClient{client}
}

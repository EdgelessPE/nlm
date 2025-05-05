package utils

import (
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadFile(url string, path string) error {
	log.Printf("Downloading file from %s to %s", url, path)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Download completed")
	return nil
}

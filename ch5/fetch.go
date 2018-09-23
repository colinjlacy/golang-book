package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, bytes int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "",0, err
	}
	defer resp.Body.Close()

	filename = path.Base(resp.Request.URL.Path)

	if filename == "/" {
		filename = "index.html"
	}

	file, err := os.Create(filename)
	if err != nil {
		return "", 0, nil
	}

	bytes, err = io.Copy(file, resp.Body)
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}

	return
}

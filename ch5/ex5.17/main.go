package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	return local, n, err
}

func main() {
	for _, url := range os.Args[1:] {
		name, numBytes, err := fetch(url)
		if err != nil {
			fmt.Println("failed to fetch url %v", err)
			continue
		}
		fmt.Println("url %s ,filename %s , num bytes got %v", url, name, numBytes)
	}
}

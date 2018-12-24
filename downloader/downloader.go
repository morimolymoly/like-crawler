package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/morimolymoly/like-crawler/config"
)

func getFilenameFromURL(ur string) (string, error) {
	return path.Base(ur), nil
}

// DownloadAll ... download given urls all
func DownloadAll(urls []string) []error {
	c := config.GetInstance()
	errs := []error{}
	for i, u := range urls {
		fmt.Printf("downloading[%d/%d] %s\n", i+1, len(urls), u)
		fname, _ := getFilenameFromURL(u)
		fpath := filepath.Join(c.SavePath, fname)
		err := downloadFile(fpath, u)
		if err != nil {
			errs = append(errs, err)
		}
		time.Sleep(1 * time.Second)
	}
	return errs
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

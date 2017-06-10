package alfred

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Download(file, url string) error {
	log.Printf("[i] Downloading %v to %v", url, file)

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status code of %s was %v", url, res.Status)
	}
	modStr := res.Header.Get("Last-Modified")
	modTime, err := http.ParseTime(modStr)
	if err != nil {
		modTime = time.Now()
	}
	tmp := file + ".tmp"
	os.Remove(tmp)
	os.Remove(file)
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, res.Body)
	res.Body.Close()
	if err != nil {
		return fmt.Errorf("error copying %v to %v: %v", url, file, err)
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Chtimes(tmp, modTime, modTime); err != nil {
		return err
	}
	if err := os.Rename(tmp, file); err != nil {
		return err
	}

	return nil
}

package alfred

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	file := "t.alfredworkflow"
	url := "https://github.com/zgs225/alfred-youdao/releases/download/v1.2.0/YouDaoDict-v1.2.0.alfredworkflow"
	err := Download(file, url)
	if err != nil {
		t.Errorf("Downloading %v error: %v", url, err)
	}

	f, err := os.Open(file)
	if err != nil {
		t.Errorf("Can't open downloaded file %v: %v", file, err)
	}
	defer f.Close()
	os.Remove(file)
}

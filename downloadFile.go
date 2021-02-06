package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	fileUrl := "http://topgoer.com/static/2/9.png"
	if err := DownloadFile("9.png", fileUrl); err != nil {
		panic(err)
	}
}

// download file会将url下载到本地文件，它会在下载时写入，而不是将整个文件加载到内存中。
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

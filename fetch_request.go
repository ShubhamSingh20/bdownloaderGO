package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	humanize "github.com/dustin/go-humanize"
	color "github.com/gookit/color"
)

//var waitGroup = sync.WaitGroup{}

//WriteCounter writing the progress
type WriteCounter struct {
	Total uint64
	url   string
}

//PrintProgress method print progress
func (wc WriteCounter) PrintProgress() {
	color.Blue.Printf("\r%s", strings.Repeat(" ", 50))
	//fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
	color.Blue.Printf("\r[+] Downloading %s.  %s complete",
		wc.url, humanize.Bytes(wc.Total))
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

/*
	contentDepo := resp.Header.Get("Content-Disposition")
	filename := strings.Replace(contentDepo, "attachment; filename=", "", -1)
	folderpath = filepath.Join(folderpath, filename)

	buff, _ := ioutil.ReadFile(tmpFile.Name())
		kind, _ := filetype.Match(buff)
		+ kind.Extension

*/
func getFileNameFromHeader(resp *http.Response) string {
	contentDepo := resp.Header.Get("Content-Disposition")
	originalFilename := strings.Replace(contentDepo, "attachment; filename=", "", -1)
	return originalFilename
}

func getFileNameFromURL(resp *http.Response) string {
	originalFilename := resp.Request.URL.String()
	urlSegments := strings.Split(originalFilename, "/")
	originalFilename = urlSegments[len(urlSegments)-1]
	return originalFilename
}

func converTempFileToOriginalAndCloseFile(tmpFile *os.File, currentfilepath string, resp *http.Response) {
	var originalFilename string
	originalFilename = getFileNameFromHeader(resp)

	if len(originalFilename) == 0 {
		originalFilename = getFileNameFromURL(resp)
	}

	tmpFile.Close()

	dirFileName := filepath.Join(currentfilepath, originalFilename)
	err := os.Rename(tmpFile.Name(), dirFileName)

	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()
	//waitGroup.Done()
}

func downloadFile(url, currentfilepath string) error {
	tmpFile, err := ioutil.TempFile(currentfilepath, "prefix-")

	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer converTempFileToOriginalAndCloseFile(tmpFile, currentfilepath, resp)

	counter := &WriteCounter{url: url}
	_, err = io.Copy(tmpFile, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	fmt.Println()

	if err != nil {
		return err
	}

	return nil
}

func downloadFromURLSlice(filepath string, downloadURLSlice []string) {
	for _, url := range downloadURLSlice {
		func(url string) {
			//waitGroup.Add(1)
			err := downloadFile(url, filepath)
			if err != nil {
				fmt.Println(err)
			}
		}(url)
		//waitGroup.Wait()
	}
}

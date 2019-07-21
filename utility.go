package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
)

func consoleMessageAndExit(message string) {
	color.Red.Println("[-] " + message)
	os.Exit(1)
}

func downloadConformation(urls []string, downloadPath string) {
	color.Yellow.Println("Download from the following URLS's")

	reader := bufio.NewReader(os.Stdin)
	//valid := "Y"

	for _, url := range urls {
		color.Yellow.Println(" |-" + url)
	}

	color.Yellow.Println("\nDownload to " + "'" +
		downloadPath + "'" + "\nConfirm (Y/N) ? ")

	confirm, _ := reader.ReadString('\n')

	confirm = strings.ToUpper(confirm)
	confirm = strings.TrimSpace(strings.Trim(confirm, "\n"))

	switch confirm {
	case "YES", "Y":
		return

	default:
		consoleMessageAndExit("exiting .")
	}

}

type commandLinePaths interface {
	validate(string)
}

//DownloadDir implements the validate interface to validate
//the folder where all the files are to be downloaded
type DownloadDir struct {
	dirPath string
}

func (d DownloadDir) createNewDownloadFolderIFNotExists() string {
	currentDir, err := os.Getwd()

	if err != nil {
		consoleMessageAndExit("[-] Could'nt get the current working directory")
	}

	downloadDirPath := filepath.Join(currentDir, "download")

	if _, err := os.Stat(downloadDirPath); os.IsNotExist(err) {
		os.Mkdir(downloadDirPath, os.ModePerm)
	}

	return downloadDirPath
}

func (d DownloadDir) validateUserDownloadDirPath(userDownloadPath string) {
	fi, err := os.Stat(userDownloadPath)

	if err != nil {
		consoleMessageAndExit(userDownloadPath + " is not a valid path")
		return
	}

	if !fi.IsDir() {
		consoleMessageAndExit(userDownloadPath + " is not a direcotry")
	}
}

func (d *DownloadDir) validate(dirPath string) {
	if len(dirPath) == 0 {
		downloadDirPath := d.createNewDownloadFolderIFNotExists()
		d.dirPath = downloadDirPath
		return
	}

	d.validateUserDownloadDirPath(dirPath)
	//check if the user provided path is valid !
	d.dirPath = dirPath
}

//URLFile implements the validate interface to validate
// a files consisting of URL's
type URLFile struct {
	urlFilePath  string
	userInputURL []string
}

func (uf URLFile) validateURLFile(urlFilePath string) {
	info, err := os.Stat(urlFilePath)

	if os.IsNotExist(err) {
		consoleMessageAndExit(urlFilePath + " does not exists ! ")
	}

	if info.IsDir() {
		consoleMessageAndExit(urlFilePath + " is a directory not a file")
	}

	urlFileExtension := filepath.Ext(strings.TrimSpace(urlFilePath))

	if urlFileExtension != ".txt" {
		consoleMessageAndExit(urlFileExtension + " is provided. Expected extension .txt ")
	}
}

func (uf *URLFile) appendURLsToArgsFromFile() {
	file, err := os.Open(uf.urlFilePath)

	if err != nil {
		consoleMessageAndExit(uf.urlFilePath + " cannot be opened ")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		uf.userInputURL = append(uf.userInputURL, scanner.Text())
	}

}

func (uf *URLFile) validate(urlFilePath string) {
	if len(urlFilePath) == 0 {
		uf.urlFilePath = ""
		return
	}

	uf.validateURLFile(urlFilePath)
	uf.urlFilePath = urlFilePath

}

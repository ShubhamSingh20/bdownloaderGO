package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func addInfoToApp(app *cli.App) {
	app.Name = "Go BDownloader"
	app.Author = "Shubham Singh"
	app.Usage = "Let you download mutliple files asynchronoulsy"
	app.Version = "0.1.0"
}

func addCommandsToApp(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "bdownload",
			Aliases: []string{"bdl"},
			Usage:   "bdl <url>",

			Flags: []cli.Flag{
				cli.StringFlag{Name: "download-dir,dout"},
				cli.StringFlag{Name: "url-list"},
			},

			Action: func(c *cli.Context) {
				userURL := c.Args()[:]

				downloadDir := c.String("dout")
				urlFilePath := c.String("url-list")

				currentDownloadDir := DownloadDir{}
				currentDownloadDir.validate(downloadDir)

				currentURLFile := URLFile{userInputURL: userURL}

				currentURLFile.validate(urlFilePath)
				currentURLFile.appendURLsToArgsFromFile()

				downloadConformation(currentURLFile.userInputURL,
					currentDownloadDir.dirPath)

				downloadFromURLSlice(currentDownloadDir.dirPath, currentURLFile.userInputURL)

			},
		},
	}

}

func main() {

	cmdApp := cli.NewApp()
	addInfoToApp(cmdApp)
	addCommandsToApp(cmdApp)

	err := cmdApp.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}

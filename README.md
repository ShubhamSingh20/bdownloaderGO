# bdownloaderGO
A command line implementation of a batch downloader that uses Go Routines.
A simple Command Line application written, for utilizing go routines and making batch downloading much simpler. 
It works by creating a temporary file and writing all the contents from the GET request into the file. And in the end request it will
get the filename from header or the url and, change the extension type of the file converting it into the given file type.

**working**
* clone the git and build it from the source using ```go build * ``` command.
* run the file as ``` ./bdownload.exe bdl -url-list <path to a file containing URLs> -dout <where to store the files> ```
* It will ask for permission to download after prompting all the files to be downloaded. choose ```yes/no```
* You are good to go :).

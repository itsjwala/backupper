package main

import (
	"github.com/jlaffaye/ftp"
	"log"
	"io"
	"os"
	"path/filepath"
)

type MyFile struct {
	dirname  string
	filename string
	size     uint64
}

const CHANNEL_SIZE int = 100000
const CONCURRENCY int = 1

const ip string = "192.168.0.102"
const port string = "9999"
const username string = "anonymous"
const password string = "anonymous"

func ensureDir(filename string) {

	dirname := filepath.Dir(filename)
	if _, serr := os.Stat(dirname); serr != nil {
		merr := os.MkdirAll(dirname, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}

}

func readFrom(path string) uint64 {
	fileinfo, err := os.Stat(path)

	if err != nil {
		return 0
	}

	filesize := fileinfo.Size()

	return uint64(filesize)

}


func download(client *ftp.ServerConn,file MyFile) {

	path := "."+file.dirname+file.filename
	log.Printf("Downloading file %s",path)

	filesize := file.size
	startFrom := readFrom(path)

	if startFrom == 0 {
		log.Printf("Fetching : %d bytes", filesize-startFrom)
	}else if startFrom < filesize{
		log.Printf("Fetching remaining : %d bytes", filesize-startFrom)
	}else{
		log.Printf("File already downloaded skipping..")
		return
	}


	if resp, err := client.RetrFrom(path, startFrom); err != nil {
		log.Fatal(err)
	} else {
		defer resp.Close()
		ensureDir(path)

		var flags int

		if startFrom == 0 {
			flags = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
		} else {
			flags = os.O_APPEND | os.O_WRONLY
		}

		newFile, err := os.OpenFile(path, flags, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()

		io.Copy(newFile,resp)
	}

}


/*

ftp => not concurrent (using) (retrieve from works)
goftp => concurrent

*/

func makeFtpClient() (*ftp.ServerConn, error) {

	client, err := ftp.Connect(ip + ":" + port)

	if err != nil {
		return nil, err
	}

	if err := client.Login(username, password); err != nil {
		return nil, err
	}

	return client, nil
}

func fetchFilesFromDirectoryRecurSively(client *ftp.ServerConn, dir string, include_dir map[string]bool, exclude_dir map[string]bool, mychan chan MyFile) error {
	include_dir[dir] = true

	entries, err := client.List(dir)

	if err != nil {
		return err
	}

	for _, entry := range entries {

		switch entry.Type {
		case ftp.EntryTypeFile:
			newfile := MyFile{dirname: dir, filename: entry.Name, size: entry.Size}

			mychan <- newfile

		case ftp.EntryTypeFolder:

			var folderName string

			folderName = dir + entry.Name + "/"

			if _, excluded := exclude_dir[folderName]; !excluded {
				if _, included := include_dir[folderName]; !included {

					fetchFilesFromDirectoryRecurSively(client, folderName, include_dir, exclude_dir, mychan)
				}
			}

		case ftp.EntryTypeLink: //Skip
		}
	}
	return nil
}

func massageDirnameWithSlash(dir string) string {
	length := len(dir)
	if length == 0 {
		return "/"
	}

	if length > 0 && dir[0] == '/' && dir[length-1] == '/' {
		return dir
	} else if dir[0] == '/' {
		return dir + "/"
	} else if dir[length-1] == '/' {
		return "/" + dir
	} else {
		return "/" + dir + "/"
	}

}
func fetchFiles(client *ftp.ServerConn, include []string, exclude []string) chan MyFile {
	file_channel := make(chan MyFile, CHANNEL_SIZE)

	exclude_dir := make(map[string]bool)

	for _, dir := range exclude {
		exclude_dir[massageDirnameWithSlash(dir)] = true
	}

	include_dir := make(map[string]bool)

	for _, dir := range include {

		dir = massageDirnameWithSlash(dir)

		fetchFilesFromDirectoryRecurSively(client, dir, include_dir, exclude_dir, file_channel)
	}

	close(file_channel)

	return file_channel
}

func main() {

	if client, err := makeFtpClient(); err == nil {

		mychan := fetchFiles(client, []string{"Pictures/Screenshots"}, []string{"/Music/NewPipe/"})

		for fileToRead := range  mychan{

				download(client,fileToRead)
		}
		client.Quit()
	} else {
		log.Fatal(err)
	}

}

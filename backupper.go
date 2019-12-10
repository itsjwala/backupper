package main

import (
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type MyFile struct {
	dirname  string
	filename string
	size     uint64
}

const (
	CHANNEL_SIZE int           = 100000
	TIME_OUT     time.Duration = 10 * time.Second
)

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

func download(client *ftp.ServerConn, file MyFile) {

	server_path := file.dirname + file.filename

	path := filepath.Join(config.Base_dir, server_path)

	log.Printf("Downloading into %s", path)

	filesize := file.size
	startFrom := readFrom(path)

	if startFrom == 0 {
		log.Printf("Fetching : %d bytes", filesize-startFrom)
	} else if startFrom < filesize {
		log.Printf("Fetching remaining : %d bytes", filesize-startFrom)
	} else {
		log.Printf("File already downloaded skipping..")
		return
	}

	if resp, err := client.RetrFrom(server_path, startFrom); err != nil {
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

		io.Copy(newFile, resp)
	}

}

/*

ftp => not concurrent (using) (retrieve from works)
goftp => concurrent

*/

func makeFtpClient() (*ftp.ServerConn, error) {

	client, err := ftp.DialTimeout(config.Server+":"+config.Port, TIME_OUT)

	if err != nil {
		return nil, err
	}

	if err := client.Login(config.Username, config.Password); err != nil {
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

func start() {

	if client, err := makeFtpClient(); err == nil {

		mychan := fetchFiles(client, config.Include_dir, config.Exclude_dir)

		for fileToRead := range mychan {
			download(client, fileToRead)
		}
		client.Quit()
	} else {
		log.Fatal(err)
	}

}

const sample_json_config string = `
{
"server":"192.168.0.102",
"port":"9999",
"username":"anonymous",
"password":"anonymous",
"include_dir":["/Download","Music"],
"exclude_dir":["/Music/NewPipe"],
"base_dir":"/home/jigar/backups"
}
`

func help() {

	log.Fatal("Please provide configuration file, sample config json is given below..", sample_json_config, "For more info visit :- github.com/jigarWala/backupper")
}

func greet() {
	log.Println("Backupper is running with following config..")
	log.Printf("server=%s", config.Server)
	log.Printf("port=%s", config.Port)
	log.Printf("username=%s", config.Username)
	log.Printf("password=%s", config.Password)
	log.Printf("include_dir=%s", config.Include_dir)
	log.Printf("exclude_dir=%s", config.Exclude_dir)
	log.Printf("base_dir=%s", config.Base_dir)
}
func main() {

	if len(os.Args) == 1 {
		help()
	}

	loadConfig(os.Args[1])
	greet()
	start()

}

package main

// import "service"

import (
	"github.com/jlaffaye/ftp"
	// "time"
	"log"
	// "io"
	"os"
	"path/filepath"
	"sync"
)

var wg sync.WaitGroup


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

	log.Printf("Already Have %d bytes", filesize)

	return uint64(filesize)

}

func fetch(path string){
	defer wg.Done()
	client, err := ftp.Connect("192.168.0.102:9999")
	log.Println(path)

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Login("anonymous", "anonymous"); err != nil {
		log.Fatal(err)
	}


	filesize, _ := client.FileSize(path)

	log.Printf("Stat from upstream : %d", filesize)
	// os.Stat
	startFrom := readFrom(path)

	if resp, err := client.RetrFrom(path, startFrom); err != nil {
		log.Fatal(err)
	} else {

		ensureDir(path)

		var flags int

		if startFrom == 0 {
			// := os.Create(path)
			flags = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
		} else {
			flags = os.O_APPEND | os.O_WRONLY
		}

		newFile, err := os.OpenFile(path, flags, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()

		data := make([]byte, 1024000)

		// io.Copy(newFile,resp)

		for {

			n, err := resp.Read(data)

			if err != nil {
				if n == 0 {
					log.Println("Read complete")
				}
				break
			}

			newFile.Write(data[:n])

		}

	}

	// log.Println(filesize)

	if err := client.Quit(); err != nil {
		log.Fatal(err)
	}
}

type MyFile struct{
	dirname string
	filename string
	size uint64
}

const CHANNEL_SIZE int = 100000
const CONCURRENCY int = 1

const ip string = "192.168.0.102"
const port string = "9999"
const username string = "anonymous"
const password string = "anonymous"

/*

ftp => not concurrent (using) (retrieve from works)
goftp => concurrent

*/

func makeFtpClient() (*ftp.ServerConn,error){

	client,err := ftp.Connect(ip+":"+port)

	if err != nil {
		return nil,err
	}

	if err := client.Login(username,password); err != nil {
		return nil,err
	}

	return client,nil
}

func fetchFilesFromDirectoryRecurSively(client * ftp.ServerConn,dir string,include_dir map[string]bool,exclude_dir map[string]bool,mychan chan MyFile)(error){
	include_dir[dir] = true


	entries,err := client.List(dir)

	if(err !=nil){
		return err
	}

	for _,entry := range entries{


		switch entry.Type{
			case ftp.EntryTypeFile :
				newfile := MyFile{dirname:dir,filename:entry.Name,size:entry.Size}
				// log.Println(newfile)

				mychan <- newfile

			case ftp.EntryTypeFolder :

				var folderName string

				folderName = dir+entry.Name+"/"

				if _, excluded := exclude_dir[folderName] ; !excluded{
						if _,included := include_dir[folderName] ; !included{

						fetchFilesFromDirectoryRecurSively(client,folderName,include_dir,exclude_dir,mychan)
					}
				}

			case ftp.EntryTypeLink : //Skip
		}
	}
	return nil
}

func massageDirnameWithSlash(dir string)string{
	length := len(dir)
	if(length == 0){
		return "/"
	}

	if(length>0 && dir[0] == '/' && dir[length-1] == '/'){
		return dir
	}else if(dir[0]=='/'){
		return dir+"/"
	}else if(dir[length-1] =='/'){
		return "/"+dir
	}else{
		return "/"+dir+"/"
	}

}
func fetchFiles(client* ftp.ServerConn,include []string,exclude []string) (chan MyFile) {
	file_channel := make(chan MyFile,CHANNEL_SIZE)


	exclude_dir := make(map[string]bool)


	for _,dir := range exclude{
		exclude_dir[massageDirnameWithSlash(dir)] = true
	}


	include_dir := make(map[string]bool)


	for _,dir := range include{


		dir = massageDirnameWithSlash(dir)

		fetchFilesFromDirectoryRecurSively(client,dir,include_dir,exclude_dir,file_channel)
		log.Println(include_dir)

	}

	close(file_channel)


	return file_channel
}

func main() {


	// path := "zzz"
	// path := "Download/lec1.mp4"
	// path := "Download/default.xlsx"
	// wg.Add(1)
	// go fetch(path)
	// path := "Download/linuxmint-19.2-xfce-32bit.iso"
	// wg.Add(1)
	// go fetch(path)

	// log.Println("Waiting for them to die")
	// wg.Wait()


	if client,err := makeFtpClient() ; err==nil {

		// fetchFiles(client,[]string{"","/Download"},[]string{"/Download"})
		for fileToRead := range fetchFiles(client,[]string{"Music/AudioBooks","/Music"},[]string{"/Music/NewPipe/"}){
			log.Println(fileToRead)
		}

		// entries,_ := client.List("Music/")

		// for _,entry := range entries{
		// 	log.Println(entry)
		// }

		client.Quit()
		// client.NoOp()
	}else{
		log.Fatal(err)
	}

}

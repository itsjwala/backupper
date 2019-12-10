> This project is still work in progress!

<div align="center">
<h2>Backupper</h2> 
<p>Command line tool to take phone's media backup to desktop wirelessly</p>
</div>

[![HitCount](http://hits.dwyl.io/jigarWala/backupper.svg)](http://hits.dwyl.io/jigarWala/backupper)

![Gopher image](https://golang.org/doc/gopher/pencil/gopherswrench.jpg)
*Gopher image by [Renee French](https://reneefrench.blogspot.com/), licensed under [Creative Commons 3.0 Attributions license](https://creativecommons.org/licenses/by/3.0/).*
 



-------------------------------------------------

It works by having ftp server on the phone(es file explorer, solid explorer..) and setting up required configuration to start the backup. It has resuming download capabilities, It is fast and configurable.

-------------------------------------------



### Features

* Quite Fast :rocket:

* Resumable Download :inbox_tray:

* Very Configurable :wrench:

### Build from source

* Make sure you have [go](https://golang.org/dl/) installed and have in your path

* fetch this repository

```bash
$ go get github.com/jigarWala/backupper
```

* make executable

```bash
$ go build $GOPATH/src/github.com/jigarWala/backupper
```

### Usage

* Start FTP server on your phone



<img align="center" height="20%" width="35%" src="https://i.imgur.com/BEUGNeW.png"/>


* prepare configuration file for your usecase


Key|Required|Default | Comments
---|---|---| ----
server|yes|NA | compulsory key
port|yes|NA | compulsory key
username | no | anonymous | `"anonymous"` is used for no username
password| no | anonymous | `"anonymous"` is used for no password
include_dir | no | [ "/" ] | `/` or the whole server directory is considered 
exclude_dir | no | [ ] | no directory is excluded
base_dir | no | check comments => | backup is created in the directory where executable `resides`


sample json config file

```json
{
"server":"192.168.0.102",
"port":"9999",
"include_dir":["Download","DCIM/Camera","Music","Movies","WhatsApp"],
"exclude_dir":["/Music/NewPipe","WhatsApp/Databases"],
"base_dir": "/home/jigar/Desktop/backups"
}
```

* Execute the command

```bash
$ ./backupper path/to/configuration-file.json
```


### Checklist

- [X] Can Resume Download from where it stopped?

- [X] Configurable Backups via json config file

- [ ] Better Error handling

- [ ] Better Logging

- [ ] Concurrent Downloads




### Why I made this?


I wanted to backup my screenshots, camera roll, music, whatsapp media etc to my local machine. Ofcourse there are adb configurable scripts available, But I wanted it to be wireless and also there are cloud backups available too like google drive. But it is quite slow as my internet speed is not so fast. I wanted to learn Golang :)

It is very fast and accessible for me and maybe others can find it useful too.

### Contributing

Please reach out to me if you wish to contribute to this project.


### Authors

* **Jigar Wala**  - [jigarWala](https://github.com/jigarWala)

See also the list of [contributors](https://github.com/jigarWala/backupper/contributors) who participated in this project.

### License

This project is licensed under the MIT - see the [LICENSE](./LICENSE) file for details



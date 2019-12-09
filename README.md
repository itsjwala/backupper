> This project is still work in progress!

<div align="center">
<h2>Backupper</h2>
<p>Command line tool to take phone's media backup to desktop wirelessly</p>
</div>


[![HitCount](http://hits.dwyl.io/jigarWala/backupper.svg)](http://hits.dwyl.io/jigarWala/backupper)


-------------------------------------------------

It works by having ftp server on the phone(es file explorer, solid explorer..) and setting up required configuration to start the backup. It has resuming download capabilities, It is fast and configurable.

-------------------------------------------



### Features

* Quite Fast :rocket:

* Resumable Download :inbox_tray:

* Highly Configurable :wrench:

### Build from source

* fetch this repository

```bash
$ go get github.com/jigarWala/backupper
```

* TODO

### Usage

* Start FTP server on your phone


<img align="center" height="15%" width="35%" src="https://i.imgur.com/YG8PQeI.png"/>


* prepare below properties file for your usecase,

```txt
username=anonymous (required)
password=anonymous (required)
server=192.168.0.102 (required)
port=9999 (required)
```

* Execute the command

```bash
$ ./backupper path/to/configuration-file
```


### Checklist

- [X] Can Resume Download from where it stopped?

- [ ] Configurable Backups via properties file

- [ ] Error handling

- [ ] Logging

- [ ] Concurrent Downloads




### Why I made this?


I wanted to backup my screenshots, camera roll, music, whatsapp media etc to my local machine. Ofcourse there are adb configurable scripts available, But I wanted it to be wireless And There are cloud backups available too like google drive. But it is quite slow for me as my internet speed is not so fast Also I wanted to learn Golang :)


It is very fast and accessible for me and maybe others can also find it useful.

### Contributing

Please reach out to me if you wish to contribute to this project.


### Authors

* **Jigar Wala**  - [jigarWala](https://github.com/jigarWala)

See also the list of [contributors](https://github.com/jigarWala/backupper/contributors) who participated in this project.

### License

This project is licensed under the MIT - see the [LICENSE](./LICENSE) file for details



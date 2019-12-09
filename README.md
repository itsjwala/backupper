> This project is still work in progress!

<div align="center">
<h2>Backupper</h2>
<p>Take phone's media backup to desktop</p>
</div>


[![HitCount](http://hits.dwyl.io/jigarWala/backupper.svg)](http://hits.dwyl.io/jigarWala/backupper)


-------------------------------------------------

It works by having ftp server on the phone(es file explorer, solid explorer..) and setting up requied configuration to start the backup. It is 

-------------------------------------------



### Features

* Blazing Fast :rocket:

* Resumable Download

* Highly Configurable

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


*github.com/jlaffaye/ftp*  isn't concurrent and *github.com/secsy/goftp* is concurrent but don't support resuming capabilities



### Why I made this?


I wanted to backup my screenshots, camera roll, music, whatsapp media etc to my local machine. There are cloud backups available like google drive. But it's quite slow for me as my internet speed is not so fast. Also I wanted to learn golang :)


It is very fast and accessible for me and maybe others can also find it useful.

### Contributing

Please reach out to me if you wish to contribute to this project.


### Authors

* **Jigar Wala**  - [jigarWala](https://github.com/jigarWala)

See also the list of [contributors](https://github.com/jigarWala/backupper/contributors) who participated in this project.

### License

This project is licensed under the MIT - see the [LICENSE](./LICENSE) file for details



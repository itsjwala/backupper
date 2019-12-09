> This project is still work in progress!

<div align="center">
<h2>Backupper</h2>
<p>Take phone's media backup to desktop</p>
</div>

[![HitCount](http://hits.dwyl.io/jigarWala/backupper.svg)](http://hits.dwyl.io/jigarWala/backupper)


-------------------------------------------------

It works by having ftp server on the phone(es file explorer, solid explorer..)

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

*

### Usage

* Start FTP server on your phone

<img align="center" src="https://imgur.com/YG8PQeI"/>


* prepare config file, grab key from below

```txt

```

* Execute the command

```bash
$ ./backupper path/to/configuration-file
```

### Checklist

[*] Can Resume Download from where it stopped?
[] Configurable Backups via properties file
[] Error handling
[] Logging
[] Concurrent Downloads

> github.com/jlaffaye/ftp  isn't concurrent and github.com/secsy/goftp is concurrent but don't support resuming capabilities



### Why I made this?


I wanted to backup my screenshots, camera roll, music, whatsapp media etc to my local machine. There are cloud backups available like google drive. But it's quite slow for me as my internet speed is not so fast.


It is very fast and accessible for me and maybe others can also find it useful.

## Contributing

Please reach out to me if you wish to contribute to this project.


## Authors

* **Jigar Wala**  - [jigarWala](https://github.com/jigarWala)

See also the list of [contributors](https://github.com/jigarWala/backupper/contributors) who participated in this project.

## License

This project is licensed under the MIT - see the [LICENSE](./LICENSE) file for details



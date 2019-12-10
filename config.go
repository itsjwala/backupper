package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type MyConfig struct {
	Server      string
	Port        string
	Username    string
	Password    string
	Base_dir    string
	Include_dir []string
	Exclude_dir []string
}

var config MyConfig

func loadConfig(config_file_path string) {

	json_dat, err := ioutil.ReadFile(config_file_path)

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(json_dat), &config)
	if err != nil {
		log.Fatal("Please provide valid json :- ", err)
	}
	if len(config.Include_dir) == 0 {
		config.Include_dir = append(config.Include_dir, "/")
	}
	if config.Username == "" {
		config.Username = "anonymous"
	}
	if config.Password == "" {
		config.Password = "anonymous"
	}
	if config.Base_dir == "" {
		ex, eerr := os.Executable()
		if eerr != nil {
			log.Fatal(eerr)
		}

		config.Base_dir = filepath.Dir(ex)
	}
}

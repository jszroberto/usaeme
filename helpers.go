package main

import (
	"errors"
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

const (
	LOCAL_CONFIG = "local_config.yml"
)

func connectDB() (*DB, error) {
	config, err := ReadLocalConfig()
	var uri string

	if err != nil {
		uri = config.DBUri
	} else {
		appEnv, err := cfenv.Current()
		if err != nil {
			return &DB{}, err
		}
		if redis_services, ok := appEnv.Services["compose-for-redis"]; ok {
			uri = redis_services[0].Credentials["uri"].(string)
		} else {
			return &DB{}, errors.New("compose-for-redis service not bound to the app")
		}
		fmt.Println(uri)
		db := NewDatabase(uri)
		return db, db.Ping()
	}

	appEnv, err := cfenv.Current()
	if err != nil {
		return &DB{}, err
	}
	if redis_services, ok := appEnv.Services["compose-for-redis"]; ok {
		db := NewDatabase(redis_services[0].Credentials["uri"].(string))
		return db, db.Ping()
	} else {
		return &DB{}, errors.New("compose-for-redis service not bound to the app")
	}
}

func ReadLocalConfig() (LocalConfig, error) {

	filename, err := filepath.Abs(LOCAL_CONFIG)
	if err != nil {
		return LocalConfig{}, err
	}
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return LocalConfig{}, err
	}

	var config LocalConfig
	err = yaml.Unmarshal(yamlFile, &config)
	return config, err
}

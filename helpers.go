package main

import (
	"encoding/json"
	"errors"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/uber-go/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

const (
	LOCAL_CONFIG = "local_config.yml"
)

func connectDB(log zap.Logger) (*DB, error) {
	config, err := ReadLocalConfig()
	var uri string

	if err == nil {
		log.Info("Found local configuration")
		uri = config.DBUri
	} else {
		log.Info("Read configuration from CF_ENV")
		appEnv, err := cfenv.Current()
		if err != nil {
			return &DB{}, err
		}
		if db_services, ok := appEnv.Services["cloudant-go-cloudant"]; ok {
			uri, _ = db_services[0].Credentials["url"].(string)
		} else {
			return &DB{}, errors.New("cloudant-go-cloudant service not bound to the app")
		}
	}
	db := NewDatabase(uri, log)
	return db, db.Ping()
}

func ReadLocalConfig() (Config, error) {

	filename, err := filepath.Abs(LOCAL_CONFIG)
	if err != nil {
		return Config{}, err
	}
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	return config, err
}

func NewLogger() zap.Logger {
	return zap.New(zap.NewJSONEncoder(zap.NoTime()))
}

func get(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

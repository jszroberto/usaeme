package main

import (
	"github.com/fjl/go-couchdb"
	"github.com/uber-go/zap"
)

type DB struct {
	client couchdb.Client
	log    zap.Logger
}

const (
	DB_NAME = "usaeme"
)

func NewDatabase(address string, log zap.Logger) *DB {

	client, err := couchdb.NewClient(address, nil)
	if err != nil {
		panic(err)
	}

	//ensure db exists
	//if the db exists the db will be returned anyway
	client.CreateDB(DB_NAME)

	return &DB{*client, log}
}

func (db *DB) Ping() error {
	return db.client.Ping()
}

func (db *DB) IsAccessible() bool {
	return db.Ping() == nil
}

func (db *DB) GetWords() ([]Word, error) {
	var words []Word
	err := db.client.DB(DB_NAME).AllDocs(words, nil)
	if err != nil {
		return nil, err
	}
	db.log.Info("Get Words", zap.Int("size", len(words)))
	return words, nil
}

package main

import (
	"github.com/fjl/go-couchdb"
	"github.com/uber-go/zap"
	"time"
)

type DB struct {
	client couchdb.Client
	log    zap.Logger
}

type AllDocsResult struct {
	TotalRows int   `json:"total_rows"`
	Offset    int   `json:"offset"`
	Rows      []Row `json:"rows"`
}

type Row struct {
	Word  Word                   `json:"doc"`
	Id    string                 `json:"id"`
	Key   string                 `json:"key"`
	Value map[string]interface{} `json:"value"`
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
func (db *DB) GetNexts(words []Word) []Word {

	nexts := make([]Word, 5)
	copy(nexts, words)
	return nexts
}
func (db *DB) GetLearning() ([]Word, error) {
	words, err := db.GetWords()

	if err != nil {
		return nil, err
	}

	learning := []Word{}

	for _, word := range words {
		if !word.IsMastered {
			learning = append(learning, word)
		}
	}
	return learning, nil
}

func (db *DB) GetMastered() ([]Word, error) {
	words, err := db.GetWords()
	if err != nil {
		return nil, err
	}
	mastered := []Word{}

	for _, word := range words {
		if word.IsMastered {
			mastered = append(mastered, word)
		}
	}
	return mastered, nil
}

func (db *DB) GetWords() ([]Word, error) {
	var result AllDocsResult
	err := db.client.DB(DB_NAME).AllDocs(&result, couchdb.Options{"include_docs": true})
	if err != nil {
		return nil, err
	}
	db.log.Info("Get Words", zap.Int("size", len(result.Rows)))
	words := []Word{}

	for _, row := range result.Rows {
		words = append(words, row.Word)
	}
	return words, nil
}

func (db *DB) GetWord(word string) (Word, error) {
	var doc Word
	db.log.Info("Get Word", zap.String("name", word))
	err := db.client.DB(DB_NAME).Get(word, &doc, couchdb.Options{})
	return doc, err
}

func (db *DB) DeleteWord(word string) error {
	db.log.Info("Get Word", zap.String("name", word))
	rev, err := db.client.DB(DB_NAME).Rev(word)

	if err != nil {
		return err
	}
	_, err = db.client.DB(DB_NAME).Delete(word, rev)
	return err
}

func (db *DB) SetWord(word Word) error {
	db.log.Info("Set Word", zap.String("name", word.Name))
	word.CreatedAt = time.Now()
	word.UpdatedAt = time.Now()
	_, err := db.client.DB(DB_NAME).Put(word.Name, word, "")
	return err
}

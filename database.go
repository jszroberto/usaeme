package main

import (
	"gopkg.in/redis.v5"
)

type DB struct {
	client redis.Client
}

func NewDatabase(address string) *DB {
	client := redis.NewClient(&redis.Options{
		Addr: address,
		// Password: password, // no password set
		DB: 0, // use default DB
	})

	return &DB{*client}
}
func (db *DB) Ping() error {
	_, err := db.client.Ping().Result()
	return err
}

func (db *DB) IsAccessible() bool {
	return db.Ping() == nil
}

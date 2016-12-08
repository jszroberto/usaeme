package main

import (
	"time"
)

type Config struct {
	DBUri string `yaml:"db_uri"`
}

type Word struct {
	Name        string    `json:"name"`
	IsMastered  bool      `json:"is_mastered"`
	Frecuency   int       `json:"frecuency"`
	TimesUsed   int       `json:"times_used"`
	NextUse     time.Time `json:"next_use"`
	Deck        string    `json:"deck"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

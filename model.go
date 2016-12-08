package main

type Config struct {
	DBUri string `yaml:"db_uri"`
}

type Word struct {
	Name        string `json:"name"`
	IsMastered  bool   `json:"is_mastered"`
	Frecuency   int    `json:"frecuency"`
	TimesUsed   int    `json:"times_used"`
	NextUse     string `json:"next_use"`
	Deck        string `json:"deck"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

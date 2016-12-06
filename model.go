package main

type LocalConfig struct {
	DBUri string `yaml:"db_uri"`
}

type Word struct {
	Name       string
	IsMastered bool
	Frecuency  int
	TimesUsed  int
	NextUse    string
	Deck       string
}

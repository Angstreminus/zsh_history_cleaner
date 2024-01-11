package main

import (
	"log"
	"zshcleaner/config"
	"zshcleaner/history"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err = history.AnalyzeHistory(config); err != nil {
		log.Fatal(err)
	}
}

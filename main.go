package main

import (
	"log"
	"sync"
)

var results chan esDoc

func main() {
	results = make(chan esDoc)
	config, err := getConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	go getTrafficStats(config)

	var wg sync.WaitGroup
	wg.Add(1)
	go esIndexer(config, &wg)
	wg.Wait()
}

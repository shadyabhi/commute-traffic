package main

import (
	"log"
	"sync"
)

func esIndexer(c config, wg *sync.WaitGroup) {
	for i := 0; i < len(c.Tracks); i++ {
		log.Printf("ES Index: %s", <-results)
	}
	wg.Done()
}

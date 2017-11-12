package main

import (
	"context"
	"log"
	"sync"

	elastic "gopkg.in/olivere/elastic.v5"
)

func esIndexer(c config, wg *sync.WaitGroup) {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		log.Fatalf("Error connecting to ES: %s", err)
	}

	for i := 0; i < len(c.Tracks); i++ {
		doc := <-results
		_, err := client.Index().
			Index("commute-time").
			Type("doc").
			BodyJson(doc).
			Do(ctx)
		if err != nil {
			log.Printf("Error indexing document: %s", err)
		}
		log.Printf("Indexed to ES: doc=%s", doc)
	}
	wg.Done()
}

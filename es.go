package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	elastic "gopkg.in/olivere/elastic.v5"
)

const mapping = `{
        "settings":{
                "number_of_shards":1,
                "number_of_replicas":0
        },
        "mappings":{
                "doc": {
                        "properties": {
                                "timestamp": { "type": "date", "format": "epoch_second", "store": "true" },
                                "start_address": { "type": "string", "store": "true" },
                                "end_address": { "type": "string", "store": "true" },
                                "distance": { "type": "long", "store": "true" },
                                "duration": { "type": "long", "store": "true" },
                                "duration_in_traffic": { "type": "long", "store": "true" }
                        }
                }
        }
}`

func esIndexer(c config, wg *sync.WaitGroup) {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		log.Fatalf("Error connecting to ES: %s", err)
	}
	err = createIndex(client, c.Elasticsearch.IndexName, ctx)
	if err != nil {
		log.Fatalf("Error creating index in ES: %s", err)
	}

	for i := 0; i < len(c.Tracks); i++ {
		doc := <-results
		_, err := client.Index().
			Index(c.Elasticsearch.IndexName).
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

func createIndex(client *elastic.Client, name string, ctx context.Context) error {
	exists, err := client.IndexExists(name).Do(ctx)
	if err != nil {
		return fmt.Errorf("Error checking if index exists: %s", err)
	}
	if !exists {
		if _, err := client.CreateIndex(name).BodyString(mapping).Do(ctx); err != nil {
			return fmt.Errorf("Error creating index: %s", err)
		}
	}
	return nil
}

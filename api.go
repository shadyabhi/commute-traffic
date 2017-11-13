package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
)

var apiURL = "https://maps.googleapis.com/maps/api/directions/json"

type esDoc struct {
	Distance          int64  `json:"distance"`
	Timestamp         string `json:"timestamp"`
	DurationInTraffic int64  `json:"duration_in_traffic"`
	StartAddress      string `json:"start_address"`
	EndAddress        string `json:"end_address"`
	Duration          int64  `json:"duration"`
}

func getTrafficStats(conf config) {

	for _, track := range conf.Tracks {
		epoch := strconv.Itoa(int(time.Now().Unix()))
		resp, err := doAPICall(conf, track.Source, track.Destination, epoch)
		if err != nil {
			log.Printf("Error querying API for src=%s, dst=%s, resp=%s", track.Source, track.Destination, resp)
		}
		publishResults(resp, epoch, track.Source, track.Destination)
	}
}

func doAPICall(conf config, src string, dst string, epoch string) ([]byte, error) {
	log.Printf("Currently working on, SRC=%s, DST=%s", src, dst)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatalf("Error creating GET request: %s", err)
	}
	q := req.URL.Query()
	q.Add("origin", src)
	q.Add("destination", dst)
	q.Add("traffic_mode", "best_guess")
	q.Add("departure_time", epoch)
	q.Add("key", conf.APIKeys[0])
	req.URL.RawQuery = q.Encode()

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error doing GET request: %s", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read body from response: %s", err)
	}
	return data, nil
}

func publishResults(resp []byte, epoch, src, dst string) {
	startAddr, _ := jsonparser.GetString(resp, "routes", "[0]", "legs", "[0]", "start_address")
	startAddr = src + ", " + startAddr
	endAddr, _ := jsonparser.GetString(resp, "routes", "[0]", "legs", "[0]", "end_address")
	endAddr = dst + ", " + endAddr
	distance, _ := jsonparser.GetInt(resp, "routes", "[0]", "legs", "[0]", "distance", "value")
	duration, _ := jsonparser.GetInt(resp, "routes", "[0]", "legs", "[0]", "duration", "value")
	durationTraffic, _ := jsonparser.GetInt(resp, "routes", "[0]", "legs", "[0]", "duration_in_traffic", "value")

	doc := esDoc{
		Distance:          distance,
		Timestamp:         epoch,
		DurationInTraffic: durationTraffic,
		StartAddress:      startAddr,
		EndAddress:        endAddr,
		Duration:          duration,
	}

	results <- doc
}

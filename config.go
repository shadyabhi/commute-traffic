package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	APIKeys []string `yaml:"apiKeys"`
	Tracks  []struct {
		Source      string `yaml:"source"`
		Destination string `yaml:"destination"`
	} `yaml:"tracks"`
}

func getConfig() (c config, err error) {
	f, err := os.Open("config.yaml")
	if err != nil {
		return config{}, fmt.Errorf("Error opening config file: %s", err)
	}
	configBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return config{}, fmt.Errorf("Error reading config file: %s", err)
	}
	err = yaml.Unmarshal(configBytes, &c)
	if err != nil {
		return config{}, fmt.Errorf("Error parsing config file: %s", err)
	}
	return c, nil
}

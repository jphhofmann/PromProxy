package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

type routes struct {
	Location string `yaml:"location"`
	Target   string `yaml:"target"`
	Path     string `yaml:"path"`
}

type config struct {
	Debug     bool              `yaml:"debug"`
	Listen    string            `yaml:"listen"`
	Whitelist []string          `yaml:"whitelist"`
	Routes    map[string]routes `yaml:"routes"`
}

var Cfg config

/* Load configuration file */
func configLoad() {
	f, err := os.Open("/etc/promproxy.yaml")
	if err != nil {
		log.Fatalf("Failed to load config, %err", err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		log.Fatalf("Failed to load config, %v", err)
	}
}

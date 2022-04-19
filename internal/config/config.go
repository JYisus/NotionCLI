package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type database struct {
	Name   string `yaml:"name"`
	Id     string `yaml:"id"`
	Key    string `yaml:"key"`
	Filter string `yaml:"filter" default:"{}"`
}

type Config struct {
	DatabaseId      string     `yaml:"DatabaseId"`
	NotionApiKey    string     `yaml:"NotionApiKey"`
	Databases       []database `yaml:"databases"`
	DefaultDatabase string     `yaml:"defaultDatabase"`
}

func NewConfig(filePath string) (*Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil
	}
	cfg := Config{}
	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("Error loading configuration: %s\n", err)
		return nil, err
	}

	return &cfg, nil
}

func GenerateFile(filePath, notionApiKey, databaseId string) {
	c := Config{
		DatabaseId:   databaseId,
		NotionApiKey: notionApiKey,
	}

	cfg, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("Error marshalling config: %s", err)
	}
	err = os.WriteFile(filePath, cfg, 0644)
	if err != nil {
		log.Fatalf("Error generating config file: %s", err)
	}
}

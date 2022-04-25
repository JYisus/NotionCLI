package config

import (
	"github.com/jyisus/notioncli/entity"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Load(filePath string) (*entity.Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil
	}
	cfg := entity.Config{}
	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("Error loading configuration: %s\n", err)
		return nil, err
	}

	return &cfg, nil
}

func GenerateFile(filePath, notionApiKey, databaseId string) error {
	config := entity.Config{
		DatabaseId:   databaseId,
		NotionApiKey: notionApiKey,
	}

	cfg, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, cfg, 0644)
	if err != nil {
		return err
	}

	return nil
}

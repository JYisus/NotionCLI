package entity

type Config struct {
	DatabaseId      string     `yaml:"DatabaseId"`
	NotionApiKey    string     `yaml:"NotionApiKey"`
	Databases       []Database `yaml:"databases"`
	DefaultDatabase string     `yaml:"defaultDatabase"`
}

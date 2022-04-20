package entity

type Database struct {
	Name   string `yaml:"name"`
	Id     string `yaml:"id"`
	Key    string `yaml:"key"`
	Filter string `yaml:"filter" default:"{}"`
}

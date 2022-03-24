/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/jyisus/notioncli/cmd"
	"github.com/jyisus/notioncli/internal/config"
	"github.com/jyisus/notioncli/internal/notion"
	"github.com/spf13/cobra"
	"log"
)

const configPath = "config.yaml"

func main() {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		fmt.Println(err)
	}

	var rootCmd = &cobra.Command{
		Use:   "notioncli",
		Short: "A brief description of your application",
	}

	if cfg != nil {
		notionClient := notion.NewClient(*cfg)

		rootCmd.AddCommand(cmd.InitAddCommand(notionClient))
		rootCmd.AddCommand(cmd.InitListCommand(notionClient))
	}

	rootCmd.AddCommand(cmd.InitConfigCommand(configPath))

	err = rootCmd.Execute()
	if err != nil {
		log.Fatalln("Error executing command!")
	}
}

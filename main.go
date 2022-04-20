package main

import (
	"fmt"
	"github.com/jyisus/notioncli/cmd"
	"github.com/jyisus/notioncli/infrastructure/notion"
	"github.com/jyisus/notioncli/usecase/config"
	"github.com/jyisus/notioncli/usecase/task"
	"log"
)

const configPath = "config.yaml"

func main() {
	cfg, err := config.Load(configPath)
	if err != nil {
		fmt.Println(err)
	}

	var rootCmd = cmd.InitRootCommand()

	if cfg != nil {
		notionTaskRepository := notion.NewClient(*cfg)
		taskService := task.NewService(*cfg, notionTaskRepository)

		rootCmd.AddCommand(cmd.InitAddCommand(taskService))
		rootCmd.AddCommand(cmd.InitDeleteCommand(taskService))
		rootCmd.AddCommand(cmd.InitListCommand(taskService))
	}

	rootCmd.AddCommand(cmd.InitConfigCommand(configPath))

	err = rootCmd.Execute()
	if err != nil {
		log.Fatalln("Error executing command!")
	}
}

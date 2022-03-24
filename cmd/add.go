/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/jyisus/notioncli/internal/notion"
	"log"

	"github.com/spf13/cobra"
)

func InitAddCommand(notionclient notion.Client) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to database",
		Long: `Add a tast to database.

To configure database and Notion API key run notioncli config`,
		Run: runAdd(notionclient),
	}
	addCmd.Flags().StringP("task", "t", "", "task to add")

	return addCmd
}

func runAdd(notionclient notion.Client) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		task, err := cmd.Flags().GetString("task")
		if err != nil {
			log.Fatalf("Error getting flag: %s", err)
		}
		if task == "" {
			log.Fatalln("You must introduce a task")
		}
		err = notionclient.AddTask(task)
		if err != nil {
			log.Fatalf("Error adding task: %s", err)
		}
	}
}

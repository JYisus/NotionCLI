/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
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

	addCmd.Flags().StringP("database", "d", "default", "Notion Database ID")

	return addCmd
}

func runAdd(notionclient notion.Client) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		var database string
		database, err := cmd.Flags().GetString("database")
		if err != nil {
			fmt.Println("Using default database")
			database = "default"
		}
		if len(cmd.Flags().Args()) < 1 {
			log.Fatalln("You must introduce a task")
		}
		task := cmd.Flags().Args()[len(cmd.Flags().Args())-1]
		err = notionclient.AddTask(database, task)
		if err != nil {
			log.Fatalf("Error adding task: %s", err)
		}
	}
}

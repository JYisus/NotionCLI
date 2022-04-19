/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/jyisus/notioncli/internal/notion"
	"log"

	"github.com/spf13/cobra"
)

func InitDeleteCommand(notionclient notion.Client) *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task to database",
		Long: `Delete a task to database.

To configure database and Notion API key run notioncli config`,
		Run: runDelete(notionclient),
	}

	deleteCmd.Flags().StringP("taskId", "t", "", "Notion Task ID")

	return deleteCmd
}

func runDelete(notionclient notion.Client) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		taskId, err := cmd.Flags().GetString("taskId")
		if err != nil {
			log.Fatalf("Error getting task ID arg: %s", err)
		}
		err = notionclient.DeleteTask(taskId)
		if err != nil {
			log.Fatalf("Error adding task: %s", err)
		}
	}
}

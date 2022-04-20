package cmd

import (
	"github.com/jyisus/notioncli/usecase/task"
	"log"

	"github.com/spf13/cobra"
)

func InitDeleteCommand(service task.Service) *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task to database",
		Long: `Delete a task to database.

To configure database and Notion API key run notioncli config`,
		Run: runDelete(service),
	}

	deleteCmd.Flags().StringP("taskId", "t", "", "Notion Task ID")

	return deleteCmd
}

func runDelete(service task.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		taskId, err := cmd.Flags().GetString("taskId")
		if err != nil {
			log.Fatalf("Error getting task ID arg: %s", err)
		}
		err = service.DeleteTask(taskId)
		if err != nil {
			log.Fatalf("Error adding task: %s", err)
		}
	}
}

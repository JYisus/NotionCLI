package cmd

import (
	"github.com/jyisus/notioncli/usecase/task"
	"log"

	"github.com/spf13/cobra"
)

func InitAddCommand(service task.Service) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to database",
		Long: `Add a tast to database.

To configure database and Notion API key run notioncli config`,
		Run: runAdd(service),
	}

	addCmd.Flags().StringP("database", "d", "default", "Notion Database ID")

	return addCmd
}

func runAdd(service task.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		var database string
		database, err := cmd.Flags().GetString("database")
		if err != nil {
			log.Fatalf("Error getting database arg: %s", err)
		}
		if len(cmd.Flags().Args()) < 1 {
			log.Fatalln("You must introduce a task")
		}
		taskText := cmd.Flags().Args()[len(cmd.Flags().Args())-1]
		err = service.AddTask(database, taskText)
		if err != nil {
			log.Fatalf("Error adding task: %s", err)
		}
	}
}

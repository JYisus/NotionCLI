package cmd

import (
	writer "github.com/jyisus/notioncli/infrastructure/terminal_writer"
	"github.com/jyisus/notioncli/usecase/task"
	"log"

	"github.com/spf13/cobra"
)

func InitListCommand(service task.Service) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List task on database",
		Run:   runList(service),
	}
	listCmd.Flags().StringP("database", "d", "default", "Notion Database ID")

	return listCmd
}

func runList(service task.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		database, err := cmd.Flags().GetString("database")
		if err != nil {
			log.Fatalf("Error getting database arg: %s", err)
		}

		w := writer.NewTerminalWriter()

		err = service.ListTasks(w, database)
		if err != nil {
			log.Fatalf("Error adding task: %s", err)
		}
	}
}

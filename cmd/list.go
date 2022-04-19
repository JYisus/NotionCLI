package cmd

import (
	"fmt"
	"github.com/jyisus/notioncli/internal/notion"
	"log"

	"github.com/spf13/cobra"
)

func InitListCommand(notionclient notion.Client) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks on database",
		Run:   runList(notionclient),
	}
	listCmd.Flags().StringP("database", "d", "default", "Notion Database ID")

	return listCmd
}

func runList(notionclient notion.Client) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		database, err := cmd.Flags().GetString("database")
		if err != nil {
			panic(err)
		}

		res, err := notionclient.ListTasks(database)
		if err != nil {
			log.Fatalf("Error adding task: %s", err)
		}

		for index, task := range res {
			fmt.Printf("%d. %s\n", index+1, task)
		}
	}
}

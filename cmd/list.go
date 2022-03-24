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
	listCmd.Flags().StringP("task", "t", "", "task to add")

	return listCmd
}

func runList(notionclient notion.Client) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		res, err := notionclient.ListTasks()
		if err != nil {
			log.Fatalf("Error adding task: %s", err)
		}

		for index, task := range res {
			fmt.Printf("%d. %s\n", index+1, task)
		}
	}
}

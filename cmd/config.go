package cmd

import (
	"fmt"
	"log"

	"github.com/jyisus/notioncli/internal/config"
	"github.com/spf13/cobra"
)

func InitConfigCommand(configPath string) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "config",
		Short: "Add configuration",
		Run:   runConfig(configPath),
	}
	listCmd.Flags().StringP("key", "k", "", "Notion API Key")
	listCmd.Flags().StringP("database", "d", "", "Notion Database ID")

	return listCmd
}

func runConfig(configPath string) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		notionKey, err := cmd.Flags().GetString("key")
		if err != nil {
			log.Println("Error getting key param")
		}
		databaseId, err := cmd.Flags().GetString("database")
		if err != nil {
			log.Println("Error getting database param")
		}
		config.GenerateFile(configPath, notionKey, databaseId)
		fmt.Println("Config file successfully created!")
	}
}

package cmd

import (
	"fmt"
	"github.com/jyisus/notioncli/usecase/config"
	"log"

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
			log.Printf("error getting key param: %s\n", err)
		}
		databaseId, err := cmd.Flags().GetString("database")
		if err != nil {
			log.Printf("error getting database param: %s\n", err)
		}

		err = config.GenerateFile(configPath, notionKey, databaseId)
		if err != nil {
			log.Printf("error generating config file: %s\n", err)
		}
		fmt.Println("Config file successfully created!")
	}
}

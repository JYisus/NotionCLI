package cmd

import (
	"github.com/spf13/cobra"
)

func InitRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "notioncli",
		Short: "Notion CLI (Command Line Interface)",
	}
}

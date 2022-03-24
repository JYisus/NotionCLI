package cmd

import "github.com/spf13/cobra"

type CobraFn func(cmd *cobra.Command, args []string)

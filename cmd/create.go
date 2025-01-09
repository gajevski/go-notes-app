package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"creation"},
	Short:   "create short",
	Long:    "create long",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("args: %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"creation"},
	Short:   "create short",
	Long:    "create long",
	Args:    cobra.ExactArgs(1),
	Run:     create,
}

func create(cmd *cobra.Command, args []string) {
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	len, err := file.WriteString(args[0] + "\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File name %s\n", file.Name())
	fmt.Printf("File length %d\n", len)
}

func init() {
	rootCmd.AddCommand(createCmd)
}

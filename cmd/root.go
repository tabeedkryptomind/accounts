package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
  Use:   "user",
  Short: "user service",
  Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use start to start a server")
		fmt.Println("Use -h to see the list of command")
	},
}


func Run() {
  rootCmd.AddCommand(accountsCmd)

  if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

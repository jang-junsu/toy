package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"toy/cmd/migrate"
)

var rootCmd = &cobra.Command{
	Use:   "toy",
	Short: "toy project command line tools",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(migrate.Cmd)
	if err := rootCmd.Execute(); err != nil {
		log.Printf("command error: %v", err)
		os.Exit(1)
	}
}

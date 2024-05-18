package migrate

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCmd = cobra.Command{
	Use:   "create-db",
	Short: "Create a database",
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func run() error {
	//config, err :=
	fmt.Println("create database")
	return nil
}

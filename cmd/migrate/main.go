package migrate

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func main() {

}

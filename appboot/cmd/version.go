package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of appboot",
	Long:  `All software has versions. This is appboot's`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: use ldflags?
		fmt.Println("appboot v0.0.1")
	},
}

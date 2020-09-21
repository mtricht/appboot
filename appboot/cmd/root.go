package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "appboot",
	Short: "appboot is an application bootstrapper",
	Long: `A cross-platform language-agnostic bootstrapper which keeps your application up to date. 
Were you expecting an application to launch? Make sure this binary is placed in the same directory as an 'app/appboot.json' file.
Complete documentation is available at https://github.com/mtricht/appboot`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

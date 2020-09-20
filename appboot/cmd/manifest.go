package cmd

import (
	"log"

	"github.com/mtricht/appboot/pkg/manifest"
	"github.com/spf13/cobra"
)

var directory string
var output string
var url string

func init() {
	rootCmd.AddCommand(manifestCmd)
	manifestCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Path to the application directory")
	manifestCmd.Flags().StringVarP(&output, "output", "o", "./manifest.json", "Output path for the generated manifest file")
	manifestCmd.Flags().StringVarP(&url, "url", "u", "", `URL prefix where the files will be hosted and can be downloaded 
	with a trailing slash`)
	manifestCmd.MarkFlagRequired("url")
}

var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "Generate a manifest file for your application",
	Long: `Generate a manifest file for your application. 
Example:

appboot manifest --directory=./my-app/ --output=./manifest.json --url=https://storage.googleapis.com/my-app/`,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Generating manifest file...")
		err := manifest.Generate(directory, output, url)
		if err != nil {
			return err
		}
		log.Printf("Saved file to '%s'...", output)
		return nil
	},
}

package main

import (
	"log"
	"os"

	"github.com/mtricht/appboot/appboot/cmd"
	"github.com/mtricht/appboot/pkg/launcher"
)

func main() {
	config := "./appboot.json"
	if os.Getenv("APPBOOT_JSON") != "" {
		config = os.Getenv("APPBOOT_JSON")
	}
	if isLauncher(config) {
		appboot, err := launcher.NewAppboot(config)
		if err != nil {
			log.Fatalln(err)
			return
		}
		log.Println("Checking for updates for " + appboot.Name)
		updateRequired, err := appboot.CheckForUpdates()
		if err != nil {
			log.Fatalln(err)
			return
		}
		if updateRequired {
			log.Println("Starting update for " + appboot.Name)
			appboot.Update()
			log.Println("Successfully updated " + appboot.Name)
		}
		// Re-read appboot.json in case it has been updated.
		appboot, err = launcher.NewAppboot(config)
		if err != nil {
			log.Fatalln(err)
			return
		}
		if err = appboot.RunCommand(); err != nil {
			log.Fatalln(err)
		}
		return
	}
	cmd.Execute()
}

func isLauncher(config string) bool {
	_, err := os.Stat(config)
	return !os.IsNotExist(err) && os.Getenv("NO_LAUNCHER") == ""
}

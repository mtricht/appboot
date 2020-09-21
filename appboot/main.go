package main

import (
	"log"
	"os"

	"github.com/mtricht/appboot/appboot/cmd"
	"github.com/mtricht/appboot/pkg/launcher"
)

func main() {
	config := "./app/appboot.json"
	if os.Getenv("APPBOOT_JSON") != "" {
		config = os.Getenv("APPBOOT_JSON")
	}
	if isLauncher(config) {
		appboot, err := launcher.NewAppboot(config)
		if err != nil {
			log.Fatalln(err)
			return
		}
		updateRequired, err := appboot.CheckForUpdates()
		if err != nil {
			log.Fatalln(err)
			return
		}
		if updateRequired {
			appboot.Update()
			log.Println("Updated!")
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

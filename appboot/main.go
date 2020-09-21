package main

import (
	"log"
	"os"

	"github.com/mtricht/appboot/appboot/cmd"
	"github.com/mtricht/appboot/pkg/launcher"
)

func main() {
	if isLauncher() {
		appboot, err := launcher.NewAppboot()
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
		appboot.RunCommand()
		return
	}
	cmd.Execute()
}

func isLauncher() bool {
	_, err := os.Stat("./app/appboot.json")
	return !os.IsNotExist(err)
}

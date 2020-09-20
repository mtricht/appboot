package main

import (
	"log"
	"os"

	"github.com/mtricht/appboot/appboot/cmd"
)

func main() {
	if isLauncher() {
		log.Fatal("appboot launcher has not yet been implemented.")
		return
	}
	cmd.Execute()
}

func isLauncher() bool {
	_, err := os.Stat("appboot.json")
	return !os.IsNotExist(err)
}

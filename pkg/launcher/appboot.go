package laucher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mtricht/appboot/pkg/manifest"
	"github.com/pkg/errors"
)

type Appboot struct {
	ManifestURL string `json:"manifest_url"`
	Command     string `json:"command"`
}

func NewAppboot() (*Appboot, error) {
	content, err := ioutil.ReadFile("app/appboot.json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to read app/appboot.json file")
	}
	var appboot Appboot
	err = json.Unmarshal(content, &appboot)
	if err != nil {
		return errors.Wrap(err, "unable to read app/appboot.json file")
	}
	return &appboot, nil
}

func (a Appboot) CheckForUpdates() error {
	manifest := manifest.Read(a.ManifestURL)
	fmt.Printf("%+v", manifest)
	return nil
}

func (a Appboot) Update() error {

}

func (a Appboot) RunCommand() error {

}

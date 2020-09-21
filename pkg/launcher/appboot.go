package launcher

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mtricht/appboot/pkg/manifest"
	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
)

type Appboot struct {
	ManifestURL   string `json:"manifest_url"`
	Command       string `json:"command"`
	remoteFiles   map[string]manifest.File
	localFiles    map[string]manifest.File
	filesToUpdate map[string]manifest.File
}

func NewAppboot() (*Appboot, error) {
	content, err := ioutil.ReadFile("app/appboot.json")
	if err != nil {
		return nil, errors.Wrap(err, "unable to read app/appboot.json file")
	}
	var appboot Appboot
	err = json.Unmarshal(content, &appboot)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read app/appboot.json file")
	}
	return &appboot, nil
}

func (a *Appboot) CheckForUpdates() (bool, error) {
	err := a.getRemoteFiles()
	if err != nil {
		return false, err
	}
	err = a.getLocalFiles()
	if err != nil {
		return false, err
	}
	a.filesToUpdate = make(map[string]manifest.File)
	for file, remoteFile := range a.remoteFiles {
		if localFile, ok := a.localFiles[file]; !ok || ok && (localFile.Checksum != remoteFile.Checksum || localFile.Size != remoteFile.Size) {
			a.filesToUpdate[file] = remoteFile
		}
	}
	return len(a.filesToUpdate) > 0, nil
}

func (a *Appboot) getRemoteFiles() error {
	entries, err := manifest.Read(a.ManifestURL)
	if err != nil {
		return err
	}
	a.remoteFiles = make(map[string]manifest.File)
	for _, entry := range entries {
		a.remoteFiles[entry.File] = entry
	}
	return nil
}

func (a *Appboot) getLocalFiles() error {
	a.localFiles = make(map[string]manifest.File)
	directory, err := filepath.Abs("./app")
	if err != nil {
		return err
	}
	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		hash, err := manifest.CalculateHash(path)
		if err != nil {
			return err
		}
		filename := manifest.GetFilename(path, directory)
		a.localFiles[filename] = manifest.File{
			File:     filename,
			Checksum: hash,
			Size:     info.Size(),
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (a Appboot) Update() error {
	for _, file := range a.filesToUpdate {
		req, _ := http.NewRequest("GET", file.URL, nil)
		resp, _ := http.DefaultClient.Do(req)
		defer resp.Body.Close()

		abs, _ := filepath.Abs("app/" + file.File)
		_ = os.MkdirAll(filepath.Dir(abs), os.ModePerm)
		f, _ := os.OpenFile("app/"+file.File, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		defer f.Close()

		bar := progressbar.DefaultBytes(
			resp.ContentLength,
			file.File,
		)
		io.Copy(io.MultiWriter(f, bar), resp.Body)
	}
	return nil
}

func (a Appboot) RunCommand() {
	log.Printf("Execute %s", a.Command)
}

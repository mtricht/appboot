package launcher

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mtricht/appboot/pkg/manifest"
	"github.com/schollz/progressbar/v3"
)

type Appboot struct {
	Name           string `json:"name"`
	ManifestURL    string `json:"manifest_url"`
	WindowsCommand string `json:"windows_command"`
	LinuxCommand   string `json:"linux_command"`
	DarwinCommand  string `json:"darwin_command"`
	directory      string
	remoteFiles    map[string]manifest.File
	localFiles     map[string]manifest.File
	filesToUpdate  map[string]manifest.File
}

func NewAppboot(config string) (*Appboot, error) {
	content, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}
	var appboot Appboot
	err = json.Unmarshal(content, &appboot)
	if err != nil {
		return nil, err
	}
	absolutePath, err := filepath.Abs(config)
	if err != nil {
		return nil, err
	}
	appboot.directory = filepath.Dir(absolutePath)
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
	err := filepath.Walk(a.directory, func(path string, info os.FileInfo, err error) error {
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
		filename := manifest.GetFilename(path, a.directory)
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
	// TODO: Check errors
	for _, file := range a.filesToUpdate {
		req, _ := http.NewRequest("GET", file.URL, nil)
		resp, _ := http.DefaultClient.Do(req)
		defer resp.Body.Close()

		path, _ := filepath.Abs(filepath.Join(a.directory, file.File))
		_ = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		f, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		defer f.Close()

		bar := progressbar.DefaultBytes(
			resp.ContentLength,
			file.File,
		)
		io.Copy(io.MultiWriter(f, bar), resp.Body)
	}
	return nil
}

func (a Appboot) RunCommand() error {
	commandString := ""
	if runtime.GOOS == "windows" {
		commandString = a.WindowsCommand
	} else if runtime.GOOS == "linux" {
		commandString = a.LinuxCommand
	} else if runtime.GOOS == "darwin" {
		commandString = a.DarwinCommand
	} else {
		panic("Unknown OS encountered")
	}
	commandSplit := strings.Fields(commandString)
	command := exec.Command(commandSplit[0], commandSplit[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

package manifest

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Generate generates a JSON manifest file to be used with appboot.
func Generate(source string, output string, URL string) error {
	path, err := filepath.Abs(source)
	if err != nil {
		return err
	}
	if URL == "" {
		return errors.New("URL may not be empty")
	}
	_, err = url.ParseRequestURI(URL)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("'%s' is not a valid URL", URL))
	}
	entries, err := getEntries(path, URL)
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		return fmt.Errorf("'%s' is an empty directory", path)
	}
	return createManifest(entries, output)
}

func getEntries(directory string, URL string) ([]File, error) {
	entries := make([]File, 0)
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		hash, err := CalculateHash(path)
		if err != nil {
			return err
		}
		file := GetFilename(path, directory)
		entries = append(entries, File{
			File:     file,
			Checksum: hash,
			URL:      URL + file,
			Size:     info.Size(),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// CalculateHash calculates a sha256 hash of a file
func CalculateHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", nil
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func GetFilename(path string, parentDirectory string) string {
	return strings.Replace(strings.ReplaceAll(strings.Replace(path, parentDirectory, "", 1), "\\", "/"), "/", "", 1)
}

func createManifest(entries []File, output string) error {
	bytes, err := json.Marshal(entries)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(output, bytes, 0644)
}

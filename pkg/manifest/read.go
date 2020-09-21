package manifest

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Read reads a manifest from a remote URL
func Read(URL string) ([]Entry, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to open manifest from '%s'", URL))
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to open manifest from '%s'", URL))
	}
	var manifest []Entry
	err = json.Unmarshal(content, &manifest)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to open manifest from '%s'", URL))
	}
	return manifest, nil
}
package manifest

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateSuccess(t *testing.T) {
	os.Remove("manifest.json")
	err := Generate("../../test-case/", "./manifest.json", "https://storage.googleapis.com/michael-personal/")
	if err != nil {
		t.Errorf("Generate() returned an error %+v", err)
	}
	content, err := ioutil.ReadFile("manifest.json")
	if err != nil {
		t.Errorf("Failed to read manifest.json: %+v", err)
	}
	text := string(content)
	expected := `[{"file":"appboot.json","checksum":"ee06c4345eb5584edfe81a54b9d635c40bea4ae5ae7d6fa390889c9aab22a942","url":"https://storage.googleapis.com/michael-personal/appboot.json","size":119},{"file":"application.sh","checksum":"8f57d0e772555a940b10f5cf8e840da14a5ff4dde085a757a2a6534d29dbe411","url":"https://storage.googleapis.com/michael-personal/application.sh","size":34},{"file":"lib/libappboot.so","checksum":"68be8bb7bc3eb5560f429b572a492ab3141bc985b74e5614b7ccaba485696379","url":"https://storage.googleapis.com/michael-personal/lib/libappboot.so","size":50}]`
	if expected != text {
		t.Error("manifest.json contents did not match expected")
	} else {
		os.Remove("manifest.json")
	}
}

// TODO: test unhappy paths

package manifest

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateSuccess(t *testing.T) {
	os.Remove("manifest.json")
	err := Generate("../../test-case/", "./manifest.json", "https://tricht.dev/appboot/")
	if err != nil {
		t.Errorf("Generate() returned an error %+v", err)
	}
	content, err := ioutil.ReadFile("manifest.json")
	if err != nil {
		t.Errorf("Failed to read manifest.json: %+v", err)
	}
	text := string(content)
	expected := `[{"file":"appboot.json","checksum":"1b6b6a1b85c541eb0c1cdbfc1e43afa03c58d2be9c6dee548c6cae58774cf883","url":"https://tricht.dev/appboot/appboot.json","size":117},{"file":"application.exe","checksum":"5957604c903e29b3f84fc57d4a870eab0f1b392d0c09ab48cacc310169399639","url":"https://tricht.dev/appboot/application.exe","size":32},{"file":"lib/libappboot.so","checksum":"68be8bb7bc3eb5560f429b572a492ab3141bc985b74e5614b7ccaba485696379","url":"https://tricht.dev/appboot/lib/libappboot.so","size":50}]`
	if expected != text {
		t.Error("manifest.json contents did not match expected")
	} else {
		os.Remove("manifest.json")
	}
}

// TODO: test unhappy paths

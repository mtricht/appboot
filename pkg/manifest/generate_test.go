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
	expected := `[{"file":"appboot.json","checksum":"940b196460644844be68a42fa56d01e9d1e35bb0e0b5272ab0a2981a02c5f702","url":"https://tricht.dev/appboot/appboot.json","size":75},{"file":"application.exe","checksum":"5957604c903e29b3f84fc57d4a870eab0f1b392d0c09ab48cacc310169399639","url":"https://tricht.dev/appboot/application.exe","size":32},{"file":"lib/libappboot.so","checksum":"68be8bb7bc3eb5560f429b572a492ab3141bc985b74e5614b7ccaba485696379","url":"https://tricht.dev/appboot/lib/libappboot.so","size":50}]`
	if expected != text {
		t.Error("manifest.json contents did not match expected")
	} else {
		os.Remove("manifest.json")
	}
}

// TODO: test unhappy paths

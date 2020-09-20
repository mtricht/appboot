package progress

import (
	"net/url"
	"os/exec"
)

type Launcher interface {
	NewLauncher(remote url.URL, command *exec.Cmd)
	Close()
}

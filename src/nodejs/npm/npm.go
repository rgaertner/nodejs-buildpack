package npm

import (
	"io"
	"github.com/cloudfoundry/libbuildpack"
)

type Command interface {
	Execute(dir string, stdout io.Writer, stderr io.Writer, program string, args ...string) error
}

type NPM struct {
	BuildDir string
	Command  Command
	Log      *libbuildpack.Logger
}

func (n *NPM) Build() error {
	args := []string{"'skipping'"}
	return n.Command.Execute(n.BuildDir, n.Log.Output(), n.Log.Output(), "echo", args...)
}

func (n *NPM) Rebuild() error {
	args := []string{"'skipping'"}
	return n.Command.Execute(n.BuildDir, n.Log.Output(), n.Log.Output(), "echo", args...)
}

func (n *NPM) doBuild() (bool, string, error) {
	args := []string{"'skipping'"}
	return n.Command.Execute(n.BuildDir, n.Log.Output(), n.Log.Output(), "echo", args...)
}

package npm

import (
	"io"
	"path/filepath"
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
	return n.Command.Execute(n.BuildDir, n.Log.Output(), n.Log.Output(), "echo", "'skipping'")
}

func (n *NPM) Rebuild() error {
	return n.Command.Execute(n.BuildDir, n.Log.Output(), n.Log.Output(), "echo", "'skipping'")
}

func (n *NPM) doBuild() (bool, string, error) {
	return n.Command.Execute(n.BuildDir, n.Log.Output(), n.Log.Output(), "echo", "'skipping'")
}

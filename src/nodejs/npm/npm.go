package npm

import (
	"io"
	"os"
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
	doBuild, source, err := n.doBuild()
	if err != nil {
		return err
	}
	if !doBuild {
		return nil
	}

	n.Log.Info("Installing node modules (%s)", source)
	npmArgs := []string{"install", "--unsafe-perm", "--userconfig", filepath.Join(n.BuildDir, ".npmrc"), "--cache", filepath.Join(n.BuildDir, ".npm")}
	return n.Command.Execute(n.BuildDir, n.Log.Output(), n.Log.Output(), "npm", npmArgs...)
}

func (n *NPM) Rebuild() error {
	return n.Command.Execute(n.BuildDir, n.Log.Output(), n.Log.Output(), "echo", []string{"'skipping'"})
}

func (n *NPM) doBuild() (bool, string, error) {
	return n.Command.Execute(n.BuildDir, n.Log.Output(), n.Log.Output(), "echo", []string{"'skipping'"})
}

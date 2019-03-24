package datadir

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rancher/norman/pkg/resolvehome"
)

const (
	DefaultDataDir     = "/var/lib/erdii/k3s"
	DefaultHomeDataDir = "${HOME}/.erdii/k3s"
	HomeConfig         = "${HOME}/.kube/k3s.yaml"
	GlobalConfig       = "/etc/erdii/k3s/k3s.yaml"
)

func Resolve(dataDir string) (string, error) {
	if dataDir == "" {
		if os.Getuid() == 0 {
			dataDir = DefaultDataDir
		} else {
			dataDir = DefaultHomeDataDir
		}
	}

	dataDir, err := resolvehome.Resolve(dataDir)
	if err != nil {
		return "", errors.Wrapf(err, "resolving %s", dataDir)
	}

	return dataDir, nil
}

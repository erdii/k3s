package main

import (
	"os"
	"path/filepath"

	"github.com/docker/docker/pkg/reexec"
	crictl2 "github.com/kubernetes-sigs/cri-tools/cmd/crictl"
	"github.com/erdii/k3s/pkg/cli/agent"
	"github.com/erdii/k3s/pkg/cli/cmds"
	"github.com/erdii/k3s/pkg/cli/crictl"
	"github.com/erdii/k3s/pkg/cli/kubectl"
	"github.com/erdii/k3s/pkg/cli/server"
	"github.com/erdii/k3s/pkg/containerd"
	kubectl2 "github.com/erdii/k3s/pkg/kubectl"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	reexec.Register("containerd", containerd.Main)
	reexec.Register("kubectl", kubectl2.Main)
	reexec.Register("crictl", crictl2.Main)
}

func main() {
	cmd := os.Args[0]
	os.Args[0] = filepath.Base(os.Args[0])
	if reexec.Init() {
		return
	}
	os.Args[0] = cmd

	app := cmds.NewApp()
	app.Commands = []cli.Command{
		cmds.NewServerCommand(server.Run),
		cmds.NewAgentCommand(agent.Run),
		cmds.NewKubectlCommand(kubectl.Run),
		cmds.NewCRICTL(crictl.Run),
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

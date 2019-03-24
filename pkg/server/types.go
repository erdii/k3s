package server

import (
	"github.com/erdii/k3s/pkg/daemons/config"
	"github.com/rancher/norman/pkg/dynamiclistener"
)

type Config struct {
	DisableAgent     bool
	DisableServiceLB bool
	TLSConfig        dynamiclistener.UserConfig
	ControlConfig    config.Control
}

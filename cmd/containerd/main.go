package main

import (
	"github.com/erdii/k3s/pkg/containerd"
	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	containerd.Main()
}

package main

import (
	"synck8s/pkgs/config"
	"synck8s/pkgs/dstk8s"
	"synck8s/pkgs/srck8s"
)

func main() {
	var configuration config.Configurations
	config.ReadConfs(&configuration, "config")

	//Set kubeconfig
	srck8s.Kubeconfig = configuration.Kubeconfig.Source
	dstk8s.Kubeconfig = configuration.Kubeconfig.Destination

	for _, deploy := range configuration.Deployments {
		deploy.Sync()

	}

}

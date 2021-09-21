package config

import (
	"log"

	"synck8s/pkgs/deploy"

	"github.com/spf13/viper"
)

//Configurations matchs data from yaml to Configurations type
type Configurations struct {
	Kubeconfig  KubeConfigurations
	Deployments []deploy.Deployment
}

//ServerConfigurations exporter listen port number
type KubeConfigurations struct {
	Source      string
	Destination string
}

//ReadConfs Read ./config.yml and import it to confs *Configurations
func ReadConfs(confs *Configurations, confFile string) (err error) {

	viper.SetConfigName(confFile)

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err = viper.ReadInConfig(); err != nil {
		log.Print("Error reading config file", err)
		return
	}

	err = viper.Unmarshal(&confs)
	if err != nil {
		log.Print("Unable to decode into struct", err)
		return
	}

	return

}

package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	configFile = "./config/config.yml"
)

type config struct {
	DriverName  string       `yaml:"driverName"`
	DataSourceName string    `yaml:"dataSourceName"`
}

var (
	c = config{}

	DriverName = c.DriverName

	DataSourceName = c.DataSourceName
)

func Configure() {
	configYaml, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(configYaml, &c); err != nil {
		panic(err)
	}
	DriverName = c.DriverName

	DataSourceName = c.DataSourceName
}

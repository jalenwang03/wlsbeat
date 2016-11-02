// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period    time.Duration    `config:"period"`
	Instances []InstanceConfig `config:"instances"`
}

//type InstancesConfig struct {
//	Instance InstanceConfig `config:"instance`
//}

type InstanceConfig struct {
	Host     string `config:"Host"`
	Port     string `config:"Port"`
	Username string `config:"Username"`
	Password string `config:"Password"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
}

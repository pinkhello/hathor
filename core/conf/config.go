package conf

import (
	"flag"
	"github.com/pinkhello/hathor/core/pkg/adapter"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"os"
)

type (
	Config struct {
		Api rest.RestConf `json:"api"`
		Log adapter.LogConf `json:"log"`
	}
	Configs struct {
		Dev Config `json:"dev"`
		Test Config `json:"test"`
		Alpha Config `json:"alpha"`
		Prod Config `json:"prod"`
	}
)

func (ca *Configs) GetEnvConfig(env string) Config {
	if env == "prod" {
		return ca.Prod
	}
	if env == "alpha"{
		return ca.Alpha
	}
	if env == "test" {
		return ca.Test
	}
	return ca.Dev
}

func LoadConfig(env string) Config {
	pwd, _ := os.Getwd()
	var configFile string
	if flag.Lookup("c") == nil {
		flag.StringVar(&configFile, "c", pwd + "/core/conf/config.yml", "config.yml")
	}
	configFile = flag.Lookup("c").Value.(flag.Getter).Get().(string)
	flag.Parse()
	var c Configs
	conf.MustLoad(configFile, &c)
	return c.GetEnvConfig(env)
}
package repository_config

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/hengkyawijaya/simple-go/constant"
	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTPServer HTTPServerConfig `yaml:"http_server"`
	GRPCServer GRPCServerConfig `yaml:"grpc_server"`
	CronServer CronServerConfig `yaml:"cron_server"`
	Repository RepositoryConfig `yaml:"repository"`
}

type HTTPServerConfig struct {
	Port string
}

type GRPCServerConfig struct {
	Port string
}

type CronServerConfig struct {
	CronHello CronHello `yaml:"cron_hello"`
}

type CronHello struct {
	Schedule string
}

type RepositoryConfig struct {
	IPInfo IPInfoRepositoryConfig `yaml:"ip_info"`
}

type IPInfoRepositoryConfig struct {
	Endpoint string
}

func NewConfig(ctx context.Context) *Config {
	env := ctx.Value(constant.EnvContextKey)
	filePath := fmt.Sprintf("./files/config/main.%s.yaml", env)
	fmt.Println("read config from", filePath)

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return &config
}

func (c *Config) ReadConfig() *Config {
	return c
}

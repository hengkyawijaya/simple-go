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
	IPInfo   IPInfoRepositoryConfig `yaml:"ip_info"`
	Database DatabaseConfig         `yaml:"database"`
}

type IPInfoRepositoryConfig struct {
	Endpoint string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func NewConfig(ctx context.Context) *Config {
	env := ctx.Value(constant.EnvContextKey)
	filePath := fmt.Sprintf("./files/config/main.%s.yaml", env)
	config := readConfig(filePath)

	filePathSecret := fmt.Sprintf("./files/config/secret.%s.yaml", env)
	configSecret := readConfig(filePathSecret)

	// override config with secret
	config.Repository.Database = configSecret.Repository.Database

	return &config
}

func readConfig(filePath string) Config {
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
	return config
}

func (c *Config) ReadConfig() *Config {
	return c
}

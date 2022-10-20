package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		RestClient RestClient `yaml:"rest_client"`
		Storage    Storage    `yaml:"storage"`
	}
	RestClient struct {
		URL string `yaml:"url"`
	}
	Storage struct {
		Connection string `yaml:"connection"`
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Database   string `yaml:"database"`
		Collection string `yaml:"collection"`
	}
)

const filePathFormat = "%s/api/cmd/config/%s.yml"

var (
	ymlConf Config
	once    sync.Once
)

func Get() Config {
	once.Do(func() {
		ymlConf = Config{}
		readFromYML(&ymlConf)
	})
	return ymlConf
}

func readFromYML(conf *Config) {
	ymlFile, err := os.ReadFile(getPath("config"))
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(ymlFile, conf); err != nil {
		log.Fatal(err)
	}
}

func getPath(fileName string) string {
	basePath, _ := os.Getwd()
	return fmt.Sprintf(filePathFormat, basePath, fileName)
}

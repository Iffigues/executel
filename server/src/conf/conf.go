package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Addr  string `yaml:"addr"`
	Port  string `yaml:"port"`
	Types string `yaml:"types"`
}

func NewConf(f string) (c *Conf) {
	c = new(Conf)
	file, err := os.Open("data.yaml")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	err = yaml.NewDecoder(file).Decode(c)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}
	return
}

func (c *Conf) GetAddr() string {
	return c.Addr
}

func (c *Conf) GetPort() string {
	return c.Port
}

func (c *Conf) GetTypes() string {
	return c.Types
}

package db

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"

	yaml "gopkg.in/yaml.v1"
)

type Configs map[string]*Config

func (cs Configs) Open() (*gorm.DB, error) {
	config, ok := cs["development"]
	if !ok {
		return nil, nil
	}
	return config.Open()
}

type Config struct {
	Datasource string `yaml:"datasource"`
}

func (c *Config) DSN() string {
	return c.Datasource
}

func (c *Config) Open() (*gorm.DB, error) {
	return gorm.Open("mysql", c.DSN())
}

func NewConfigsFromFile() (Configs, error) {
	f, err := os.Open("dbconfig.yml")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewConfigs(f)
}

func NewConfigs(r io.Reader) (Configs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var configs Configs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

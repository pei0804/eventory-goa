package database

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v1"
)

// Configs have configuration for each environment.
type Configs map[string]*Config

// Open creates connection between database for each environment.
func (cs Configs) Open(env string) (*gorm.DB, error) {
	config, ok := cs[env]
	if !ok {
		return nil, nil
	}
	return config.Open()
}

// Config is a database configuration.
// It's save as sql-migrate schema style.
//
// see also: https://github.com/rubenv/sql-migrate
type Config struct {
	Datasource string `yaml:"datasource"`
}

// DSN returns data source name configured.
func (c *Config) DSN() string {
	return c.Datasource
}

// Open connets database.
// NOTE: supports mysql only.
func (c *Config) Open() (*gorm.DB, error) {
	return gorm.Open("mysql", c.DSN())
}

// NewConfigsFromFile reads settings from file.
func NewConfigsFromFile(path string) (Configs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewConfigs(f)
}

// NewConfigs reads configs from io.Reader.
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

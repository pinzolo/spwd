package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	"github.com/pinzolo/xdgdir"
)

var app = xdgdir.NewApp("spwd")

// Config is configurations holder of spwd app.
type Config struct {
	// IdentityFile is file path of ssl secret file.
	IdentityFile string `yaml:"identity_file"`
	// DataFile is file path of storing encrypted passwords.
	DataFile string `yaml:"data_file"`
}

// GetConfig return merged configuration.
func GetConfig() (Config, error) {
	dcfg, err := DefaultConfig()
	if err != nil {
		return dcfg, err
	}
	fcfg, ok := FileConfig()
	if ok {
		return dcfg.Merge(fcfg), nil
	}
	return dcfg, nil
}

// Merge config values and returns new Config.
func (cfg Config) Merge(other Config) Config {
	newCfg := Config{
		IdentityFile: cfg.IdentityFile,
		DataFile:     cfg.DataFile,
	}
	if other.IdentityFile != "" {
		newCfg.IdentityFile = other.IdentityFile
	}
	if other.DataFile != "" {
		newCfg.DataFile = other.DataFile
	}
	return newCfg
}

// DefaultConfig return sefault configuration.
func DefaultConfig() (Config, error) {
	df, err := app.DataFile("data.yml")
	if err != nil {
		return Config{}, err
	}
	return Config{
		IdentityFile: filepath.Join(homeDir(), ".ssh", "id_rsa"),
		DataFile:     df,
	}, nil
}

// FileConfig return configuration of config file.
// Second bool return value is existence of config file.
func FileConfig() (Config, bool) {
	cf, err := app.ConfigFile("config.yml")
	if err != nil {
		return Config{}, false
	}
	if _, err = os.Stat(cf); err != nil {
		return Config{}, false
	}
	p, err := ioutil.ReadFile(cf)
	if err != nil {
		return Config{}, false
	}
	var cfg Config
	if err = yaml.Unmarshal(p, &cfg); err != nil {
		return Config{}, false
	}
	return cfg, true
}

func homeDir() string {
	home := os.Getenv("HOME")
	if home != "" {
		return home
	}
	return os.Getenv("USERPROFILE")
}

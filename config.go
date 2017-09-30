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
	// KeyFile is file path of secret file.
	KeyFile string `yaml:"key_file"`
	// DataFile is file path of storing encrypted passwords.
	DataFile string `yaml:"data_file"`
	// FilteringCommand is command for filtering in search subcommand.
	FilteringCommand string `yaml:"filtering_command"`
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
		KeyFile:          cfg.KeyFile,
		DataFile:         cfg.DataFile,
		FilteringCommand: cfg.FilteringCommand,
	}
	if other.KeyFile != "" {
		newCfg.KeyFile = other.KeyFile
	}
	if other.DataFile != "" {
		newCfg.DataFile = other.DataFile
	}
	if other.FilteringCommand != "" {
		newCfg.FilteringCommand = other.FilteringCommand
	}
	return newCfg
}

// DefaultConfig return sefault configuration.
func DefaultConfig() (Config, error) {
	df, err := app.DataFile("data.dat")
	if err != nil {
		return Config{}, err
	}
	return Config{
		KeyFile:          filepath.Join(homeDir(), ".ssh", "id_rsa"),
		DataFile:         df,
		FilteringCommand: "peco",
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

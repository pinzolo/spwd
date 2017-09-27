package main

import (
	"os"
	"path/filepath"
)

// Initialize makes data directory if not exist.
func Initialize(cfg Config) error {
	dp := filepath.Dir(cfg.DataFile)
	if _, err := os.Stat(dp); err != nil {
		err = os.MkdirAll(dp, 0755)
		return err
	}
	return nil
}

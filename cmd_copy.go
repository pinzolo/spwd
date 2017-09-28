package main

import (
	"errors"
	"fmt"

	"github.com/atotto/clipboard"
)

var cmdCopy = &Command{
	Run:       runCopy,
	UsageLine: "copy NAME",
	Short:     "Copy password to clipboard",
	Long:      `Find password and copy to clipboard.`,
}

func runCopy(ctx context, args []string) error {
	if len(args) == 0 {
		return errors.New("item name required")
	}
	cfg, err := GetConfig()
	if err != nil {
		return err
	}
	Initialize(cfg)
	key, err := GetKey(cfg.KeyFile)
	if err != nil {
		return err
	}
	is, err := LoadItems(key, cfg.DataFile)
	if err != nil {
		return err
	}
	it := is.Find(args[0])
	if it == nil {
		return fmt.Errorf("item not found: %s", args[0])
	}
	clipboard.WriteAll(it.Password)
	fmt.Fprintln(ctx.out, fmt.Sprintf("password of '%s' copy to clipboard successfully", it.Name))
	return nil
}

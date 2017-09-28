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
	is, err := LoadItems(cfg.DataFile)
	if err != nil {
		return err
	}
	it := is.Find(args[0])
	if it == nil {
		return fmt.Errorf("item not found: %s", args[0])
	}

	key, err := GetKey(cfg.KeyFile)
	if err != nil {
		return err
	}
	pwd, err := Decrypt(key, it.Encrypted)
	if err != nil {
		return err
	}
	clipboard.WriteAll(pwd)
	fmt.Fprintln(ctx.out, fmt.Sprintf("password of '%s' copy to clipboard successfully", it.Name))
	return nil
}

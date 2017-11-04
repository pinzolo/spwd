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
		return errors.New("item name is required")
	}
	cfg, err := GetConfig()
	if err != nil {
		return err
	}
	Initialize(cfg)
	is, err := LoadItemsWithConfig(cfg)
	if err != nil {
		return err
	}
	if is.HasMaster() {
		if err = confirmMasterPassword(is.Master()); err != nil {
			return err
		}
	}
	it := is.Find(args[0])
	if it == nil || it.Master {
		return fmt.Errorf("item not found: %s", args[0])
	}
	clipboard.WriteAll(it.Password)
	PrintSuccess(ctx.out, "password of '%s' is copied to clipboard successfully", it.Name)
	return nil
}

func confirmMasterPassword(it *Item) error {
	pwd, err := scanPassword("Master password: ")
	if err != nil {
		return err
	}
	if it.Password != pwd {
		return errMasterPasswordNotMatch
	}
	return nil
}

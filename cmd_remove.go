package main

import (
	"errors"
	"fmt"
)

var cmdRemove = &Command{
	Run:       runRemove,
	UsageLine: "remove NAME",
	Short:     "Remove saved password",
	Long:      `Remove saved password by input name.`,
}

func runRemove(ctx context, args []string) error {
	if len(args) == 0 {
		return errors.New("item name is required")
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

	name := args[0]
	fit := is.Find(name)
	if fit == nil {
		return fmt.Errorf("item not found: %s", name)
	}

	nis := Items([]Item{})
	for _, it := range is {
		if it.Name != fit.Name {
			nis = append(nis, it)
		}
	}
	nis.Save(key, cfg.DataFile)
	PrintSuccess(ctx.out, "password of '%s' is removed successfully", name)
	return nil
}

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

	name := args[0]
	fit := is.Find(name)
	if fit == nil {
		return fmt.Errorf("item not found: %s", name)
	}

	nis := Items(make([]Item, len(is)-1))
	for _, it := range is {
		if it.Name != fit.Name {
			nis = append(nis, it)
		}
	}
	nis.Save(cfg.DataFile)
	fmt.Fprintln(ctx.out, fmt.Sprintf("password of '%s' is removed successfully", name))
	return nil
}

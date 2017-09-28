package main

import (
	"errors"
	"fmt"
)

var cmdMigrate = &Command{
	Run:       runMigrate,
	UsageLine: "migrate",
	Short:     "Regenerate data file",
	Long:      `Regenerate data file with new key file.`,
}

// MigrateFileName is file name that created with migrate subcommand.
const MigrateFileName = "spwd-migrated.yml"

func runMigrate(ctx context, args []string) error {
	if len(args) == 0 {
		return errors.New("new key file required")
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

	if len(is) == 0 {
		fmt.Fprintln(ctx.out, "no password.")
		return nil
	}

	key, err := GetKey(cfg.KeyFile)
	if err != nil {
		return err
	}
	nkey, err := GetKey(args[0])
	if err != nil {
		return err
	}

	nis := Items(make([]Item, len(is)))
	var pwd, enc string
	for i, it := range is {
		pwd, err = Decrypt(key, it.Encrypted)
		if err != nil {
			return err
		}
		enc, err = Encrypt(nkey, pwd)
		if err != nil {
			return err
		}
		nis[i] = NewItem(it.Name, it.Description, enc)
	}

	err = nis.Save(MigrateFileName)
	if err != nil {
		return err
	}
	fmt.Fprintln(ctx.out, fmt.Sprintf("new data file saved as %s successfully", MigrateFileName))
	return nil
}

package main

import (
	"errors"
	"fmt"
)

var cmdMigrate = &Command{
	Run:       runMigrate,
	UsageLine: "migrate NEW_KEY_FILE",
	Short:     "Regenerate data file",
	Long:      `Regenerate data file with new key file.`,
}

// MigrateFileName is file name that created with migrate subcommand.
const MigrateFileName = "spwd-migrated.dat"

func runMigrate(ctx context, args []string) error {
	if len(args) == 0 {
		return errors.New("new key file is required")
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

	if is.HasMaster() && cfg.IsProtective(ctx.cmdName) {
		if err = confirmMasterPassword(is.Master()); err != nil {
			return err
		}
	}

	if len(is) == 0 {
		fmt.Fprintln(ctx.out, "no password.")
		return nil
	}

	nkey, err := GetKey(args[0])
	if err != nil {
		return err
	}
	err = is.Save(nkey, MigrateFileName)
	if err != nil {
		return err
	}
	PrintSuccess(ctx.out, "new data file saved as %s successfully", MigrateFileName)
	return nil
}

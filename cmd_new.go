package main

import (
	"errors"
	"fmt"
)

var cmdNew = &Command{
	Run:       runNew,
	UsageLine: "new",
	Short:     "Register new password item interactively",
	Long: `Register new password item to data file interactively.
If input name already exists, you can update it.
`,
}

func runNew(ctx context, args []string) error {
	cfg, err := GetConfig()
	if err != nil {
		return err
	}

	err = Initialize(cfg)
	if err != nil {
		return err
	}

	key, err := GetKey(cfg.KeyFile)
	if err != nil {
		return err
	}
	is, err := LoadItems(key, cfg.DataFile)
	if err != nil {
		return err
	}

	if is.HasMaster() && cfg.IsProtective(ctx.cmdName) {
		if err = confirmMasterPassword(is.Master()); err != nil {
			return err
		}
	}

	name, desc, pwd, err := scan()
	if err != nil {
		return err
	}
	nit := NewItem(name, desc, pwd)
	if it := is.Find(name); it != nil {
		b, berr := scanBool(fmt.Sprintf("item '%s' already exists, update? [y/N]: ", name))
		if berr != nil {
			return berr
		}
		if !b {
			return nil
		}
		is = is.Update(nit)
	} else {
		is = append(is, nit)
	}

	err = is.Save(key, cfg.DataFile)
	if err != nil {
		return err
	}
	PrintSuccess(ctx.out, "password of '%s' is saved successfully", name)
	return nil
}

func scan() (name string, desc string, pwd string, err error) {
	name, err = scanText("Name: ")
	if name == "" {
		err = errors.New("name is required")
		return
	}
	desc, err = scanText("Description: ")
	pwd, err = scanPassword("Password: ")
	if pwd == "" {
		err = errors.New("password is required")
	}
	fmt.Println()
	return
}

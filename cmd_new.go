package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
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
	Initialize(cfg)
	is, err := LoadItems(cfg.DataFile)
	if err != nil {
		return err
	}

	name, desc, pwd, err := scan()
	if err != nil {
		return err
	}
	key, err := GetKey(cfg.KeyFile)
	if err != nil {
		return err
	}
	enc, err := Encrypt(key, string(pwd))
	nit := NewItem(name, desc, Encode(enc))
	if err != nil {
		return err
	}
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

	err = is.Save(cfg.DataFile)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("item '%s' saved!!", name))
	return nil
}

func scan() (name string, desc string, pwd string, err error) {
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("Name: ")
	in.Scan()
	if err = in.Err(); err != nil {
		return
	}
	name = in.Text()
	fmt.Print("Description: ")
	in.Scan()
	if err = in.Err(); err != nil {
		return
	}
	desc = in.Text()
	fmt.Print("Password: ")
	var p []byte
	p, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return
	}
	pwd = string(p)
	fmt.Println()
	return
}

func scanBool(prompt string) (bool, error) {
	fmt.Print(prompt)
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		return false, err
	}

	return strings.ToLower(in.Text()) == "y", nil
}

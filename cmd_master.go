package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var cmdMaster = &Command{
	Run:       runMaster,
	UsageLine: "master",
	Short:     "Handle master password item interactively",
	Long: `Register master password item file interactively.
If master password already exists, you can update or delete it.
`,
}

var errMasterPasswordNotMatch = errors.New("master password is not matched")

func runMaster(ctx context, args []string) error {
	rand.Seed(time.Now().UnixNano())
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
	var msg string
	if is.HasMaster() {
		method, err := scanText("master password already exists, update or delete? [u(pdate)/D(ELETE)]: ")
		if err != nil {
			return err
		}
		if isInvalidMasterMethod(method) {
			return fmt.Errorf("invalid input: %s", method)
		}
		pwd, err := scanPassword("Current master password: ")
		if err != nil {
			return err
		}
		mst := is.Master()
		if mst.Password != pwd {
			return errMasterPasswordNotMatch
		}
		if isDeleteMasterMethod(method) {
			is = is.Remove(mst.Name)
			msg = "master password is deleted successfully"
		} else {
			newPwd, err := scanPassword("New master password: ")
			if err != nil {
				return err
			}
			is = is.Update(NewMasterItem(mst.Name, newPwd))
			msg = "master password is updated successfully"
		}
	} else {
		pwd, err := scanPassword("Master password: ")
		if err != nil {
			return err
		}
		is = append(is, NewMasterItem(genMasterItemName(is), pwd))
		msg = "master password is saved successfully"
	}
	err = is.Save(key, cfg.DataFile)
	if err != nil {
		return err
	}
	PrintSuccess(ctx.out, msg)
	return nil
}

func genMasterItemName(is Items) string {
	name := Encode([]byte(strconv.Itoa(rand.Int())))
	if is.Find(name) == nil {
		return name
	}
	return genMasterItemName(is)
}

func isInvalidMasterMethod(method string) bool {
	if isUpdateMasterMethod(method) {
		return false
	}
	if isDeleteMasterMethod(method) {
		return false
	}
	return true
}

func isUpdateMasterMethod(method string) bool {
	lmethod := strings.ToLower(method)
	return lmethod == "u" || lmethod == "update"
}

func isDeleteMasterMethod(method string) bool {
	return method == "D" || method == "DELETE"
}

package main

import (
	"fmt"

	"github.com/olekukonko/tablewriter"
)

var cmdList = &Command{
	Run:       runList,
	UsageLine: "list",
	Short:     "List name and description",
	Long:      `List name and description of saved passwords.`,
}

func runList(ctx context, args []string) error {
	cfg, err := GetConfig()
	if err != nil {
		return err
	}

	err = Initialize(cfg)
	if err != nil {
		return err
	}

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
		_, err = fmt.Fprintln(ctx.out, "no password.")
		return err
	}

	tw := tablewriter.NewWriter(ctx.out)
	tw.SetHeader([]string{"Name", "Description"})
	tw.SetColumnSeparator("")
	tw.SetCenterSeparator(" ")
	tw.SetBorder(false)
	tw.AppendBulk(is.ToDataTable())
	tw.Render()
	return nil
}

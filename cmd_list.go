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

	if len(is) == 0 {
		fmt.Fprintln(ctx.out, "no password.")
		return nil
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

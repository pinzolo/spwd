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
	key, err := GetKey(cfg.KeyFile)
	if err != nil {
		return err
	}
	is, err := LoadItems(key, cfg.DataFile)
	if err != nil {
		return err
	}

	if len(is) == 0 {
		fmt.Fprintln(ctx.out, "no password.")
		return nil
	}

	data := make([][]string, len(is))
	for i, it := range is {
		data[i] = []string{it.Name, it.Description}
	}

	tw := tablewriter.NewWriter(ctx.out)
	tw.SetHeader([]string{"Name", "Description"})
	tw.SetColumnSeparator("")
	tw.SetCenterSeparator(" ")
	tw.SetBorder(false)
	tw.AppendBulk(data)
	tw.Render()
	return nil
}

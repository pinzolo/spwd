package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/olekukonko/tablewriter"
)

var cmdSearch = &Command{
	Run:       runSearch,
	UsageLine: "search",
	Short:     "Search keywords",
	Long:      `Search keywords interactively (default filtering tool: peco)`,
}

func runSearch(ctx context, args []string) error {
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

	buf := &bytes.Buffer{}
	err = runCommand(cfg.FilteringCommand, strings.NewReader(filteringText(is)), buf)
	if err != nil {
		return err
	}
	name := strings.TrimSpace(strings.Split(buf.String(), "|")[0])
	return runCopy(ctx, []string{name})
}

func filteringText(is Items) string {
	buf := &bytes.Buffer{}
	tw := tablewriter.NewWriter(buf)
	tw.SetColumnSeparator("|")
	tw.SetBorder(false)
	tw.AppendBulk(is.ToDataTable())
	tw.Render()
	return buf.String()
}

func runCommand(cmdText string, r io.Reader, w io.Writer) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", cmdText)
	} else {
		cmd = exec.Command("sh", "-c", cmdText)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd.Run()
}

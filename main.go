package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

// A Command is an implementation of a spwd command
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(ctx context, args []string) error

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'spwd help' output.
	Short string

	// Long is the long message shown in the 'spwd help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet
}

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// Usage prints usage of command
func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}

// Commands lists the available commands and help topics.
// The order here is the order in which they are printed by 'spwd help'.
var commands = []*Command{
	cmdCopy,
	cmdList,
	cmdMigrate,
	cmdNew,
	cmdRemove,
	cmdSearch,
	cmdVersion,
}

func main() {
	flag.Usage = usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			cmd.Flag.Usage = func() { cmd.Usage() }

			cmd.Flag.Parse(args[1:])
			args = cmd.Flag.Args()

			ow := bufio.NewWriter(os.Stdout)
			err := cmd.Run(newContext(ow), args)
			code := 0
			if err != nil {
				PrintError(os.Stderr, err)
				code = 2
			}
			ow.Flush()
			os.Exit(code)
		}
	}

	fmt.Fprintf(os.Stderr, "spwd: unknown subcommand %q\nRun ' spwd help' for usage.\n", args[0])
	os.Exit(2)
}

var usageTemplate = `spwd is secret file based password management tool.
Input password is encrypted with AES-256 using your secret file.

Usage:

	spwd command [arguments]

The commands are:
{{range .}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "spwd help [command]" for more information about a command.

`

var helpTemplate = `usage: spwd {{.UsageLine}}

{{.Long | trim}}
`

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func printUsage(w io.Writer) {
	bw := bufio.NewWriter(w)
	tmpl(bw, usageTemplate, commands)
	bw.Flush()
}

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

// help implements the 'help' command.
func help(args []string) {
	if len(args) == 0 {
		printUsage(os.Stdout)
		// not exit 2: succeeded at 'spwd help'.
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: spwd help command\n\nToo many arguments given.\n")
		os.Exit(2) // failed at 'spwd help'
	}

	arg := args[0]

	for _, cmd := range commands {
		if cmd.Name() == arg {
			tmpl(os.Stdout, helpTemplate, cmd)
			// not exit 2: succeeded at 'spwd help cmd'.
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic %#q.  Run 'spwd help'.\n", arg)
	os.Exit(2) // failed at 'spwd help cmd'
}

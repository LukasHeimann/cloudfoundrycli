//go:build go1.13
// +build go1.13

package main

import (
	"fmt"
	"os"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/cmd"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/common"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/command_parser"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/configv3"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/panichandler"
	plugin_util "github.com/LukasHeimann/cloudfoundrycli/v8/util/plugin"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/ui"
)

func main() {
	var exitCode int
	defer panichandler.HandlePanic()

	config, err := configv3.GetCFConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error: %s\n", err.Error())
		os.Exit(1)
	}

	commandUI, err := ui.NewUI(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error: %s\n", err.Error())
		os.Exit(1)
	}

	p, err := command_parser.NewCommandParser()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error: %s\n", err.Error())
		os.Exit(1)
	}

	exitCode, err = p.ParseCommandFromArgs(commandUI, os.Args[1:])
	if err == nil {
		os.Exit(exitCode)
	}

	if unknownCommandError, ok := err.(command_parser.UnknownCommandError); ok {
		plugin, commandIsPlugin := plugin_util.IsPluginCommand(os.Args[1:])

		switch {
		case commandIsPlugin:
			err = plugin_util.RunPlugin(plugin)
			if err != nil {
				exitCode = 1
			}

		case common.ShouldFallbackToLegacy:
			cmd.Main(os.Getenv("CF_TRACE"), os.Args)
			//NOT REACHED, legacy main will exit the process

		default:
			unknownCommandError.Suggest(plugin_util.PluginCommandNames())
			fmt.Fprintf(os.Stderr, "%s\n", unknownCommandError.Error())
			os.Exit(1)
		}
	}

	os.Exit(exitCode)
}

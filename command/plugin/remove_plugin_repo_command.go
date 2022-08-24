package plugin

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/command"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/flag"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/translatableerror"
)

type RemovePluginRepoCommand struct {
	RequiredArgs    flag.PluginRepoName `positional-args:"yes"`
	usage           interface{}         `usage:"CF_NAME remove-plugin-repo REPO_NAME\n\nEXAMPLES:\n   CF_NAME remove-plugin-repo PrivateRepo"`
	relatedCommands interface{}         `related_commands:"list-plugin-repos"`
}

func (RemovePluginRepoCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (RemovePluginRepoCommand) Execute(args []string) error {
	return translatableerror.UnrefactoredCommandError{}
}

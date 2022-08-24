package pluginrepo

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commandregistry"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/flags"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/terminal"

	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
)

type ListPluginRepos struct {
	ui     terminal.UI
	config coreconfig.Reader
}

func init() {
	commandregistry.Register(&ListPluginRepos{})
}

func (cmd *ListPluginRepos) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "list-plugin-repos",
		Description: T("List all the added plugin repositories"),
		Usage: []string{
			T("CF_NAME list-plugin-repos"),
		},
	}
}

func (cmd *ListPluginRepos) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	usageReq := requirements.NewUsageRequirement(commandregistry.CLICommandUsagePresenter(cmd),
		T("No argument required"),
		func() bool {
			return len(fc.Args()) != 0
		},
	)

	reqs := []requirements.Requirement{
		usageReq,
	}
	return reqs, nil
}

func (cmd *ListPluginRepos) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	return cmd
}

func (cmd *ListPluginRepos) Execute(c flags.FlagContext) error {
	repos := cmd.config.PluginRepos()

	table := cmd.ui.Table([]string{T("Repo Name"), T("URL")})

	for _, repo := range repos {
		table.Add(repo.Name, repo.URL)
	}

	cmd.ui.Ok()
	cmd.ui.Say("")

	err := table.Print()
	if err != nil {
		return err
	}

	cmd.ui.Say("")
	return nil
}

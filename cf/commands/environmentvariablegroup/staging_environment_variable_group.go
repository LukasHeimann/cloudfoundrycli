package environmentvariablegroup

import (
	"sort"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/api/environmentvariablegroups"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commandregistry"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/flags"
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/models"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/terminal"
)

type StagingEnvironmentVariableGroup struct {
	ui                           terminal.UI
	config                       coreconfig.ReadWriter
	environmentVariableGroupRepo environmentvariablegroups.Repository
}

func init() {
	commandregistry.Register(&StagingEnvironmentVariableGroup{})
}

func (cmd *StagingEnvironmentVariableGroup) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "staging-environment-variable-group",
		Description: T("Retrieve the contents of the staging environment variable group"),
		ShortName:   "sevg",
		Usage: []string{
			T("CF_NAME staging-environment-variable-group"),
		},
	}
}

func (cmd *StagingEnvironmentVariableGroup) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	usageReq := requirements.NewUsageRequirement(commandregistry.CLICommandUsagePresenter(cmd),
		T("No argument required"),
		func() bool {
			return len(fc.Args()) != 0
		},
	)

	reqs := []requirements.Requirement{
		usageReq,
		requirementsFactory.NewLoginRequirement(),
	}
	return reqs, nil
}

func (cmd *StagingEnvironmentVariableGroup) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.environmentVariableGroupRepo = deps.RepoLocator.GetEnvironmentVariableGroupsRepository()
	return cmd
}

func (cmd *StagingEnvironmentVariableGroup) Execute(c flags.FlagContext) error {
	cmd.ui.Say(T("Retrieving the contents of the staging environment variable group as {{.Username}}...", map[string]interface{}{
		"Username": terminal.EntityNameColor(cmd.config.Username())}))

	stagingEnvVars, err := cmd.environmentVariableGroupRepo.ListStaging()
	if err != nil {
		return err
	}

	cmd.ui.Ok()

	table := cmd.ui.Table([]string{T("Variable Name"), T("Assigned Value")})
	sortedEnvVars := models.EnvironmentVariableList(stagingEnvVars)
	sort.Sort(sortedEnvVars)
	for _, envVar := range sortedEnvVars {
		table.Add(envVar.Name, envVar.Value)
	}
	err = table.Print()
	if err != nil {
		return err
	}
	return nil
}

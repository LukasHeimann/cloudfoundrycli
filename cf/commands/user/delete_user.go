package user

import (
	"fmt"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/api"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commandregistry"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/errors"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/flags"
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/terminal"
)

type DeleteUser struct {
	ui       terminal.UI
	config   coreconfig.Reader
	userRepo api.UserRepository
}

func init() {
	commandregistry.Register(&DeleteUser{})
}

func (cmd *DeleteUser) MetaData() commandregistry.CommandMetadata {
	fs := make(map[string]flags.FlagSet)
	fs["f"] = &flags.BoolFlag{ShortName: "f", Usage: T("Force deletion without confirmation")}

	return commandregistry.CommandMetadata{
		Name:        "delete-user",
		Description: T("Delete a user"),
		Usage: []string{
			T("CF_NAME delete-user USERNAME [-f]"),
		},
		Flags: fs,
	}
}

func (cmd *DeleteUser) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires an argument\n\n") + commandregistry.Commands.CommandUsage("delete-user"))
		return nil, fmt.Errorf("Incorrect usage: %d arguments of %d required", len(fc.Args()), 1)
	}

	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
	}

	return reqs, nil
}

func (cmd *DeleteUser) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.userRepo = deps.RepoLocator.GetUserRepository()
	return cmd
}

func (cmd *DeleteUser) Execute(c flags.FlagContext) error {
	username := c.Args()[0]
	force := c.Bool("f")

	if !force && !cmd.ui.ConfirmDelete(T("user"), username) {
		return nil
	}

	cmd.ui.Say(T("Deleting user {{.TargetUser}} as {{.CurrentUser}}...",
		map[string]interface{}{
			"TargetUser":  terminal.EntityNameColor(username),
			"CurrentUser": terminal.EntityNameColor(cmd.config.Username()),
		}))

	users, err := cmd.userRepo.FindAllByUsername(username)

	switch err.(type) {
	case nil:
		if len(users) > 1 {
			return fmt.Errorf(T(
				"Error deleting user {{.Username}} \nThe user exists in multiple origins.",
				map[string]interface{}{
					"Username": username,
				}))
		}
	case *errors.ModelNotFoundError:
		cmd.ui.Ok()
		cmd.ui.Warn(T("User {{.TargetUser}} does not exist.", map[string]interface{}{"TargetUser": username}))
		return nil
	default:
		return err
	}

	err = cmd.userRepo.Delete(users[0].GUID)
	if err != nil {
		return err
	}

	cmd.ui.Ok()
	return nil
}

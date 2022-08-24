package v7

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3/constant"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/flag"
	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
)

type CreateAppCommand struct {
	BaseCommand

	RequiredArgs    flag.AppName `positional-args:"yes"`
	AppType         flag.AppType `long:"app-type" choice:"buildpack" choice:"docker" description:"App lifecycle type to stage and run the app" default:"buildpack"`
	usage           interface{}  `usage:"CF_NAME create-app APP_NAME [--app-type (buildpack | docker)]"`
	relatedCommands interface{}  `related_commands:"app, apps, push"`
}

func (cmd CreateAppCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	user, err := cmd.Actor.GetCurrentUser()
	if err != nil {
		return err
	}

	cmd.UI.DisplayTextWithFlavor("Creating app {{.AppName}} in org {{.CurrentOrg}} / space {{.CurrentSpace}} as {{.CurrentUser}}...", map[string]interface{}{
		"AppName":      cmd.RequiredArgs.AppName,
		"CurrentSpace": cmd.Config.TargetedSpace().Name,
		"CurrentOrg":   cmd.Config.TargetedOrganization().Name,
		"CurrentUser":  user.Name,
	})

	_, warnings, err := cmd.Actor.CreateApplicationInSpace(
		resources.Application{
			Name:          cmd.RequiredArgs.AppName,
			LifecycleType: constant.AppLifecycleType(cmd.AppType),
		},
		cmd.Config.TargetedSpace().GUID,
	)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		switch err.(type) {
		case ccerror.NameNotUniqueInSpaceError:
			cmd.UI.DisplayText(err.Error())
		default:
			return err
		}
	}

	cmd.UI.DisplayOK()

	return nil
}

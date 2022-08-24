package v7

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/flag"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/v7/shared"
	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
)

type OrgQuotaCommand struct {
	BaseCommand

	RequiredArgs    flag.OrganizationQuota `positional-args:"yes"`
	usage           interface{}            `usage:"CF_NAME org-quota QUOTA"`
	relatedCommands interface{}            `related_commands:"org, org-quotas"`
}

func (cmd OrgQuotaCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(false, false)
	if err != nil {
		return err
	}

	user, err := cmd.Actor.GetCurrentUser()
	if err != nil {
		return err
	}

	quotaName := cmd.RequiredArgs.OrganizationQuotaName

	cmd.UI.DisplayTextWithFlavor(
		"Getting org quota {{.QuotaName}} as {{.Username}}...",
		map[string]interface{}{
			"QuotaName": quotaName,
			"Username":  user.Name,
		})
	cmd.UI.DisplayNewline()

	orgQuota, warnings, err := cmd.Actor.GetOrganizationQuotaByName(quotaName)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	quotaDisplayer := shared.NewQuotaDisplayer(cmd.UI)
	quotaDisplayer.DisplaySingleQuota(resources.Quota(orgQuota.Quota))

	return nil
}

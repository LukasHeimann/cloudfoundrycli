package ccv3

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3/internal"
	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
)

func (client *Client) GetApplicationRevisions(appGUID string, query ...Query) ([]resources.Revision, Warnings, error) {
	var revisions []resources.Revision

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetApplicationRevisionsRequest,
		Query:        query,
		URIParams:    internal.Params{"app_guid": appGUID},
		ResponseBody: resources.Revision{},
		AppendToList: func(item interface{}) error {
			revisions = append(revisions, item.(resources.Revision))
			return nil
		},
	})
	return revisions, warnings, err
}

func (client *Client) GetApplicationRevisionsDeployed(appGUID string) ([]resources.Revision, Warnings, error) {
	var revisions []resources.Revision

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetApplicationRevisionsDeployedRequest,
		URIParams:    internal.Params{"app_guid": appGUID},
		ResponseBody: resources.Revision{},
		AppendToList: func(item interface{}) error {
			revisions = append(revisions, item.(resources.Revision))
			return nil
		},
	})
	return revisions, warnings, err
}

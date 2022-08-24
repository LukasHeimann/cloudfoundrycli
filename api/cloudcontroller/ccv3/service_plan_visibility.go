package ccv3

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3/internal"
	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
)

func (client *Client) GetServicePlanVisibility(servicePlanGUID string) (resources.ServicePlanVisibility, Warnings, error) {
	var result resources.ServicePlanVisibility

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.GetServicePlanVisibilityRequest,
		URIParams:    internal.Params{"service_plan_guid": servicePlanGUID},
		ResponseBody: &result,
	})

	return result, warnings, err
}

func (client *Client) UpdateServicePlanVisibility(servicePlanGUID string, planVisibility resources.ServicePlanVisibility) (resources.ServicePlanVisibility, Warnings, error) {
	var result resources.ServicePlanVisibility

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PostServicePlanVisibilityRequest,
		URIParams:    internal.Params{"service_plan_guid": servicePlanGUID},
		RequestBody:  planVisibility,
		ResponseBody: &result,
	})

	return result, warnings, err
}

func (client *Client) DeleteServicePlanVisibility(servicePlanGUID, organizationGUID string) (Warnings, error) {

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName: internal.DeleteServicePlanVisibilityRequest,
		URIParams:   internal.Params{"service_plan_guid": servicePlanGUID, "organization_guid": organizationGUID},
	})

	return warnings, err
}

package ccv3

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3/internal"
	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
)

func (client *Client) GetSpaceFeature(spaceGUID string, featureName string) (bool, Warnings, error) {
	var responseBody resources.SpaceFeature

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.GetSpaceFeatureRequest,
		URIParams:    internal.Params{"space_guid": spaceGUID, "feature": featureName},
		ResponseBody: &responseBody,
	})

	return responseBody.Enabled, warnings, err
}

func (client *Client) UpdateSpaceFeature(spaceGUID string, enabled bool, featureName string) (Warnings, error) {
	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName: internal.PatchSpaceFeaturesRequest,
		URIParams:   internal.Params{"space_guid": spaceGUID, "feature": featureName},
		RequestBody: struct {
			Enabled bool `json:"enabled"`
		}{Enabled: enabled},
	})

	return warnings, err
}

package ccv3

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3/internal"
	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
)

// GetStacks lists stacks with optional filters.
func (client *Client) GetStacks(query ...Query) ([]resources.Stack, Warnings, error) {
	var stacks []resources.Stack

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetStacksRequest,
		Query:        query,
		ResponseBody: resources.Stack{},
		AppendToList: func(item interface{}) error {
			stacks = append(stacks, item.(resources.Stack))
			return nil
		},
	})

	return stacks, warnings, err
}

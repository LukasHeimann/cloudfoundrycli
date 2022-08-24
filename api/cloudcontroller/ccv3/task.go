package ccv3

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3/internal"
	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
)

// CreateApplication resources.Task runs a command in the Application environment
// associated with the provided Application GUID.
func (client *Client) CreateApplicationTask(appGUID string, task resources.Task) (resources.Task, Warnings, error) {
	var responseBody resources.Task

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PostApplicationTasksRequest,
		URIParams:    internal.Params{"app_guid": appGUID},
		RequestBody:  task,
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

// GetApplicationTasks returns a list of tasks associated with the provided
// application GUID. Results can be filtered by providing URL queries.
func (client *Client) GetApplicationTasks(appGUID string, query ...Query) ([]resources.Task, Warnings, error) {
	var tasks []resources.Task

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetApplicationTasksRequest,
		URIParams:    internal.Params{"app_guid": appGUID},
		Query:        query,
		ResponseBody: resources.Task{},
		AppendToList: func(item interface{}) error {
			tasks = append(tasks, item.(resources.Task))
			return nil
		},
	})

	return tasks, warnings, err
}

// UpdateTaskCancel cancels a task.
func (client *Client) UpdateTaskCancel(taskGUID string) (resources.Task, Warnings, error) {
	var responseBody resources.Task

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName: internal.PutTaskCancelRequest,
		URIParams: internal.Params{
			"task_guid": taskGUID,
		},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

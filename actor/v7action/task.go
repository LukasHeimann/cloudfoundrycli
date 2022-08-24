package v7action

import (
	"strconv"

	"sort"

	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/actionerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3"
	"github.com/LukasHeimann/cloudfoundrycli/v8/resources"
)

// Run resources.Task runs the provided command in the application environment associated
// with the provided application GUID.
func (actor Actor) RunTask(appGUID string, task resources.Task) (resources.Task, Warnings, error) {
	createdTask, warnings, err := actor.CloudControllerClient.CreateApplicationTask(appGUID, resources.Task(task))
	if err != nil {
		if e, ok := err.(ccerror.TaskWorkersUnavailableError); ok {
			return resources.Task{}, Warnings(warnings), actionerror.TaskWorkersUnavailableError{Message: e.Error()}
		}
	}

	return resources.Task(createdTask), Warnings(warnings), err
}

// GetApplicationTasks returns a list of tasks associated with the provided
// appplication GUID.
func (actor Actor) GetApplicationTasks(appGUID string, sortOrder SortOrder) ([]resources.Task, Warnings, error) {
	tasks, warnings, err := actor.CloudControllerClient.GetApplicationTasks(appGUID)
	actorWarnings := Warnings(warnings)
	if err != nil {
		return nil, actorWarnings, err
	}

	allTasks := []resources.Task{}
	for _, task := range tasks {
		allTasks = append(allTasks, resources.Task(task))
	}

	if sortOrder == Descending {
		sort.Slice(allTasks, func(i int, j int) bool { return allTasks[i].SequenceID > allTasks[j].SequenceID })
	} else {
		sort.Slice(allTasks, func(i int, j int) bool { return allTasks[i].SequenceID < allTasks[j].SequenceID })
	}

	return allTasks, actorWarnings, nil
}

func (actor Actor) GetTaskBySequenceIDAndApplication(sequenceID int, appGUID string) (resources.Task, Warnings, error) {
	tasks, warnings, err := actor.CloudControllerClient.GetApplicationTasks(
		appGUID,
		ccv3.Query{Key: ccv3.SequenceIDFilter, Values: []string{strconv.Itoa(sequenceID)}},
	)
	if err != nil {
		return resources.Task{}, Warnings(warnings), err
	}

	if len(tasks) == 0 {
		return resources.Task{}, Warnings(warnings), actionerror.TaskNotFoundError{SequenceID: sequenceID}
	}

	return resources.Task(tasks[0]), Warnings(warnings), nil
}

func (actor Actor) TerminateTask(taskGUID string) (resources.Task, Warnings, error) {
	task, warnings, err := actor.CloudControllerClient.UpdateTaskCancel(taskGUID)
	return resources.Task(task), Warnings(warnings), err
}

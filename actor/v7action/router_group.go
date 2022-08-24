package v7action

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/actionerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/router"
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/router/routererror"
)

type RouterGroup router.RouterGroup

func (actor Actor) GetRouterGroups() ([]RouterGroup, error) {
	var routerGroups []RouterGroup

	apiRouterGroups, err := actor.RoutingClient.GetRouterGroups()
	if err != nil {
		return nil, err
	}

	for _, group := range apiRouterGroups {
		routerGroups = append(routerGroups, RouterGroup(group))
	}

	return routerGroups, err
}

func (actor Actor) GetRouterGroupByName(name string) (RouterGroup, error) {
	apiRouterGroup, err := actor.RoutingClient.GetRouterGroupByName(name)
	if err != nil {
		if _, ok := err.(routererror.ResourceNotFoundError); ok {
			return RouterGroup{}, actionerror.RouterGroupNotFoundError{Name: name}
		}

		return RouterGroup{}, err
	}

	return RouterGroup(apiRouterGroup), nil
}

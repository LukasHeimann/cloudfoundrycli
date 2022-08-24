package ccv3

import "github.com/LukasHeimann/cloudfoundrycli/v8/resources"

type IncludedResources struct {
	Users            []resources.User            `json:"users,omitempty"`
	Organizations    []resources.Organization    `json:"organizations,omitempty"`
	Spaces           []resources.Space           `json:"spaces,omitempty"`
	ServiceInstances []resources.ServiceInstance `json:"service_instances,omitempty"`
	ServiceOfferings []resources.ServiceOffering `json:"service_offerings,omitempty"`
	ServiceBrokers   []resources.ServiceBroker   `json:"service_brokers,omitempty"`
	ServicePlans     []resources.ServicePlan     `json:"service_plans,omitempty"`
	Apps             []resources.Application     `json:"apps,omitempty"`
}

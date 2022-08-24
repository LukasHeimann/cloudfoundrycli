package requirements

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/api"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DomainRequirement

type DomainRequirement interface {
	Requirement
	GetDomain() models.DomainFields
}

type domainAPIRequirement struct {
	name       string
	config     coreconfig.Reader
	domainRepo api.DomainRepository
	domain     models.DomainFields
}

func NewDomainRequirement(name string, config coreconfig.Reader, domainRepo api.DomainRepository) (req *domainAPIRequirement) {
	req = new(domainAPIRequirement)
	req.name = name
	req.config = config
	req.domainRepo = domainRepo
	return
}

func (req *domainAPIRequirement) Execute() error {
	var apiErr error
	req.domain, apiErr = req.domainRepo.FindByNameInOrg(req.name, req.config.OrganizationFields().GUID)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (req *domainAPIRequirement) GetDomain() models.DomainFields {
	return req.domain
}

package commands_test

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/api/stacks/stacksfakes"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commandregistry"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/flags"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/models"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements/requirementsfakes"
	testcmd "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/commands"
	testconfig "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/configuration"
	testterm "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands"
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/matchers"
)

var _ = Describe("stacks command", func() {
	var (
		ui                  *testterm.FakeUI
		repo                *stacksfakes.FakeStackRepository
		config              coreconfig.Repository
		requirementsFactory *requirementsfakes.FakeFactory
		deps                commandregistry.Dependency
	)

	updateCommandDependency := func(pluginCall bool) {
		deps.UI = ui
		deps.Config = config
		deps.RepoLocator = deps.RepoLocator.SetStackRepository(repo)
		commandregistry.Commands.SetCommand(commandregistry.Commands.FindCommand("stacks").SetDependency(deps, pluginCall))
	}

	BeforeEach(func() {
		ui = &testterm.FakeUI{}
		config = testconfig.NewRepositoryWithDefaults()
		requirementsFactory = new(requirementsfakes.FakeFactory)
		requirementsFactory.NewLoginRequirementReturns(requirements.Passing{})
		repo = new(stacksfakes.FakeStackRepository)
	})

	Describe("login requirements", func() {
		It("fails if the user is not logged in", func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Failing{Message: "not logged in"})

			Expect(testcmd.RunCLICommand("stacks", []string{}, requirementsFactory, updateCommandDependency, false, ui)).To(BeFalse())
		})

		Context("when arguments are provided", func() {
			var cmd commandregistry.Command
			var flagContext flags.FlagContext

			BeforeEach(func() {
				cmd = &commands.ListStacks{}
				cmd.SetDependency(deps, false)
				flagContext = flags.NewFlagContext(cmd.MetaData().Flags)
			})

			It("should fail with usage", func() {
				flagContext.Parse("blahblah")

				reqs, err := cmd.Requirements(requirementsFactory, flagContext)
				Expect(err).NotTo(HaveOccurred())

				err = testcmd.RunRequirements(reqs)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Incorrect Usage"))
				Expect(err.Error()).To(ContainSubstring("No argument required"))
			})
		})
	})

	It("lists the stacks", func() {
		stack1 := models.Stack{
			Name:        "Stack-1",
			Description: "Stack 1 Description",
		}
		stack2 := models.Stack{
			Name:        "Stack-2",
			Description: "Stack 2 Description",
		}

		repo.FindAllReturns([]models.Stack{stack1, stack2}, nil)
		testcmd.RunCLICommand("stacks", []string{}, requirementsFactory, updateCommandDependency, false, ui)

		Expect(ui.Outputs()).To(ContainSubstrings(
			[]string{"Getting stacks in org", "my-org", "my-space", "my-user"},
			[]string{"OK"},
			[]string{"Stack-1", "Stack 1 Description"},
			[]string{"Stack-2", "Stack 2 Description"},
		))
	})
})

package environmentvariablegroup_test

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/api/environmentvariablegroups/environmentvariablegroupsfakes"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commandregistry"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/environmentvariablegroup"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/configuration/coreconfig"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/flags"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/models"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements/requirementsfakes"
	testcmd "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/commands"
	testconfig "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/configuration"
	. "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/matchers"
	testterm "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("running-environment-variable-group command", func() {
	var (
		ui                           *testterm.FakeUI
		requirementsFactory          *requirementsfakes.FakeFactory
		configRepo                   coreconfig.Repository
		environmentVariableGroupRepo *environmentvariablegroupsfakes.FakeRepository
		deps                         commandregistry.Dependency
	)

	updateCommandDependency := func(pluginCall bool) {
		deps.UI = ui
		deps.RepoLocator = deps.RepoLocator.SetEnvironmentVariableGroupsRepository(environmentVariableGroupRepo)
		deps.Config = configRepo
		commandregistry.Commands.SetCommand(commandregistry.Commands.FindCommand("running-environment-variable-group").SetDependency(deps, pluginCall))
	}

	BeforeEach(func() {
		ui = &testterm.FakeUI{}
		configRepo = testconfig.NewRepositoryWithDefaults()
		requirementsFactory = new(requirementsfakes.FakeFactory)
		environmentVariableGroupRepo = new(environmentvariablegroupsfakes.FakeRepository)
	})

	runCommand := func(args ...string) bool {
		return testcmd.RunCLICommand("running-environment-variable-group", args, requirementsFactory, updateCommandDependency, false, ui)
	}

	Describe("requirements", func() {
		It("requires the user to be logged in", func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Failing{Message: "not logged in"})
			Expect(runCommand()).ToNot(HavePassedRequirements())
		})

		Context("when arguments are provided", func() {
			var cmd commandregistry.Command
			var flagContext flags.FlagContext

			BeforeEach(func() {
				cmd = &environmentvariablegroup.RunningEnvironmentVariableGroup{}
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

	Describe("when logged in", func() {
		BeforeEach(func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Passing{})
			environmentVariableGroupRepo.ListRunningReturns(
				[]models.EnvironmentVariable{
					{Name: "abc", Value: "123"},
					{Name: "def", Value: "456"},
				}, nil)
		})

		It("Displays the running environment variable group", func() {
			runCommand()

			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"Retrieving the contents of the running environment variable group as my-user..."},
				[]string{"OK"},
				[]string{"Variable Name", "Assigned Value"},
				[]string{"abc", "123"},
				[]string{"def", "456"},
			))
		})
	})
})

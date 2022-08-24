package internal_test

import (
	"reflect"
	"strings"

	"github.com/LukasHeimann/cloudfoundrycli/v8/command/common"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/common/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("test help all display", func() {
	var fromCommandList []string
	var fromHelpAllDisplay []string

	BeforeEach(func() {
		commandList := common.Commands
		handler := reflect.TypeOf(commandList)
		for i := 0; i < handler.NumField(); i++ {
			fieldTag := handler.Field(i).Tag
			commandName := fieldTag.Get("command")
			if !(strings.HasPrefix(commandName, "v3-") || (commandName == "")) {
				fromCommandList = append(fromCommandList, commandName)
			}
		}

		for _, category := range internal.HelpCategoryList {
			for _, row := range category.CommandList {
				for _, command := range row {
					if !(strings.HasPrefix(command, "v3-") || (command == "")) {
						fromHelpAllDisplay = append(fromHelpAllDisplay, command)
					}
				}
			}
		}

		for _, category := range internal.ExperimentalHelpCategoryList {
			for _, row := range category.CommandList {
				for _, command := range row {
					if command != "" {
						fromHelpAllDisplay = append(fromHelpAllDisplay, command)
					}
				}
			}
		}
	})

	It("lists all commands from command list in at least one category", func() {
		for _, command := range fromCommandList {
			Expect(fromHelpAllDisplay).To(ContainElement(command))
		}
	})
})

package command_parser

import (
	"strings"

	"github.com/LukasHeimann/cloudfoundrycli/v8/command/translatableerror"
)

func preventExtraArgs(args []string) error {
	if len(args) > 0 {
		return translatableerror.TooManyArgumentsError{
			ExtraArgument: strings.Join(args, " "),
		}
	}
	return nil
}

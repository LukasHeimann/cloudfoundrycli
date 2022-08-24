package v7pushaction

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/translatableerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/manifestparser"
)

func HandleInstancesOverride(manifest manifestparser.Manifest, overrides FlagOverrides) (manifestparser.Manifest, error) {
	if overrides.Instances.IsSet {
		if manifest.ContainsMultipleApps() {
			return manifest, translatableerror.CommandLineArgsWithMultipleAppsError{}
		}

		webProcess := manifest.GetFirstAppWebProcess()
		if webProcess != nil {
			webProcess.Instances = &overrides.Instances.Value
		} else {
			app := manifest.GetFirstApp()
			app.Instances = &overrides.Instances.Value
		}
	}

	return manifest, nil
}

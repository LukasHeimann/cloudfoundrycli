package v7pushaction

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/errors"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/translatableerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/manifestparser"
)

func HandleAppNameOverride(manifest manifestparser.Manifest, overrides FlagOverrides) (manifestparser.Manifest, error) {
	if manifest.ContainsMultipleApps() && manifest.HasAppWithNoName() {
		return manifest, errors.New("Found an application with no name specified.")
	}

	if overrides.AppName != "" {
		newApps := make([]manifestparser.Application, 1)

		foundApp := false
		for _, app := range manifest.Applications {
			if app.Name == overrides.AppName {
				newApps[0] = app
				foundApp = true
				break
			}
		}

		if !foundApp {
			if len(manifest.Applications) == 1 {
				manifest.Applications[0].Name = overrides.AppName
				return manifest, nil
			}

			return manifest, manifestparser.AppNotInManifestError{Name: overrides.AppName}
		}

		manifest.Applications = newApps
	} else if manifest.HasAppWithNoName() {
		return manifest, translatableerror.AppNameOrManifestRequiredError{}
	}

	return manifest, nil
}

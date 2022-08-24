package v7pushaction

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/manifestparser"
)

func HandleDefaultRouteOverride(manifest manifestparser.Manifest, overrides FlagOverrides) (manifestparser.Manifest, error) {
	for i := range manifest.Applications {
		if manifest.Applications[i].RandomRoute || manifest.Applications[i].NoRoute {
			continue
		}
		manifest.Applications[i].DefaultRoute = true
	}

	return manifest, nil
}

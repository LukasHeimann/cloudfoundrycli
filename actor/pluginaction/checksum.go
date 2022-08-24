package pluginaction

import "github.com/LukasHeimann/cloudfoundrycli/v8/util/configv3"

func (actor Actor) ValidateFileChecksum(path string, checksum string) bool {
	plugin := configv3.Plugin{Location: path}
	return plugin.CalculateSHA1() == checksum
}

package terrago

import "log"

// Get calls terraform get and return stdout/stderr.
func Get(options *Options) string {
	out, err := GetE(options)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

// GetE calls terraform get and return stdout/stderr.
func GetE(options *Options) (string, error) {
	return RunTerraformCommandE(options, "get", "-update")
}

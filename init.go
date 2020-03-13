package terrago

import (
	"fmt"
)

// Init calls `terraform init`, and returns stderr/stdout
func Init(options *Options) (string, error) {
	args := []string{"init", fmt.Sprintf("-upgrade=%t", options.Upgrade)}
	args = append(args, FormatTerraformBackendConfigAsArgs(options.BackendConfig)...)
	return RunTerraformCommand(options, args...)
}

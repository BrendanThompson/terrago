package terrago

import (
	"fmt"
	"log"
)

// Init calls `terraform init`, and returns stdout
func Init(options *Options) string {
	out, err := InitE(options)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

// InitE calls `terraform init`, and returns stderr/stdout
func InitE(options *Options) (string, error) {
	args := []string{"init", fmt.Sprintf("-upgrade=%t", options.Upgrade)}
	args = append(args, FormatTerraformBackendConfigAsArgs(options.BackendConfig)...)
	return RunTerraformCommandE(options, args...)
}

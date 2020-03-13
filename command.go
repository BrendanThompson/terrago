package terrago

import (
	"fmt"
	"log"

	"brendanthompson.dev/terrago/internal/retry"
	"brendanthompson.dev/terrago/internal/shell"
	"github.com/gruntwork-io/terratest/modules/collections"
)

// GetCommonOptions extracts commons terraform options
func GetCommonOptions(options *Options, args ...string) (*Options, []string) {
	if options.NoColor && !collections.ListContains(args, "-no-color") {
		args = append(args, "-no-color")
	}

	if options.TerraformBinary == "" {
		options.TerraformBinary = "terraform"
	}

	// Initialize EnvVars, if it hasn't been set yet
	if options.EnvVars == nil {
		options.EnvVars = map[string]string{}
	}

	return options, args
}

// RunTerraformCommand runs `terraform` with the given arguments
// and options. It then returns stdout/stderr
func RunTerraformCommand(additionalOptions *Options, additionalArgs ...string) (string, error) {
	//TODO: Implement `GetCommonOptions`
	options, args := GetCommonOptions(additionalOptions, additionalArgs...)

	cmd := shell.Command{
		Command:           options.TerraformBinary,
		Args:              args,
		WorkingDir:        options.TerraformDir,
		Env:               options.EnvVars,
		OutputMaxLineSize: options.OutputMaxLineSize,
	}

	description := fmt.Sprintf("%s %v", options.TerraformBinary, args)

	return retry.DoWithRetryableErrorsE(description, options.RetryableTerraformErrors, options.MaxRetries, options.TimeBetweenRetries, func() (string, error) {
		return shell.RunCommandAndGetOutputE(cmd)
	})
}

// GetExitCodeForTerraformCommand runs terraform with the given arguments and options and returns exit code
func GetExitCodeForTerraformCommand(additionalOptions *Options, additionalArgs ...string) (int, error) {
	options, args := GetCommonOptions(additionalOptions, additionalArgs...)

	log.Printf("Running %s with args %v", options.TerraformBinary, args)
	cmd := shell.Command{
		Command:           options.TerraformBinary,
		Args:              args,
		WorkingDir:        options.TerraformDir,
		Env:               options.EnvVars,
		OutputMaxLineSize: options.OutputMaxLineSize,
	}

	_, err := shell.RunCommandAndGetOutputE(cmd)
	if err == nil {
		return DefaultSuccessExitCode, nil
	}
	exitCode, getExitCodeErr := shell.GetExitCodeForRunCommandError(err)
	if getExitCodeErr == nil {
		return exitCode, nil
	}
	return DefaultErrorExitCode, getExitCodeErr
}

package terrago

import (
	"log"
)

// InitAndPlan runs terraform init and plan with the given options and returns stdout/stderr from the plan command.
// This will fail the test if there is an error in the command.
func InitAndPlan(options *Options) string {
	out, err := InitAndPlanE(options)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

// InitAndPlanE runs terraform init and plan with the given options and returns stdout/stderr from the plan command.
func InitAndPlanE(options *Options) (string, error) {
	if _, err := InitE(options); err != nil {
		return "", err
	}

	if _, err := GetE(options); err != nil {
		return "", err
	}

	return PlanE(options)
}

// Plan runs terraform plan with the given options and returns stdout/stderr.
// This will fail the test if there is an error in the command.
func Plan(options *Options) string {
	out, err := PlanE(options)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

// PlanE runs terraform plan with the given options and returns stdout/stderr.
func PlanE(options *Options) (string, error) {
	return RunTerraformCommandE(options, FormatArgs(options, "plan", "-input=false", "-lock=false")...)
}

// InitAndPlanWithExitCode runs terraform init and plan with the given options and returns exitcode for the plan command.
// This will fail the test if there is an error in the command.
func InitAndPlanWithExitCode(options *Options) int {
	exitCode, err := InitAndPlanWithExitCodeE(options)
	if err != nil {
		log.Fatal(err)
	}
	return exitCode
}

// InitAndPlanWithExitCodeE runs terraform init and plan with the given options and returns exitcode for the plan command.
func InitAndPlanWithExitCodeE(options *Options) (int, error) {
	if _, err := InitE(options); err != nil {
		return DefaultErrorExitCode, err
	}

	return PlanExitCodeE(options)
}

// PlanExitCode runs terraform plan with the given options and returns the detailed exitcode.
// This will fail the test if there is an error in the command.
func PlanExitCode(options *Options) int {
	exitCode, err := PlanExitCodeE(options)
	if err != nil {
		log.Fatal(err)
	}
	return exitCode
}

// PlanExitCodeE runs terraform plan with the given options and returns the detailed exitcode.
func PlanExitCodeE(options *Options) (int, error) {
	return GetExitCodeForTerraformCommandE(options, FormatArgs(options, "plan", "-input=false", "-detailed-exitcode")...)
}

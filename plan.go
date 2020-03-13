package terrago

// InitAndPlan runs terraform init and plan with the given options and returns stdout/stderr from the plan command.
func InitAndPlan(options *Options) (string, error) {
	if _, err := Init(options); err != nil {
		return "", err
	}

	if _, err := GetE(options); err != nil {
		return "", err
	}

	return Plan(options)
}

// Plan runs terraform plan with the given options and returns stdout/stderr.
func Plan(options *Options) (string, error) {
	return RunTerraformCommand(options, FormatArgs(options, "plan", "-input=false", "-lock=false")...)
}

// InitAndPlanWithExitCode runs terraform init and plan with the given options and returns exitcode for the plan command.
func InitAndPlanWithExitCode(options *Options) (int, error) {
	if _, err := Init(options); err != nil {
		return DefaultErrorExitCode, err
	}

	return PlanExitCode(options)
}

// PlanExitCode runs terraform plan with the given options and returns the detailed exitcode.
func PlanExitCode(options *Options) (int, error) {
	return GetExitCodeForTerraformCommand(options, FormatArgs(options, "plan", "-input=false", "-detailed-exitcode")...)
}

package terrago

import (
	"log"
)

// InitAndApply runs terraform init and apply with the given options and return stdout/stderr from the apply command. Note that this
// method does NOT call destroy and assumes the caller is responsible for cleaning up any resources created by running
// apply.
func InitAndApply(options *Options) string {
	out, err := InitAndApplyE(options)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

// InitAndApplyE runs terraform init and apply with the given options and return stdout/stderr from the apply command. Note that this
// method does NOT call destroy and assumes the caller is responsible for cleaning up any resources created by running
// apply.
func InitAndApplyE(options *Options) (string, error) {
	if _, err := InitE(options); err != nil {
		return "", err
	}

	if _, err := GetE(options); err != nil {
		return "", err
	}

	return ApplyE(options)
}

// Apply runs terraform apply with the given options and return stdout/stderr. Note that this method does NOT call destroy and
// assumes the caller is responsible for cleaning up any resources created by running apply.
func Apply(options *Options) string {
	out, err := ApplyE(options)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

// ApplyE runs terraform apply with the given options and return stdout/stderr. Note that this method does NOT call destroy and
// assumes the caller is responsible for cleaning up any resources created by running apply.
func ApplyE(options *Options) (string, error) {
	return RunTerraformCommandE(options, FormatArgs(options, "apply", "-input=false", "-auto-approve")...)
}

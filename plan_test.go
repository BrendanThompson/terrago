package terrago

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/stretchr/testify/require"
)

func TestInitAndPlanWithError(t *testing.T) {
	t.Parallel()

	testFolder, err := files.CopyTerraformFolderToTemp("test/fixtures/terraform-with-plan-error", t.Name())
	require.NoError(t, err)

	options := &Options{
		TerraformDir: testFolder,
	}

	_, err = InitAndPlanE(options)
	require.Error(t, err)
}

func TestInitAndPlanWithNoError(t *testing.T) {
	t.Parallel()

	testFolder, err := files.CopyTerraformFolderToTemp("test/fixtures/terraform-no-error", t.Name())
	require.NoError(t, err)

	options := &Options{
		TerraformDir: testFolder,
	}

	out, err := InitAndPlanE(options)
	require.NoError(t, err)
	require.Contains(t, out, "No changes. Infrastructure is up-to-date.")
}

func TestPlanWithExitCodeWithNoChanges(t *testing.T) {
	t.Parallel()
	testFolder, err := files.CopyTerraformFolderToTemp("test/fixtures/terraform-no-error", t.Name())
	require.NoError(t, err)

	options := &Options{
		TerraformDir: testFolder,
	}
	exitCode := InitAndPlanWithExitCode(options)
	require.Equal(t, DefaultSuccessExitCode, exitCode)
}

func TestPlanWithExitCodeWithChanges(t *testing.T) {
	t.Parallel()
	testFolder, err := files.CopyTerraformFolderToTemp("test/fixtures/terraform-basic-configuration", t.Name())
	require.NoError(t, err)

	options := &Options{
		TerraformDir: testFolder,
		Vars: map[string]interface{}{
			"cnt": 1,
		},
	}
	exitCode := InitAndPlanWithExitCode(options)
	require.Equal(t, TerraformPlanChangesPresentExitCode, exitCode)
}

func TestPlanWithExitCodeWithFailure(t *testing.T) {
	t.Parallel()

	testFolder, err := files.CopyTerraformFolderToTemp("test/fixtures/terraform-with-plan-error", t.Name())
	require.NoError(t, err)

	options := &Options{
		TerraformDir: testFolder,
	}

	exitCode, getExitCodeErr := InitAndPlanWithExitCodeE(options)
	require.NoError(t, getExitCodeErr)
	require.Equal(t, exitCode, 1)
}

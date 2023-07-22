package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestRecreate(t *testing.T) {
	t.Parallel()
	uniqueID := random.UniqueId()
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		Vars: map[string]interface{}{
			"cluster_name": strings.ToLower("gke-dev-" + uniqueID),
		},
		TerraformDir: "../modules/gke",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

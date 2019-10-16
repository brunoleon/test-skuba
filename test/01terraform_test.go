package test

import (
	"testing"
  "fmt"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/gruntwork-io/terratest/modules/shell"
  test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
func Test01Terraform(t *testing.T) {
	t.Parallel()

  exp_ip_load_balancer := "10.17.1.0"
  exp_ip_masters := []string{}
  exp_ip_masters = append(exp_ip_masters, "[10.17.2.0]")
  exp_ip_workers := []string{"[10.17.3.0 10.17.3.1]"}

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/01terraform",

//		// Variables to pass to our Terraform code using -var options
//		Vars: map[string]interface{}{
//			"example": expectedText,
//
//			// We also can see how lists and maps translate between terratest and terraform.
//			"example_list": expectedList,
//			"example_map":  expectedMap,
//		},

		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"terraform.auto.tfvars.json"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
//	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

//	// Run `terraform output` to get the values of output variables
	ip_load_balancer := terraform.Output(t, terraformOptions, "ip_load_balancer")
	ip_masters := terraform.OutputList(t, terraformOptions, "ip_masters")
	ip_workers := terraform.OutputList(t, terraformOptions, "ip_workers")
//
//	// Verify we're getting back the outputs we expect
	assert.Equal(t, exp_ip_load_balancer, ip_load_balancer)
	assert.Equal(t, exp_ip_masters, ip_masters)
	assert.Equal(t, exp_ip_workers, ip_workers)
}

// Deploy CaaSP using skuba
func skuba(t *testing.T) {
	test_structure.RunTestStage(t, "skubaInit", func() {
		skubaInit(t)
	})
	test_structure.RunTestStage(t, "skubaBootstrap", func() {
		skubaBootstrap(t)
	})
	test_structure.RunTestStage(t, "skubaJoin1", func() {
		skubaJoin1(t)
	})
	test_structure.RunTestStage(t, "skubaJoin2", func() {
		skubaJoin2(t)
	})
}

func skubaInit(t *testing.T) {
  cluster := "company-cluster"
 // expectedText := "[bootstrap] successfully bootstrapped core add-ons on node \"10.17.2.0\""
 // defer os.RemoveAll(cluster)
  cmdArgs := []string{}
  cmdArgs = append(cmdArgs, "cluster", "init", "--control-plane", "testing-lb.caasp.local", cluster)

  cmd := shell.Command{
    Command: "skuba",
    Args:    cmdArgs,
  }
  out, err :=  shell.RunCommandAndGetOutputE(t, cmd)
  fmt.Println(out)
  fmt.Println(err)
//	assert.Equal(t, expectedText, out)

  assert.DirExists(t, "company-cluster")
}

func skubaBootstrap(t *testing.T) {
  cmdArgs := []string{}
  cmdArgs = append(cmdArgs, "node", "bootstrap", "--user", "sles", "--sudo", "--target", "10.17.2.0", "my-master")

  cmd := shell.Command{
    Command: "skuba",
    Args:    cmdArgs,
    WorkingDir: "company-cluster",
  }
  out, err :=  shell.RunCommandAndGetOutputE(t, cmd)

  fmt.Println(out)
  fmt.Println(err)
}

func skubaJoin1(t *testing.T) {
  cmdArgs := []string{}
  cmdArgs = append(cmdArgs, "node", "join", "--role", "worker", "--user", "sles", "--sudo", "--target", "10.17.3.0", "worker0")

  cmd := shell.Command{
    Command: "skuba",
    Args:    cmdArgs,
    WorkingDir: "company-cluster",
  }
  out, err :=  shell.RunCommandAndGetOutputE(t, cmd)
  fmt.Println(out)
  fmt.Println(err)
}

func skubaJoin2(t *testing.T) {
  cmdArgs := []string{}
  cmdArgs = append(cmdArgs, "node", "join", "--role", "worker", "--user", "sles", "--sudo", "--target", "10.17.3.1", "worker1")

  cmd := shell.Command{
    Command: "skuba",
    Args:    cmdArgs,
    WorkingDir: "company-cluster",
  }
  out, err :=  shell.RunCommandAndGetOutputE(t, cmd)
  fmt.Println(out)
  fmt.Println(err)
}

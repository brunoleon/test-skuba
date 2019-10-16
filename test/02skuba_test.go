package test

import (
	"testing"
  "fmt"
  //"os"

	"github.com/stretchr/testify/assert"
	"github.com/gruntwork-io/terratest/modules/shell"
  test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func Test02Skuba(t *testing.T) {
	t.Parallel()
  skuba(t)

}

// Deploy CaaSP using skuba
func skuba(t *testing.T) {
 // cluster := "company-cluster"
 // defer os.RemoveAll(cluster)
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

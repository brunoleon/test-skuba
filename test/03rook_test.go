package test

import (
	"testing"
  "fmt"

	"github.com/gruntwork-io/terratest/modules/shell"
)

func Test03Rook(t *testing.T) {
	t.Parallel()
  rook(t)

}

// Deploy CaaSP using skuba
func rook(t *testing.T) {
	cluster := "company-cluster"
  rook := "rook"
 // defer os.RemoveAll(cluster)
  arrayOne := [3]string{"common.yaml", "operator.yaml", "cluster-test.yaml"}
	for index,element := range arrayOne{
		fmt.Println(index)
		fmt.Println(element)

		cmdArgs := []string{}
		cmdArgs = append(cmdArgs, "--kubeconfig", fmt.Sprintf("%s/admin.conf", cluster), "create", "-f", fmt.Sprintf("%s/%s", rook, element))

		cmd := shell.Command{
			Command: "kubectl",
			Args:    cmdArgs,
		}
		out, err :=  shell.RunCommandAndGetOutputE(t, cmd)
		fmt.Println(out)
		fmt.Println(err)
		}
}

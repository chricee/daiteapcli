package daiteapcli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/spf13/cobra"
)

var k8sRenameCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "rename",
	Aliases:       []string{},
	Short:         "Command to rename Kubernetes cluster",
	Args:          cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		requiredFlags := []string{"cluster", "name"}
		checkForRequiredFlags(requiredFlags, cmd)

        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		clusterID, _ := cmd.Flags().GetString("cluster")
		isKubernetes, err := IsKubernetes(clusterID)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if isKubernetes == false {
			fmt.Println("Please enter valid Kubernetes cluster ID")
			os.Exit(0)
		}

		name, _ := cmd.Flags().GetString("name")
		method := "POST"
		endpoint := "/renameCluster"
		requestBody := "{\"clusterID\": \"" + clusterID + "\", \"clusterName\": \"" + name + "\"}"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, requestBody)

		if err != nil {
			fmt.Println(err)
		} else {
			output, _ := json.MarshalIndent(responseBody, "", "    ")
			fmt.Println(string(output))
		}
	},
}

func init() {
	k8sCmd.AddCommand(k8sRenameCmd)

	parameters := [][]interface{}{
		[]interface{}{"cluster", "ID of the Kubernetes cluster", "string"},
		[]interface{}{"name", "new name of the Kubernetes cluster", "string"},
	}

	addParameterFlags(parameters, k8sRenameCmd)
}
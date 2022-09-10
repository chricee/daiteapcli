package daiteapcli

import (
	"encoding/json"
	"fmt"

	"github.com/Daiteap/daiteapcli/pkg/daiteapcli"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var projectsListCmd = &cobra.Command{
	SilenceUsage:  true,
	SilenceErrors: true,
	Use:           "list",
	Aliases:       []string{},
	Short:         "Command to list projects from current tenant",
	Args:          cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		outputFormat, _ := cmd.Flags().GetString("output")
		method := "GET"
		endpoint := "/projects"
		responseBody, err := daiteapcli.SendDaiteapRequest(method, endpoint, "")

		if err != nil {
			fmt.Println(err)
		} else {
			if outputFormat == "json" {
				output, _ := json.MarshalIndent(responseBody, "", "    ")
				fmt.Println(string(output))
			} else {
				tbl := table.New("Name", "Created at", "Contact")

				for _, project := range responseBody["data"].([]interface{}) {
					projectObject := project.(map[string]interface{})
					tbl.AddRow(projectObject["name"], projectObject["created_at"], projectObject["contact"])
				}

				tbl.Print()
			}
		}
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}

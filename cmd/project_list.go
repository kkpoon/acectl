package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	aceproject "github.com/kkpoon/go-aceproject"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectListCmd represents the list command
var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "list project",
	Long:  `list the project related to you`,
	Run: func(cmd *cobra.Command, args []string) {

		guid := viper.GetString("GUID")
		guidInfo := aceproject.GUIDInfo{GUID: guid}
		projectSvc := aceproject.NewProjectService(&http.Client{}, &guidInfo)

		projects, _, err := projectSvc.List()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		data := [][]string{}
		for _, p := range projects {
			data = append(data, []string{
				strconv.Itoa(p.ID), p.Name, p.Type, p.ProjectNumber,
			})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Type", "Project Number"})
		table.SetBorder(false)
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}

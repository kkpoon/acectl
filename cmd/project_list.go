package cmd

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	aceproject "github.com/kkpoon/go-aceproject"
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

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
		fmt.Fprintf(w, "ID\tCode\tType\tName\n")
		for _, p := range projects {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", p.ID, p.ProjectNumber, p.Type, p.Name)
		}
		w.Flush()
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}

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

// taskListCmd represents the list command
var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Long:  `list my tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		projectID := viper.GetInt64("project_id")
		guid := viper.GetString("GUID")
		guidInfo := aceproject.GUIDInfo{GUID: guid}
		taskSvc := aceproject.NewTaskService(&http.Client{}, &guidInfo)

		var tasks []aceproject.Task
		var err error
		if projectID == 0 {
			tasks, _, err = taskSvc.List()
		} else {
			tasks, _, err = taskSvc.ListWithProject(projectID)
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
		fmt.Fprintf(w, "ID\tProject Name\tTask Type\tTask Name\n")
		for _, t := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", t.ID, t.ProjectName, t.TaskGroupName, t.Name)
		}
		w.Flush()
	},
}

func init() {
	taskCmd.AddCommand(taskListCmd)
	taskListCmd.Flags().Int64P("projectid", "p", 0, "the Project ID")
	viper.BindPFlag("project_id", taskListCmd.Flags().Lookup("projectid"))
}

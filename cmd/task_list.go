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

		data := [][]string{}
		for _, i := range tasks {
			data = append(data, []string{
				strconv.Itoa(i.ID),
				fmt.Sprintf("%.25s", i.Name),
				i.TaskGroupName,
				strconv.Itoa(i.ProjectID),
				i.ProjectName,
			})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Type", "P_ID", "Project Name"})
		table.SetBorder(false)
		table.SetAutoWrapText(true)
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	},
}

func init() {
	taskCmd.AddCommand(taskListCmd)
	taskListCmd.Flags().Int64P("projectid", "p", 0, "the Project ID")
	viper.BindPFlag("project_id", taskListCmd.Flags().Lookup("projectid"))
}

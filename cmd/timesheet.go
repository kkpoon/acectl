package cmd

import "github.com/spf13/cobra"

// timesheetCmd represents the timesheet command
var timesheetCmd = &cobra.Command{
	Use:   "timesheet",
	Short: "a command about timesheet",
	Long:  `a command to do timesheet action`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// TODO: Work your own magic here
	// 	fmt.Println("timesheet called")
	// },
}

func init() {
	RootCmd.AddCommand(timesheetCmd)
}

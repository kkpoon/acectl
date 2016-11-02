package cmd

import "github.com/spf13/cobra"

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "task command",
	Long:  `command about task`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// TODO: Work your own magic here
	// 	fmt.Println("task called")
	// },
}

func init() {
	RootCmd.AddCommand(taskCmd)
}

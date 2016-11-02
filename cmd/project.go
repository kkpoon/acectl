package cmd

import "github.com/spf13/cobra"

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "project command",
	Long:  `commands about project`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// TODO: Work your own magic here
	// 	fmt.Println("project called")
	// },
}

func init() {
	RootCmd.AddCommand(projectCmd)
}

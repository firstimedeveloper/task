package cmd

import (
	"fmt"

	"github.com/firstimedeveloper/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current tasks",
	Run: func(cmd *cobra.Command, args []string) {
		task, _ := db.ListTasks()
		for _, t := range task {
			fmt.Printf("%d: %s\n", t.Key, t.Value)

		}
	},
}

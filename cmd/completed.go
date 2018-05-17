package cmd

import (
	"fmt"

	"github.com/firstimedeveloper/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(compCmd)
}

var compCmd = &cobra.Command{
	Use:   "compl",
	Short: "List completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		task, _ := db.ListTasks("done")
		for _, t := range task {
			fmt.Printf("%d: %s\n", t.Key, t.Value)
		}
	},
}

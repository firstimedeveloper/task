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
	Use:   "list [\"compl\"]",
	Short: "List current tasks. \"list compl\" will show the completed list",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			if args[0] == "compl" {
				task, _ := db.ListTasks("done")
				for _, t := range task {
					fmt.Printf("%d: %s\n", t.Key, t.Value)
				}
			}
		} else {
			task, _ := db.ListTasks("tasks")
			for _, t := range task {
				fmt.Printf("%d: %s\n", t.Key, t.Value)
			}
		}

	},
}

package cmd

import (
	"strings"

	"github.com/firstimedeveloper/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db.AddTask(strings.Join(args, " "), "tasks")
	},
}

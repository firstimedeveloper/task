package cmd

import (
	"strconv"

	"github.com/firstimedeveloper/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete specified task",
	Run: func(cmd *cobra.Command, args []string) {
		taskNum, _ := strconv.Atoi(args[0])
		db.DeleteTask(taskNum)
	},
}

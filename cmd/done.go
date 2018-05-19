package cmd

import (
	"fmt"
	"strconv"

	"github.com/firstimedeveloper/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doneCmd)
}

var doneCmd = &cobra.Command{
	Use:   "done [task #]",
	Short: "Done with task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskNum, _ := strconv.Atoi(args[0])
		nameOfTask, _ := db.GetValueFromKey(taskNum)
		db.DoneTask(taskNum)
		fmt.Printf("Done with task '%s'\n", nameOfTask)

	},
}

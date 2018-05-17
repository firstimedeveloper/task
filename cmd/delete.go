package cmd

import (
	"fmt"
	"strconv"

	"github.com/firstimedeveloper/task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [taskNum] optional(\"compl\")",
	Short: "Delete specified task",
	Long: `deletes the task with taskNum
	 in the tasklist. when \"compl\" is 
	 added after the taskNum, the corresponding
	 task in the completed list is deleted`,
	Run: func(cmd *cobra.Command, args []string) {
		taskNum, _ := strconv.Atoi(args[0])
		if len(args) == 2 {
			if args[1] == "compl" {
				db.DeleteTask(taskNum, "done")
				fmt.Printf("Deleted task %d in completed tasks list\n", taskNum)
			} else {
				fmt.Println("Check your arguments")
			}
		} else {
			db.DeleteTask(taskNum, "tasks")
			fmt.Printf("Deleted task %d in task list\n", taskNum)
		}

	},
}

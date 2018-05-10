package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doneCmd)
}

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Done with task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: Done with task")
	},
}

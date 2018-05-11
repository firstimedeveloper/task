package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"

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

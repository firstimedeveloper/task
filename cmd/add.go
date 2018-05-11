package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"

)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: add task")

	},
}

func addTask() {
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([])
	})	
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"

)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current tasks",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("TODO: List current tasks")
		listTasks()
	},
}

func listTasks() {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		c := b.Cursor()

		for k, v := c.First(); k!= nil; k, v = c.Next() {
			fmt.Printf("task %s: %s", k, v)
		}
		return nil
	})
}
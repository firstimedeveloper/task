package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
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
		fmt.Println("TODO: Done with task")
		doneWithTask(args[0])
	},
}

func doneWithTask(taskNum string) error {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("task"))

		err := b.Delete([]byte(taskNum))
		if err == nil {
			fmt.Println("Not a valid task number")
			return nil
		}
		return err
	})
}

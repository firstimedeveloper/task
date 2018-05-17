package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var doneBucket = []byte("done")
var db *bolt.DB

//Task is a struct that contains a Key and aValue
type Task struct {
	Key   int
	Value string
}

//Init is different from init, and it is first called to initialize the database
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		_, err = tx.CreateBucketIfNotExists(doneBucket)

		return err
	})
}

//taken from https://github.com/boltdb/bolt/issues/436
// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

//AddTask takes a task of type string, adds the string to the database and returns an id (and an error)
func AddTask(task string, nameOfBucket string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(nameOfBucket))

		// c := b.Cursor()
		// for k,v := c.First(); k != nil; k,v := c.Next() {

		// }

		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(int(id64))

		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

//ListTasks returns a slice of task structs (and an error)
func ListTasks(nameOfBucket string) ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(nameOfBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//DoneTask takes in an integar (task number) and marks it complete (moves to doneBucket and deletes it)
func DoneTask(num int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		bDone := tx.Bucket(doneBucket)
		value := b.Get(itob(num))
		b.Delete(itob(num))
		return bDone.Put(itob(num), []byte(value))
	})
	return err
}

//DeleteTask deletes a task specified by its taskNum
func DeleteTask(num int, nameOfBucket string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(nameOfBucket))
		return b.Delete(itob(num))
	})
	return err
}

// ResetTasks resets the ordering of the tasks
func ResetTasks(nameOfBucket string) error {
	var tasks []Task
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(nameOfBucket))

		var count = 1
		return b.ForEach(func(k, v []byte) error {
			value := b.Get(k)
			tasks = append(tasks, Task{
				Key:   count,
				Value: string(value),
			})
			b.Delete(k)
			count++
			return nil
		})
		for _, t := range tasks {
			b.Put(itob(t.Key), []byte(t.Value))
		}
		b.SetSequence(uint64(count))
		return nil
	})

	return err
}

func getLenBucket(nameOfBucket string) (int, error) {
	var count int
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(nameOfBucket))
		return b.ForEach(func(k, v []byte) error {
			count++
			return nil
		})
	})
	return count, err
}

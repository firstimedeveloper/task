package main

import (
	"path/filepath"

	"github.com/firstimedeveloper/task/cmd"
	"github.com/firstimedeveloper/task/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}

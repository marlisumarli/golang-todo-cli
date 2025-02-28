package cmd

import (
	"flag"
	"fmt"
	"golang-todo-cli/todo"
	"os"
)

func UpdateTask(todos *todo.Todos, args []string) {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateID := updateCmd.Int("i", 0, "Id bakalan diupdate")
	updateCat := updateCmd.String("c", "", "Category update")
	updateTask := updateCmd.String("t", "", "Task update")
	updateDone := updateCmd.Int("d", 2, "Done update")

	updateCmd.Parse(args)

	if *updateID < 0 {
		fmt.Println("Error: flag butuh flag -i.")
		os.Exit(1)
	}

	err := todos.Update(*updateID, *updateTask, *updateCat, *updateDone)
	if err != nil {
		fmt.Println(err)
	}

	err = todos.Store(GetJsonFile())
	if err != nil {
		fmt.Println(err)
	}

	todos.Print(2, "")
	fmt.Println("Berhasil update.")
}

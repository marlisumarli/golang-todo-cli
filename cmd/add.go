package cmd

import (
	"flag"
	"fmt"
	"golang-todo-cli/todo"
	"log"
	"os"
)

func AddTask(todos *todo.Todos, args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTask := addCmd.String("t", "", "Isi dari todo baru")
	addCategory := addCmd.String("c", "Uncategorized", "Category todo")

	addCmd.Parse(args)

	if len(*addTask) == 0 {
		fmt.Println("Error: Butuh -t flag buat jalanin command \"add\".")
		os.Exit(1)
	}

	todos.Add(*addTask, *addCategory)
	err := todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Berhasil nambahin todo.")
}

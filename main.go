package main

import (
	"flag"
	"fmt"
	"golang-todo-cli/cmd"
	"golang-todo-cli/todo"
	"os"
)

func main() {
	todos := &todo.Todos{}
	flag.Parse()

	switch flag.Arg(0) {
	case "init":
		cmd.Init()
	case "add":
		cmd.RemaindInit(todos)
		cmd.AddTask(todos, os.Args[2:])
	case "list":
		cmd.RemaindInit(todos)
		cmd.ListTasks(todos, os.Args[2:])
	default:
		fmt.Println("Invalid command.")
	}
}

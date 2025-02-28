package cmd

import (
	"flag"
	"golang-todo-cli/todo"
)

func ListTasks(todos *todo.Todos, args []string) {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listDone := listCmd.Int("l", 2, "Nampilin status todo")
	listCategory := listCmd.String("c", "", "Nampilin category todo")

	listCmd.Parse(args)
	todos.Print(*listDone, *listCategory)
}

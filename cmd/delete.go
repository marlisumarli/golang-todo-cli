package cmd

import (
	"flag"
	"fmt"
	"golang-todo-cli/todo"
	"log"
)

func DeleteTask(todos *todo.Todos, args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.Int("i", 0, "ID dari task yang mau dihapus")
	deleteAll := deleteCmd.Bool("a", false, "Hapus semua task")

	deleteCmd.Parse(args)

	if *deleteAll {
		err := todos.DeleteAll()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Semua task berhasil dihapus.")
	} else {
		err := todos.Delete(*deleteID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Berhasil dihapus.")
	}

	err := todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
}

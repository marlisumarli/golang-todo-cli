package cmd

import (
	"bufio"
	"fmt"
	"golang-todo-cli/todo"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetJsonFile() string {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(curDir, "todos.json")
}

func RemaindInit(todos *todo.Todos) {
	_, err := os.Stat(GetJsonFile())
	if err != nil {
		fmt.Println("Harus jalanin \"init\" dulu buat bikin JSON file buat nyimpen data.")
		os.Exit(1)
	} else {
		if err := todos.Load(GetJsonFile()); err != nil {
			log.Fatal(err)
		}
	}
}

func GetUserApproval() bool {
	confirmMsg := "Butuh bikin file \"todo.json\" broo, mau lanjut bikin ga? (Y/n) "

	r := bufio.NewReader(os.Stdin)
	var s string

	fmt.Print(confirmMsg)
	s, _ = r.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	for {
		if s == "y" || s == "yes" || s == "" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}

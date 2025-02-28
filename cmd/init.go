package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Init() {
	ok := GetUserApproval()
	if !ok {
		fmt.Print("Lah gamau? yaudah.")
		os.Exit(0)
	}

	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filepath := filepath.Join(curDir, "todos.json")

	_, err = os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(filepath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			fmt.Println("File JSON udah dibuat, coba cek.")
		} else {
			log.Fatal("Error njirr.")
		}
	} else {
		fmt.Println("File JSON udah ada broo.")
	}
}

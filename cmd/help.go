package cmd

import "fmt"

func Help() {
	fmt.Println("\"gm gm gm lfgggg!\"")
	fmt.Println("Usage: todo [command] [arguments...]")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println("  init              - Membuat file JSON untuk menyimpan todo list")
	fmt.Println("  add -t <task> -c <category>  - Menambahkan task baru dengan kategori")
	fmt.Println("  list -l          - Menampilkan semua task dalam todo list")
	fmt.Println("  update <id> -t <task> -c <category>  - Memperbarui task berdasarkan ID")
	fmt.Println("  delete <id>       - Menghapus task berdasarkan ID")
	fmt.Println("  delete -a         - Menghapus semua task")
	fmt.Println("  help              - Menampilkan bantuan")
	fmt.Println()
	fmt.Println("Example Usage:")
	fmt.Println("  todo add -t \"Belajar Go\" -c \"Programming\"")
	fmt.Println("  todo list -l")
	fmt.Println("  todo update 1 -t \"Belajar Golang\" -c \"Coding\"")
	fmt.Println("  todo delete 2")

}

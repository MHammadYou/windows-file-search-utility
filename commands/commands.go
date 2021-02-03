package commands

import (
	"bufio"
	"fmt"
	"os"
)


func HandleCommandInput(files []string, PATH string) {

	scanner := bufio.NewScanner(os.Stdin)

	if len(files) > 0 {
		fmt.Println("")
		fmt.Println("Use below commands")
		fmt.Println("q: to quit")
		fmt.Println("r: to rename files")
		fmt.Println("e: to change extension")
		fmt.Println("d: to delete files")

		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "q":
			os.Exit(0)
		case "d":
			removeFiles(files)
		case "r":
			renameFiles(files, PATH)
		case "e":
			changeExtension(files)
		default:
			fmt.Println("Please enter a valid command")
			HandleCommandInput(files, PATH)
		}
	}
}

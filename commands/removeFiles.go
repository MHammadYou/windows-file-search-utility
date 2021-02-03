package commands

import (
	"bufio"
	"fmt"
	"os"
)

func removeFiles(files []string) {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Files will be deleted permanently and can't be retrieved later from recycle bin. Are you sure to delete these files?")
	fmt.Println("y: yes")
	fmt.Println("n: cancel")

	scanner.Scan()
	command := scanner.Text()

	switch command {
	case "n":
		fmt.Println("Canceled")
	case "y":
		fmt.Println("Deleted files")
		for _, file := range files {
			err := os.Remove(file)

			if err != nil {
				panic(err)
			}
		}
	default:
		fmt.Println("Please choose a valid command")
		removeFiles(files)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func handleCommandInput(files []string) {

	scanner := bufio.NewScanner(os.Stdin)

	if len(files) > 0 {
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
		}
	}
}


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
			log.Fatal(os.Remove(file))
		}
	default:
		fmt.Println("Please choose a valid command")
		removeFiles(files)
	}

}

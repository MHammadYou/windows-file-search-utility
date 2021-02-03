package commands

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func renameFiles(files []string, PATH string) {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter New filename")
	scanner.Scan()
	newName := scanner.Text()

	fmt.Println("Are you sure you want to change filename?")
	fmt.Println("y: yes")
	fmt.Println("n: cancel")

	scanner.Scan()
	command := scanner.Text()

	switch command {
	case "n":
		os.Exit(0)
	case "y":
		for _, file := range files {
			oldExten := path.Ext(file)
			err := os.Rename(file, PATH + newName + oldExten)

			if err != nil {
				panic(err)
			}
		}
		fmt.Println("Filename changed")
	default:
		fmt.Println("Please choose a valid command")
		renameFiles(files, PATH)
	}
}


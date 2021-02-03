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
		for index, file := range files {
			oldExten := path.Ext(file)

			var err error
			if index > 0 {
				newPath := fmt.Sprintf("%s\\%s (%d)%s", PATH, newName, index + 1, oldExten)
				err = os.Rename(file, newPath)
			} else {
				newPath := fmt.Sprintf("%s\\%s%s", PATH, newName, oldExten)
				err = os.Rename(file, newPath)
			}

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


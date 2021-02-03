package commands

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func changeExtension(files []string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter New extension")
	scanner.Scan()
	newExten := scanner.Text()

	fmt.Println("Are you sure you want to change file extension?")
	fmt.Println("y: yes")
	fmt.Println("n: cancel")

	scanner.Scan()
	command := scanner.Text()

	switch command {
	case "n":
		fmt.Println("Canceled")
		os.Exit(0)
	case "y":
		for _, file := range files {
			fileName := strings.TrimSuffix(file, path.Ext(file))
			err := os.Rename(file, fileName + "." + newExten)

			if err != nil {
				panic(err)
			}
		}
		fmt.Println("Changed file extension")
	default:
		fmt.Println("Please enter a valid command")
		changeExtension(files)
	}
}

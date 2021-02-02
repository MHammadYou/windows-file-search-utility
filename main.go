package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter filename (filename can be empty)")
	scanner.Scan()
	filename := strings.ToLower(scanner.Text())

	fmt.Println("Enter Extension (extension can be empty)")
	scanner.Scan()
	extension := strings.ToLower(scanner.Text())

	PATH, _ := os.Getwd()
	var allFiles []string

	err := filepath.Walk(PATH, func(file string, info os.FileInfo, err error) error {

		names := strings.Split(file, PATH)[1]

		if len(names) < 1 {
			return nil
		}
		fullName := removeInitialSlash(names)
		files := strings.Split(fullName, ".")

		name := files[0]
		exten := files[1]

		if filename != "" && extension != "" {
			if strings.Contains(name, filename) && strings.Contains(exten, extension) {
				allFiles = append(allFiles, fullName)
			}
		} else if filename != "" && extension == "" {
			if strings.Contains(name, filename) {
				allFiles = append(allFiles, fullName)
			}
		} else if extension != "" && filename == "" {
			if strings.Contains(exten, extension) {
				allFiles = append(allFiles, fullName)
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	printFiles(allFiles)
}
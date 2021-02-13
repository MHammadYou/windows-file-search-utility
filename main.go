package main

import (
	"WindowsFileSearch/commands"
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

	PATH, err := os.Getwd()
	if err != nil { panic(err) }

	var files []string
	var filesWithFullPath []string

	err = filepath.Walk(PATH, func(file string, info os.FileInfo, err error) error {

		names := strings.Split(file, PATH)[1]

		if len(names) < 1 { return nil }

		fullName := removeInitialSlash(names)
		splitNameList := strings.Split(fullName, ".")

		if len(splitNameList) > 1 {
			name := splitNameList[0]
			exten := splitNameList[1]

			if filename != "" && extension != "" {
				if strings.Contains(name, filename) && strings.Contains(exten, extension) {
					files = append(files, fullName)
					filesWithFullPath = append(filesWithFullPath, file)
				}
			} else if filename != "" && extension == "" {
				if strings.Contains(name, filename) {
					files = append(files, fullName)
					filesWithFullPath = append(filesWithFullPath, file)
				}
			} else if extension != "" && filename == "" {
				if strings.Contains(exten, extension) {
					files = append(files, fullName)
					filesWithFullPath = append(filesWithFullPath, file)
				}
			}
		}

		return nil
	})

	if err != nil { panic(err) }

	printFiles(files)

	commands.HandleCommandInput(filesWithFullPath, PATH)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Print("Enter filename: ")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	filename := scanner.Text()


	var files []string

	root, _ := os.Getwd()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		file := strings.Split(path, root + "\\")[1:]

		if len(file) >= 1 {
			files = append(files, file[0])
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Contains(file, filename) {
			fmt.Println(file)
		}
	}
}

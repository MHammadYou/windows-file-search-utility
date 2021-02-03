package main

import (
	"fmt"
	"strings"
)

func removeInitialSlash(str string) string {

	newStr := ""

	for i := 1; i < len(str); i++ {
		newStr += string(str[i])
	}
	return strings.ToLower(newStr)
}


func printFiles(files []string) {

	if len(files) < 1 {
		fmt.Println("No files found")
	} else {
		fmt.Printf("Found %d file(s)\n", len(files))
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
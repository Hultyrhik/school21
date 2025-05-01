package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
)

func main() {
	oldFile := flag.String("old", "", "Enter path to old file")
	newFile := flag.String("new", "", "Enter path to old file")
	flag.Parse()

	oldFilepaths, err := ReadFile(*oldFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newFilepaths, err := ReadFile(*newFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(oldFilepaths) == 0 && len(newFilepaths) == 0 {
		fmt.Println("Both files are empty")
		os.Exit(1)
	}

	for _, filepath := range oldFilepaths {
		if !slices.Contains(newFilepaths, filepath) {
			fmt.Println("REMOVED", filepath)
		}
	}

	for _, filepath := range newFilepaths {
		if !slices.Contains(oldFilepaths, filepath) {
			fmt.Println("ADDED", filepath)
		}
	}

}

func ReadFile(filename string) ([]string, error) {
	var filepaths []string
	file, err := os.Open(filename)
	if err != nil {
		return filepaths, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filepaths = append(filepaths, scanner.Text())
	}
	err = scanner.Err()
	return filepaths, err
}

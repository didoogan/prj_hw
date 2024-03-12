package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Editor struct {
	lines        []string
	searchResult []string
}

func (e *Editor) loadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var line string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		e.lines = append(e.lines, line)
	}

	err = scanner.Err()

	return err
}

func (e *Editor) search(searchText string) {
	for _, l := range e.lines {
		if strings.Contains(l, searchText) {
			e.searchResult = append(e.searchResult, l)
		}
	}
}

func printLines(lines []string) {
	for _, l := range lines {
		fmt.Println(l)
	}
}

func (e *Editor) printSearchResult() {
	printLines(e.searchResult)
}

func (e *Editor) printLines() {
	printLines(e.lines)
}

func main() {

	const search = "текст"

	editor := Editor{}

	error := editor.loadFromFile("data.txt")
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println("Editor text:")
	editor.printLines()
	editor.search(search)
	fmt.Println("Editor search result:")
	editor.printSearchResult()
}

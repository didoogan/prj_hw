package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFileLines(path string) ([]string, error) {
	var lines []string

	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	var line string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		// Розумію, що можна прям тут робити grep, щоб не ітеруватись другий раз, але single responsebilty і все таке
		lines = append(lines, line)
	}

	err = scanner.Err()

	return lines, err
}

func grep(string1, string2 string) {
	// Print string1 if it contains string2

	if strings.Contains(string1, string2) {
		fmt.Println(string1)
	}
}

func main() {

	const search = "текст"

	lines, err := readFileLines("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	for _, s := range lines {
		grep(s, search)
	}
}

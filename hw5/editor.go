package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"
)

const allEngilishWordsFilePath = "words.txt"
const sourceFilePath = "source.txt"

const lineWordsCopacity = 15

type Editor struct {
	bufferLines  []string
	searchResult []string
	index        map[string][]int
}

func (e *Editor) loadBufferFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var line string

	scanner := bufio.NewScanner(file)
	e.bufferLines = make([]string, 0)
	for scanner.Scan() {
		line = scanner.Text()
		e.bufferLines = append(e.bufferLines, line)
	}

	err = scanner.Err()

	return err
}

func (e *Editor) search(searchText string) {
	e.searchResult = make([]string, 0)

	for _, l := range e.bufferLines {
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

func (e *Editor) printBuffer() {
	printLines(e.bufferLines)
}

func (e *Editor) GenerateSourceFile(lines []string) error {
	f, err := os.Create(sourceFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range lines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func shuffle(lines []string) {
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
}

func (e *Editor) GenerateLines(filePath string, linesCount, uniqWordsCount int) ([]string, error) {
	var result []string
	var lineList []string

	err := e.loadBufferFromFile(filePath)
	if err != nil {
		return result, err
	}

	shuffle(e.bufferLines)

	for i := 0; i < linesCount; i++ {
		lineList = make([]string, lineWordsCopacity)
		for y := 0; y < lineWordsCopacity; y++ {
			randomIndex := rand.Intn(uniqWordsCount)
			lineList = append(lineList, e.bufferLines[randomIndex])
		}
		result = append(result, strings.Join(lineList, " "))
	}

	return result, nil
}

func timeCount(f func(string), arg string) time.Duration {
	startTime := time.Now()
	f(arg)
	elapsedTime := time.Since(startTime)
	return elapsedTime
}

func (e *Editor) makeIndex() {
	e.index = make(map[string][]int)

	for lineNumber, line := range e.bufferLines {
		for _, word := range strings.Split(line, " ") {
			if _, ok := e.index[word]; !ok {
				e.index[word] = make([]int, 1)
			}
			if !slices.Contains(e.index[word], lineNumber) {
				e.index[word] = append(e.index[word], lineNumber)
			}
		}
	}
}

func (e *Editor) searchIndex(searchText string) {
	e.searchResult = make([]string, 0)

	for _, lineNumber := range e.index[searchText] {
		e.searchResult = append(e.searchResult, e.bufferLines[lineNumber])
	}
}

func compareSearches(editor *Editor, linesCount, uniqWordsCount int) {

	lines, err := editor.GenerateLines(allEngilishWordsFilePath, linesCount, uniqWordsCount)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = editor.GenerateSourceFile(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = editor.loadBufferFromFile(sourceFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	searchWord := strings.Split(lines[0], " ")[0]
	searchTime := timeCount(editor.search, searchWord)

	startTime := time.Now()
	editor.makeIndex()
	elapsedIndexTime := time.Since(startTime)

	indexSearchTime := timeCount(editor.searchIndex, searchWord)
	fmt.Printf("Income params: %v lines, %v uniq words. Search time: %v ns, index search time: %v ns (index time: %v ns)\n", linesCount, uniqWordsCount, searchTime.Nanoseconds(), indexSearchTime.Nanoseconds(), elapsedIndexTime.Nanoseconds())
	fmt.Println("================================================================================================")
}

func main() {

	editor := Editor{}

	presets := [3][2]int{
		{10, 50},
		{100, 500},
		{10000, 1000},
	}

	for _, preset := range presets {
		linesCount := preset[0]
		uniqWordsCount := preset[1]

		compareSearches(&editor, linesCount, uniqWordsCount)
	}

}

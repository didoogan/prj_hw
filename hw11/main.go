package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const simpleNumberPatternStr = `^\d{10}`
const numberWithParenthesesPatternStr = `^\(\d{3}\)[\s]?\d{3}-\d{4}$`
const delimitedNumberPatternStr = `^\d{3}[\s\.-]{1}\d{3}[\s\.-]{1}\d{4}$`

type SearchEngine struct {
	filePath string
	tokens   []string
	result   []string
}

func (e *SearchEngine) showTokens() {
	for _, w := range e.tokens {
		fmt.Println(w)
	}
}

func (e *SearchEngine) clearTokens() {
	e.tokens = make([]string, 0)
}

func (e *SearchEngine) clearResult() {
	e.result = make([]string, 0)
}

func (e *SearchEngine) extractTokensFromFile(callBack func(e *SearchEngine, fileLine string)) error {
	e.clearTokens()

	f, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		fileLine := scn.Text()

		callBack(e, fileLine)
	}

	if err := scn.Err(); err != nil {
		return err
	}

	return nil
}

func (e *SearchEngine) extractLinesFromFile() error {
	cb := func(e *SearchEngine, fileLine string) {
		e.tokens = append(e.tokens, fileLine)
	}

	return e.extractTokensFromFile(cb)
}

func (e *SearchEngine) extractWordsFromFile() error {
	redundantSuffixes := []string{".", ",", ";", ":", "!", "?"}

	cb := func(e *SearchEngine, fileLine string) {

		for _, w := range strings.Fields(fileLine) {
			for _, s := range redundantSuffixes {
				if strings.HasSuffix(w, s) {
					w = strings.TrimSuffix(w, s)
				}
			}
			e.tokens = append(e.tokens, w)
		}
	}

	return e.extractTokensFromFile(cb)
}

func (e *SearchEngine) searchByPattern(re *regexp.Regexp) {
	e.clearResult()

	for _, w := range e.tokens {
		isValid := re.MatchString(w)

		if isValid {
			e.result = append(e.result, w)
		}
	}
}

func (e *SearchEngine) showResult(pattern string) {
	fmt.Printf("Found %v result(s) for pattern `%v` in the file %v:\n", len(e.result), pattern, e.filePath)

	for _, w := range e.result {
		fmt.Println(w)
	}

	fmt.Println("==================================")
}

func main() {
	e := SearchEngine{
		filePath: "numbers.txt",
		tokens:   make([]string, 0),
		result:   make([]string, 0),
	}

	err := e.extractLinesFromFile()
	if err != nil {
		fmt.Println(err)
	}

	e.showTokens()

	numbersPatterns := []string{simpleNumberPatternStr, numberWithParenthesesPatternStr, delimitedNumberPatternStr}

	for _, p := range numbersPatterns {
		re := regexp.MustCompile(p)
		e.searchByPattern(re)
		e.showResult(p)
	}

	e.filePath = "text.txt"
	e.extractWordsFromFile()

	ukraineVowelsChars := "аеиіоуяюїєАЕИІОУЯЮЇЄ"
	ukraineConsonantsCharsR := "бвгджзйклмнпрстфхцчшщБВГДЖЗЙКЛМНПРСТФХЦЧШЩ"

	patterns := make([]string, 0)

	pattern1 := fmt.Sprintf("^[%v][%v%v]*[%v]$", ukraineConsonantsCharsR, ukraineConsonantsCharsR, ukraineVowelsChars, ukraineVowelsChars)
	pattern2 := fmt.Sprintf("^[%v][%v%v]*[%v]$", ukraineVowelsChars, ukraineConsonantsCharsR, ukraineVowelsChars, ukraineConsonantsCharsR)
	pattern3 := fmt.Sprintf("^[%v][%v%v]*[%v]$", ukraineVowelsChars, ukraineConsonantsCharsR, ukraineVowelsChars, ukraineVowelsChars)
	pattern4 := fmt.Sprintf("^[%v][%v%v]*[%v]$", ukraineVowelsChars, ukraineConsonantsCharsR, ukraineVowelsChars, ukraineVowelsChars)
	pattern5 := `^(\w)\w*\1$` // unfortunately golang doesn't understand \1 :(

	patterns = append(patterns, pattern1, pattern2, pattern3, pattern4, pattern5)

	for _, p := range patterns {
		re, err := regexp.Compile(p)

		if err != nil {
			fmt.Printf("Pattern %v is not valid reqular expression: %v\n", p, err)
			return
		}

		e.searchByPattern(re)
		e.showResult(p)
	}
}

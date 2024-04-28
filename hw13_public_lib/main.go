package main

import (
	"fmt"
	"github.com/didoogan/searchphonenumber"
	"regexp"
)

func main() {
	e := searchphonenumber.SearchEngine{
		FilePath: "numbers.txt",
		Tokens:   make([]string, 0),
		Result:   make([]string, 0),
	}

	err := e.ExtractLinesFromFile()
	if err != nil {
		fmt.Println(err)
	}

	e.ShowTokens()

	numbersPatterns := []string{
		searchphonenumber.SimpleNumberPatternStr,
		searchphonenumber.NumberWithParenthesesPatternStr,
		searchphonenumber.DelimitedNumberPatternStr}

	for _, p := range numbersPatterns {
		re := regexp.MustCompile(p)
		e.SearchByPattern(re)
		e.ShowResult(p)
	}

	e.FilePath = "text.txt"
	e.ExtractWordsFromFile()

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

		e.SearchByPattern(re)
		e.ShowResult(p)
	}
}

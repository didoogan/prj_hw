package store

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const keyValueDelimiter = ":"
const defaultFilePath = "db.txt"

type File struct {
	Path string
}

func (s *File) Save(key, value string) error {
	f, err := os.OpenFile(s.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	defer f.Close()

	keyValueStr := fmt.Sprintf("%v%v%v\n", key, keyValueDelimiter, value)
	if _, err := f.WriteString(keyValueStr); err != nil {
		return err
	}
	fmt.Printf("%v saved into db\n", key)
	return nil
}

func (s *File) searchForKey(key string) (string, bool, error) {
	f, err := os.Open(s.Path)
	if err != nil {
		return "", false, err
	}
	defer f.Close()

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		line := fs.Text()
		searchPrefix := fmt.Sprintf("%v%v", key, keyValueDelimiter)
		if strings.HasPrefix(line, searchPrefix) {
			return line[len(searchPrefix):], true, nil
		}
	}

	if err := fs.Err(); err != nil {
		return "", false, err
	}

	return "", false, nil
}

func (s *File) GetAllKeys() ([]string, error) {
	keys := make([]string, 0)
	f, err := os.Open(s.Path)
	if err != nil {
		return keys, err
	}
	defer f.Close()

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		line := fs.Text()
		keyValue := strings.Split(line, keyValueDelimiter)

		keys = append(keys, keyValue[0])
	}

	if err := fs.Err(); err != nil {
		return keys, err
	}

	return keys, nil
}

func (s *File) Get(key string) (string, error) {
	value, exists, err := s.searchForKey(key)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", errors.New("not found")
	}
	return value, nil
}

func (s *File) HasKey(key string) (bool, error) {
	_, found, err := s.searchForKey(key)
	if err != nil {
		return false, err
	}
	return found, nil
}

func NewFileStore() *File {
	return &File{
		Path: defaultFilePath,
	}
}
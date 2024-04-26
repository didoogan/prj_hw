package passwordService

import (
	"errors"
	"fmt"
	argumentSrv "hw12/internal/argument-service"
)

const actionGet = "get"
const actionAdd = "add"

const passwordMinLength = 6
const passwordMaxLength = 20

type Store interface {
	Save(key, value string) error
	Get(key string) (string, error)
	HasKey(key string) (bool, error)
	GetAllKeys() ([]string, error)
}

type PasswordSrv struct {
	Store       Store
	ArgumentSrv *argumentSrv.ArgumentSrv
}

func (s *PasswordSrv) validatePassword(password *string) bool {
	if password == nil {
		return false
	}

	if len(*password) < passwordMinLength || len(*password) > passwordMaxLength {
		fmt.Printf("Your password shoud be more than %v and less than %v simbols\n", passwordMinLength, passwordMaxLength)
		return false
	}

	return true
}

func (s *PasswordSrv) Save(passwordName string) error {
	alreadyExists, err := s.Store.HasKey(passwordName)
	if err != nil {
		return err
	}

	if alreadyExists {
		return errors.New("already exists")
	}

	var passwordValue string

	for {
		fmt.Print("Enter password: ")

		fmt.Scan(&passwordValue)

		if !s.validatePassword(&passwordValue) {
			continue
		}

		err = s.Store.Save(passwordName, passwordValue)

		if err != nil {
			return err
		}

		fmt.Printf("Password for %v successfully saved\n", passwordName)

		return nil
	}

	return nil
}

func (s *PasswordSrv) ShowPassword(passwordName string) error {
	passwordValue, err := s.Store.Get(passwordName)

	if err != nil {
		return err
	}

	fmt.Printf("Password for %v is: %v\n", passwordName, passwordValue)

	return nil
}

func (s *PasswordSrv) ShowPasswordsNames() error {
	keys, err := s.Store.GetAllKeys()

	if err != nil {
		return err
	}

	fmt.Printf("Found %v key(s)\n", len(keys))

	for i, k := range keys {
		fmt.Printf("%v. %v\n", i+1, k)
	}

	return nil
}

func (s *PasswordSrv) ProcessRequest() error {
	argsLen := s.ArgumentSrv.GetArgumentsLen()

	switch argsLen {
	case 0:
		return s.ShowPasswordsNames()
	case 2:
		args := s.ArgumentSrv.GetArguments()
		action := args[0]
		passwordName := args[1]

		if action != actionGet && action != actionAdd {
			return errors.New("first argument should be 'get' or 'save'")
		}

		switch action {
		case actionGet:
			return s.ShowPassword(passwordName)

		case actionAdd:
			return s.Save(passwordName)
		}
	default:
		return errors.New("you should pass zero or two arguments")
	}

	return nil
}

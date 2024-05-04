package arguments

import (
	"errors"
	"os"
)

type Service struct {
	args []string
}

func (s *Service) GetArguments() []string {
	args := make([]string, 0)

	for _, arg := range s.args {
		args = append(args, arg)
	}

	return args
}

func (s *Service) GetArgumentsLen() int {
	return len(s.args)
}

func (s *Service) GetArgument(index int) (string, error) {
	if index >= s.GetArgumentsLen() {
		return "", errors.New("index error")
	}

	return s.args[index], nil
}

func NewService() *Service {
	return &Service{args: os.Args[1:]}
}

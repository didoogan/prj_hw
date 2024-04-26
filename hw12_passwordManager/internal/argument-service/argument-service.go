package argument_service

import (
	"errors"
	"os"
)

type ArgumentSrv struct {
	args []string
}

func (s *ArgumentSrv) GetArguments() []string {
	args := make([]string, 0)

	for _, arg := range s.args {
		args = append(args, arg)
	}

	return args
}

func (s *ArgumentSrv) GetArgumentsLen() int {
	return len(s.args)
}

func (s *ArgumentSrv) GetArgument(index int) (string, error) {
	if index >= s.GetArgumentsLen() {
		return "", errors.New("index error")
	}

	return s.args[index], nil
}

func NewArgumentSrv() *ArgumentSrv {
	return &ArgumentSrv{args: os.Args[1:]}
}

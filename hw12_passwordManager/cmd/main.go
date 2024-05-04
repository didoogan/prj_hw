package main

import (
	"fmt"
	"hw12/internal/arguments"
	"hw12/internal/password"
	"hw12/internal/store"
)

func main() {

	argService := arguments.NewService()

	fileStore := store.NewFileStore()

	passwordSrv := password.Service{Store: fileStore, ArgumentService: argService}

	err := passwordSrv.ProcessRequest()

	if err != nil {
		fmt.Println(err)
	}
}

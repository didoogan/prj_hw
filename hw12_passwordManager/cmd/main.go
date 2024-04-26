package main

import (
	"fmt"
	argumentSrv "hw12/internal/argument-service"
	"hw12/internal/store"
	passwordService "hw12/pkg/password-service"
)

func main() {

	argSrv := argumentSrv.NewArgumentSrv()

	fileStore := store.NewFileStore()

	passwordSrv := passwordService.PasswordSrv{Store: fileStore, ArgumentSrv: argSrv}

	err := passwordSrv.ProcessRequest()

	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"os"

	"MSaaS-Framework/MSaaS/cmd/wizcraft/app"
	"MSaaS-Framework/MSaaS/pkg/base"
)

func main() {

	command := app.NewWizcraftCommand()
	code := base.Run(command)
	os.Exit(code)
}

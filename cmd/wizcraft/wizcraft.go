package main

import (
	"Larshavin/MSaaS/cmd/wizcraft/app"
	"Larshavin/MSaaS/pkg/base"
	"os"
)

func main() {

	command := app.NewWizcraftCommand()
	code := base.Run(command)
	os.Exit(code)
}

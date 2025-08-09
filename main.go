package main

import (
	"Ewallet-grpc/cmd"
	"Ewallet-grpc/helpers"
)

func main() {
	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// helpers.SetupMySql()

	// running grpc
	go cmd.ServeGRPC()

	cmd.ServeHTTP()
}

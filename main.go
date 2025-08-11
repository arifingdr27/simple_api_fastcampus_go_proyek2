package main

import (
	"ewallet-ums/cmd"
	"ewallet-ums/helpers"
)

func main() {
	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	helpers.SetupMySql()

	// running grpc
	go cmd.ServeGRPC()

	cmd.ServeHTTP()
}

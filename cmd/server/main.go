package main

import (
	"github.com/dopaemon/artus/internal/auth"
	"github.com/dopaemon/artus/internal/cli"
	"github.com/dopaemon/artus/internal/db"
	"github.com/dopaemon/artus/internal/server"
)

func main() {
	db.InitDB()

	if !auth.RegisterOrLogin() {
		return
	}

	server.StartGRPCServer(":50051")

	cli.ShowMetricsCLI()
}

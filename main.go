package main

import (
	"embed"
	"fmt"

	"github.com/andig/evcc/cmd"
	"github.com/andig/evcc/server"
)

//go:embed dist/*
var assets embed.FS

func init() {
	// use embedded assets unless live assets are already loaded
	if server.Assets == nil {
		fmt.Println("--- EMBED ---")
		server.Assets = assets
	}
}

func main() {
	cmd.Execute()
}

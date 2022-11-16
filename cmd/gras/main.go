package main

import (
	"context"
	"log"
	"os"

	"github.com/remiposo/gras/cmd/gras/subcmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gras",
		Usage: "manage gras",
		Commands: []*cli.Command{
			subcmd.NewServer(),
		},
	}
	if err := app.RunContext(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

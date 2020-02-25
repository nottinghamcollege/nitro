package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"sort"

	"github.com/urfave/cli/v2"

	"github.com/pixelandtonic/nitro/action"
	"github.com/pixelandtonic/nitro/command"
)

func main() {
	executor := action.NewSyscallExecutor("multipass")

	app := &cli.App{
		Name:        "nitro",
		Usage:       "Develop Craft CMS websites locally with ease",
		Version:     "1.0.0",
		Description: "A better way to develop Craft CMS applications without Docker or Vagrant",
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "machine",
				Aliases:     []string{"m"},
				Value:       "nitro-dev",
				Usage:       "Provide a machine name",
				DefaultText: "nitro-dev",
			},
		},
		Commands: []*cli.Command{
			command.Initialize(),
			command.SSH(executor),
			command.Delete(),
			command.Stop(),
		},
	}

	// find the path to multipass and set value in context
	multipass, err := exec.LookPath("multipass")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.WithValue(context.Background(), "multipass", multipass)

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}

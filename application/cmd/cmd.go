package cmd

import (
	"github.com/urfave/cli"
	"os"
)

func Start() {
	app := cli.NewApp()

	app.Name = "gih"
	app.Usage = "Github Client CLI."
	app.Version = "0.0.1"

	app.Action = action

	app.Flags = []cli.Flag{
		&cli.BoolFlag{Name: "config", Aliases: []string{"c"}},
	}

	_ = app.Run(os.Args)
}

func action(context *cli.Context) error {

	if context.Bool("config") {

	}

	return nil
}



package cmd

import (
	"fmt"
	"gih/registry"
	"github.com/urfave/cli/v2"
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
		&cli.BoolFlag{Name: "issues", Aliases: []string{"i"}},
	}

	_ = app.Run(os.Args)
}

func action(context *cli.Context) error {

	if context.Bool("config") {
		editConfig()
		return nil
	}

	config := readConfig()
	s := registry.NewGitService(config.AccessToken, config.RepositoryName, config.Organization)
	fmt.Println(s.GetIssues())

	return nil
}

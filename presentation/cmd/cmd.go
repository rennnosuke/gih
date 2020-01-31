package cmd

import (
	"fmt"
	action2 "github.com/rennnosuke/gih/presentation/cmd/action"
	"github.com/rennnosuke/gih/registry"
	"github.com/urfave/cli/v2"
	"os"
)

func Start() {
	app := cli.NewApp()

	app.Name = "gih"
	app.Usage = "GitHub Client CLI."
	app.Version = "0.0.1"

	app.Action = action

	app.Flags = []cli.Flag{
		&cli.BoolFlag{Name: "init"},
		&cli.BoolFlag{Name: "create", Aliases: []string{"c"}},
		&cli.BoolFlag{Name: "update", Aliases: []string{"u"}},
		&cli.BoolFlag{Name: "close", Aliases: []string{"d"}},
		&cli.BoolFlag{Name: "browse", Aliases: []string{"w"}},
	}

	_ = app.Run(os.Args)
}

func action(context *cli.Context) error {

	if context.Args().Get(0) == "init" {
		editConfig()
		return nil
	}

	config := readConfig()
	if config == nil {
		fmt.Println("Configuration is not found.")
		fmt.Println("Execute `$ gih init` to set git hosting service configuration.")
		os.Exit(1)
	}

	if context.Bool("browse") {
		return browse(config.RepositoryPath)
	}

	s := registry.NewGitService(config.AccessToken, config.RepositoryName, config.Organization)

	if context.Bool("create") {
		return action2.createIssue(s, context)
	}

	if context.Bool("update") {
		return action2.updateIssue(s, context)
	}

	if context.Bool("close") {
		return action2.closeIssue(s, context)
	}

	return action2.listIssues(s)
}

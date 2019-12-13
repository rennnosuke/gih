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
		&cli.BoolFlag{Name: "browse", Aliases: []string{"w"}},
	}

	_ = app.Run(os.Args)
}

func action(context *cli.Context) error {

	if context.Bool("config") {
		editConfig()
		return nil
	}

	config := readConfig()
	if config == nil {
		fmt.Println("Configuration is not found.")
		fmt.Println("Execute `$ gih -c` to set git hosting service configuration.")
		os.Exit(1)
	}

	if context.Bool("browse") {
		err := openWebBrowser(config.RepositoryPath)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		return nil
	}

	s := registry.NewGitService(config.AccessToken, config.RepositoryName, config.Organization)
	fmt.Println(s.GetIssues())

	return nil
}

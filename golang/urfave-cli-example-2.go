package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := initApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initApp() *cli.App {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "Language for the greeting",
				EnvVars: []string{"LEGACY_COMPAT_LANG", "APP_LANG", "LANG"},
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
		Commands: []*cli.Command{
			cmdAdd,
			cmdComplete,
			cmdTemplate,
		},
	}

	return app
}

// Command Info
var cmdAdd = &cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "add a task to the list",
	Action:  doAdd,
}

var cmdComplete = &cli.Command{
	Name:    "complete",
	Aliases: []string{"c"},
	Usage:   "complete a task on the list",
	Action:  doComplete,
}

var cmdTemplate = &cli.Command{
	Name:    "template",
	Aliases: []string{"t"},
	Usage:   "options for task templates",
	Subcommands: []*cli.Command{
		cmdTemplateAdd,
		cmdTemplateRemove,
	},
}

var cmdTemplateAdd = &cli.Command{
	Name:   "add",
	Usage:  "add a new template",
	Action: doTemplateAdd,
}

var cmdTemplateRemove = &cli.Command{
	Name:   "remove",
	Usage:  "remove an existing template",
	Action: doTemplateRemove,
}

// Command Action
func doAdd(c *cli.Context) error {
	fmt.Println("added task: ", c.Args().First())
	return nil
}

func doComplete(c *cli.Context) error {
	fmt.Println("completed task: ", c.Args().First())
	return nil
}

func doTemplateAdd(c *cli.Context) error {
	fmt.Println("new task template: ", c.Args().First())
	return nil
}

func doTemplateRemove(c *cli.Context) error {
	fmt.Println("removed task template: ", c.Args().First())
	return nil
}

package main

import  (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app:= cli.NewApp()
	app.Name = "subcommand_client"
	app.Usage = "github.com/codegangsta/cliを試してます"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.BoolFlag {
			Name: "dryrun d",
			Usage: "グローバルオプション dry runです",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "hello",
			Aliases: []string{"h"},
			Usage:  "hello worldを表示します",
			Action: helloAction,
		},
		{
			Name:   "fuga",
			Aliases: []string{"f"},
			Usage:  "fugaを表示します",
			Action: fugaAction,
		},
	}

	app.Before = func(c *cli.Context) error {
		fmt.Println("開始")
		return nil
	}

	app.After = func(c *cli.Context) error {
		fmt.Println("終了")
		return nil
	}

	app.Run(os.Args)
}

func helloAction(c *cli.Context) {
	var isDry = c.GlobalBool("dryrun")
	if isDry {
		fmt.Println("this is dry-run")
	}

	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First()
	}

	fmt.Printf("Hello world! %s\n", paramFirst)
}

func fugaAction(c *cli.Context) {
	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First()
	}

	fmt.Printf("Fuga! %s\n", paramFirst)
}

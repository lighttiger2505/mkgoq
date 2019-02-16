package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	err := newApp().Run(os.Args)
	var exitCode = 0
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		exitCode = 255
	}
	os.Exit(exitCode)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "mkgoq"
	app.HelpName = "mkgoq"
	app.Usage = "markdown query."
	app.Version = "0.0.1"
	app.Author = "lighttiger2505"
	app.Email = "lighttiger2505@gmail.com"
	app.Action = run
	return app
}

func run(c *cli.Context) error {
	return nil
}

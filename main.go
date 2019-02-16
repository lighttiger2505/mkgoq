package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	"github.com/lighttiger2505/mkgoq/lib/markdown"
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
	if len(c.Args()) == 0 {
		return fmt.Errorf("The required arguments were not provided: <markdown file>")
	}
	fpath := c.Args().First()

	if !isExists(fpath) {
		return fmt.Errorf("not found markdown file: %s", fpath)
	}

	file, err := os.OpenFile(fpath, os.O_RDONLY, 0666)
	if err != nil {
		return fmt.Errorf("cannot open file, %s", err.Error())
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("cannot read file, %s", err.Error())
	}

	fname := filepath.Base(file.Name())
	abspath, err := filepath.Abs(file.Name())
	if err != nil {
		return err
	}

	headers := markdown.ParseHeader(string(b))
	for _, header := range headers {
		header.Name = fname
		header.Path = abspath
		fmt.Println(fmt.Sprintf(
			"%s((%d, %d), (%d, %d)): %s",
			header.Path,
			header.StPos.Line,
			header.StPos.Row,
			header.EnPos.Line,
			header.EnPos.Row,
			header.RowString,
		))
	}

	return nil
}

func isExists(filename string) bool {
	_, err := os.Stat(filename)

	if pathError, ok := err.(*os.PathError); ok {
		if pathError.Err == syscall.ENOTDIR {
			return false
		}
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func readContents(fpath string) (string, error) {
	file, err := os.OpenFile(fpath, os.O_RDONLY, 0666)
	if err != nil {
		return "", fmt.Errorf("cannot open file, %s", err.Error())
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("cannot read file, %s", err.Error())
	}
	return string(b), nil
}


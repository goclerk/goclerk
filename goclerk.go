package main

import (
	"os"

	"github.com/jonaswouters/goclerk/cmd"

	"github.com/codegangsta/cli"
)

const APP_VER = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "GoClerk"
	app.Usage = "Accounting & More"
	app.Version = APP_VER
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.Setup,
		cmd.Migrate,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}

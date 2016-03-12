package main

import (
	"os"
	"runtime"

	"github.com/goclerk/goclerk/cmd"

	"github.com/codegangsta/cli"
)

const APP_VER = "0.0.1"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "GoClerk"
	app.Usage = "Accounting & More"
	app.Version = APP_VER
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		cmd.Web,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}

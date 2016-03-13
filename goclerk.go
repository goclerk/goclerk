package main

import (
	"os"
	"runtime"

	"github.com/goclerk/goclerk/cmd"

	"github.com/codegangsta/cli"
	"github.com/goclerk/goclerk/modules/setting"
)

const APP_VER = "0.0.1"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	setting.Connection.Username =  "jonaswouters"
	//setting.Connection.Database = "goclerk"
	app := cli.NewApp()
	app.Name = "GoClerk"
	app.Usage = "Accounting & More"
	app.Version = APP_VER
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.Migrate,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}

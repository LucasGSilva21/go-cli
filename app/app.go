package app

import "github.com/urfave/cli"

// generate cli app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line interface"
	app.Usage = "Find Ip's and server names"

	return app
}

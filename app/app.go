package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// generate cli app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line interface"
	app.Usage = "Find Ip's and server names"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Find Ip's",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
				},
			},
			Action: findIpsAndPrint,
		},
	}

	return app
}

func findIpsAndPrint(context *cli.Context) {
	host := context.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

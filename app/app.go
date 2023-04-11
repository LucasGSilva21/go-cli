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

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find Ip's",
			Flags:  flags,
			Action: findIpsAndPrint,
		},
		{
			Name:   "server",
			Usage:  "Find Servers",
			Flags:  flags,
			Action: findServersAndPrint,
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

func findServersAndPrint(context *cli.Context) {
	host := context.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

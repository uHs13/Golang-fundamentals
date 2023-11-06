package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "Command Line Application"
	app.Usage = "Search for IPs and server names"

	app.Commands = configureCommands()

	return app
}

func configureCommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs from internet addresses",
			Flags:  configureFlags(),
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search server info from internet addresses",
			Flags:  configureFlags(),
			Action: searchServers,
		},
	}

	return commands
}

func configureFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	return flags
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

package main

import "github.com/liamnaddell/LDCSP"
import "github.com/urfave/cli"

//import "fmt"
import "os"

var Version string

func main() {
	app := cli.NewApp()
	app.Name = "ldcsp"
	app.Usage = "initiate a conversation with another computer regardless of operating system"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:  "join",
			Usage: "join another server",
			Action: func(c *cli.Context) error {
				ldcsp.JoinServer()
				return nil
			},
		},
		{
			Name:  "host",
			Usage: "host a server",
			Action: func(c *cli.Context) error {
				ldcsp.HostServer()
				return nil
			},
		},
	}
	app.Run(os.Args)
}

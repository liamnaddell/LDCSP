package main

import "github.com/liamnaddell/LDCSP/ldcsp/nhelo"
import "github.com/liamnaddell/LDCSP/ldcsp/lnet"
import "fmt"
import "os"

const (
	usage = `
USAGE: ldcsp [ join | host | help { command } ]
`
	usage_join = `
	USAGE FOR JOIN: 'ldcsp join'
	'ldcsp join' joins port 9000 on a seperate computer and gives you a basic interface for talking with the person on the other end of the port
	ldcsp is cross platform.
	`
	usage_host = `
	USAGE FOR HANDLE: 'ldcsp host'
	'ldcsp handle' handles a server on port 9000, then waits for other people to join you.
	It is cross platform.
	`
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		os.Exit(0)
	}
	for i := 1; i < len(os.Args); i++ {
		switch arg := os.Args[i]; {
		case arg == "help":
			if len(os.Args) == 2 {
				fmt.Println(usage)
				os.Exit(0)
			}
			switch arg2 := os.Args[i+1]; {
			case arg2 == "join":
				fmt.Println(usage_join)
			case arg2 == "host":
				fmt.Println(usage_host)
		case arg == "join":
				lnet.JoinServer()
		case arg == "host":
				nhelo.HostServer()
			}
		}
	}
}

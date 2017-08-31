package main

import "github.com/liamnaddell/LDCSP"

//import "fmt"
import "os"

func main() {
	arg1 := os.Args[1]
	if arg1 == "join" {
		ldcsp.JoinServer()
	} else {
		ldcsp.HostServer()
	}
}

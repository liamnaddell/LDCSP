package main

import "net"
import "time"

func main() {
	ln, err := net.Listen("tcp", ":9000")
	checkerr(err)
	time.Sleep(4 * time.Second)
	con1, err := ln.Accept()
	time.Sleep(1 * time.Second)
	b := make([]byte, 1024)
	con1.Read(b)
	con1.Write(b)
}
func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import "net"
import "time"
import "fmt"

var conns []net.Conn

func main() {
	ln, err := net.Listen("tcp", ":9000")
	checkerr(err)

	//accepter
	go func() {
		for {
			conn, err := ln.Accept()
			if err == nil {
				conns = append(conns, conn)
			}
		}
	}()

	time.Sleep(5 * time.Second)

	fmt.Println(conns[0].LocalAddr())
	go read()
	for {
		time.Sleep(1 * time.Second)
	}
}

func read() {
	for {
		b := make([]byte, 1024)
		for _, i := range conns {
			n, _ := i.Read(b)
			if n > 0 {
				buf := make([]byte, n)
				copy(buf, b[:n])
				fmt.Println(string(buf))
			}
		}
	}
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

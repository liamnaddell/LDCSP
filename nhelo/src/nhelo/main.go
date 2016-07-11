package main

import "net"
import "time"
import "fmt"

var conns []net.Conn

func main() {
	serv, err := net.Listen("tcp", ":8382")
	if err != nil {
		panic(err)
	}
	hcon, _ := net.Dial("tcp", "skilstak.sh:8382")
	conns = append(conns, hcon)
	go getcon(serv)
	go func() {
		b := make([]byte, 1024)

		for {
			n, _ := conn.Read(b)
			if n > 0 {
				res := make([]byte, n)

				copy(res, b[:n])
				fmt.Println(string(res))
			}
		}
	}()
	for {
		time.Sleep(1 * time.Second)
	}
	fmt.Println("endin")
}

func getcon(serv net.Listener) {
	for {
		fmt.Println("accepting")
		conn, err := serv.Accept()
		if err == nil {
			conns = append(conns, conn)
		}

	}
}

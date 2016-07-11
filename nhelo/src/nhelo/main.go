package main

import (
	"fmt"
	//	"io"
	//	"io/ioutil"
	"net"
	"time"
)

type Client struct {
	LocalAdr net.Addr
	Conn     net.Conn
}

func main() {
	//create server
	ln, err := net.Listen("tcp", ":8452")
	if err != nil {
		panic(err)
	}
	//get 2 conns
	con1, con2 := get2conn(ln)

	defer ln.Close()

	fmt.Println(con1.LocalAdr)
	fmt.Println(con2.LocalAdr)
	pipe(con1.Conn, con2.Conn)
	for {
		fmt.Println("notext")
		time.Sleep(10 * time.Second)
	}
}

func pipe(conn, conn2 net.Conn) {
	chan1 := chancon(conn)
	chan2 := chancon(conn2)
	go func() {
		for {
			select {
			case b1 := <-chan1:
				if b1 == nil {
					return
				} else {
					conn2.Write(b1)
				}
			case b2 := <-chan2:
				if b2 == nil {
					return
				} else {
					conn.Write(b2)
				}
			}

		}
	}()
}

func chancon(conn net.Conn) chan []byte {
	c := make(chan []byte)

	go func() {
		b := make([]byte, 1024)

		for {
			n, err := conn.Read(b)
			if n > 0 {
				res := make([]byte, n)
				copy(res, b[:n])
				c <- res
			}
			if err != nil {
				c <- nil
				fmt.Println("stopping read")
				break
			}
		}
	}()
	return c
}

func get2conn(ln net.Listener) (Client, Client) {
	var fincon1 net.Conn
	var fincon2 net.Conn
	for sucA := 0; sucA < 2; {
		//get conn1
		for sucA < 1 {
			var Conn1, err = ln.Accept()
			if err != nil {
				fmt.Println("no conn: ", err)
			} else {
				if err == nil {
					fincon1 = Conn1
					sucA++
				}
			}
		}
		for sucA < 2 {
			var Conn2, err = ln.Accept()
			if err != nil {
				fmt.Println("no conn2: ", err)
			} else {
				if err == nil {
					fincon2 = Conn2
					sucA++
				}
			}
		}
	}
	fc1 := Client{fincon1.LocalAddr(), fincon1}
	fc2 := Client{fincon2.LocalAddr(), fincon2}
	return fc1, fc2
}

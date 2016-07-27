package main

import "net"
import "fmt"
import "github.com/pelletier/go-toml"
import "io/ioutil"
import "bytes"

var msgs = make(chan string)

type Client struct {
	Ch chan<- string
}

var clients []Client

func main() {
	//delete me senpi
	h, err1 := toml.LoadFile("h.toml")
	checkerr(err1)
	fmt.Println(h.Get("h").(string))
	//delete me senpi
	ln, err := net.Listen("tcp", ":9000")
	checkerr(err)
	go broad()
	//accepter
	for {
		conn, err := ln.Accept()
		conn.Write([]byte("accepted"))
		if err != nil {
			panic(err)

		} else {
			fmt.Println("inHandle")
			go handleIn(conn)
		}
	}
}

func broad() {
	for {
		if b := <-msgs; b != "" {
			for _, cl := range clients {
				fmt.Println(b)
				v, err := toml.Load(b)
				ioutil.WriteFile("esrr.txt", []byte(fmt.Sprintf("%v", err)), 0644)

				checkerr(err)
				g := v.Get("msg")

				cl.Ch <- g.(string)
			}
		}
	}
}

func handleIn(conn net.Conn) {
	ch := make(chan string)
	var User = Client{ch}
	clients = append(clients, User)
	go clientWriter(conn, ch)
	for {
		b := make([]byte, 1024)
		n, _ := conn.Read(b)
		if n > 0 {
			msgs <- string(bytes.Trim(b, "\x00"))
		}
	}
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for {
		for msg := range ch {
			conn.Write([]byte(msg))
		}
	}
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

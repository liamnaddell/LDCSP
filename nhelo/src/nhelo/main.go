package main

import "net"
import "fmt"
import "github.com/pelletier/go-toml"
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
		//conn.Write([]byte("accepted"))
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
			var g interface{}
			var name interface{}
			for _, cl := range clients {
				v, err := toml.Load(b)
				checkerr(err)
				g = v.Get("msg")
				name = v.Get("name")
				cl.Ch <- name.(string) + ": " + g.(string)
			}
			fmt.Println(name.(string) + ": " + g.(string))
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
			mgs := string(bytes.Trim(b, "\x00"))
			msgs <- mgs
		}
	}
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for {
		msg := <-ch
		conn.Write([]byte(msg))
	}
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

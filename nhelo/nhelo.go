package nhelo

import "net"
import "bufio"
import "fmt"
import "github.com/pelletier/go-toml"
import "bytes"
import "os"

var msgs = make(chan string)

type Client struct {
	Ch chan<- string
}

var clients []Client

func HostServer() {
	ln, err := net.Listen("tcp", ":9000")
	var rame = bufio.NewScanner(os.Stdin)
	fmt.Println("whats your name?")
	rame.Scan()

	var NAME = rame.Text()
	//serve person is the text interface for the server administrator. basically copied from the lnet interface
	go servePerson(NAME)

	checkerr(err)
	go broad()
	//accepter
	for {
		conn, err := ln.Accept()
		//conn.Write([]byte("accepted"))
		if err != nil {
			panic(err)

		} else {
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

func servePerson(NAME string) {
	bf := bufio.NewScanner(os.Stdin)
	for {
		bf.Scan()
		if bf.Text() != "" {
			var msgg = bf.Text()
			qw := `HOST[` + NAME + `]: ` + msgg
			for _, b := range clients {
				b.Ch <- qw
			}
		}
	}
}

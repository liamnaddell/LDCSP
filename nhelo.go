package ldcsp

import "net"
import "bufio"
import "fmt"
import "strings"
import "os"

var msgs = make(chan string)

type Client struct {
	Ch chan<- string
}

var clients []Client

type umsg struct {
	Msg  string
	Name string
}

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
			var g string
			var name string
			for _, cl := range clients {
				//v, err := toml.Load(b)
				g, name = getInfo(b)
				//g = v.Get("msg")
				//name = v.Get("name")
				cl.Ch <- name + ": " + g
			}
			fmt.Println(name + ": " + g)
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
			//	mgs := string(bytes.Trim(b, "\x00"))
			msgs <- string(b)
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

func getInfo(s string) (string, string) {
	ss := strings.Split(s, "←")
	return ss[0], ss[1]
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "skilstak.sh:8452")
	fmt.Println("past uname")
	if err != nil {
		panic(err)
	}
	fmt.Println("what is your nickname: ")
	nname := bufio.NewScanner(os.Stdin)
	nname.Scan()
	conv(conn, nname.Text())
	defer conn.Close()
}
func conv(conn net.Conn, nname string) {
	RemAdr := conn.RemoteAddr()
	fmt.Println("start talking: ")
	//manage reads
	go func() {
		b := make([]byte, 1024)

		for {
			n, _ := conn.Read(b)
			if n > 0 {
				res := make([]byte, n)

				copy(res, b[:n])
				msgprnt(string(res))
			}
		}
	}()
	go func() {
		NTEXT := bufio.NewScanner(os.Stdin)
		for {
			ntext := ""
			NTEXT.Scan()
			ntext = NTEXT.Text()
			if ntext != "" {
				conn.Write([]byte("[" + nname + "]" + ntext))
			}
			ntext = ""
		}
	}()
	//echo -n "Old line"; echo -e "\e[<N>A new line"
	for conn.RemoteAddr() == RemAdr {
		time.Sleep(1 * time.Second)
	}
	fmt.Println("closed")
}

func msgprnt(msg string) {
	fmt.Println(msg)
}

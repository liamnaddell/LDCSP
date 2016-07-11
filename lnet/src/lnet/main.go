package main

import "net"
import "time"
import "fmt"

func main() {
	conn, err := net.Dial("tcp", "skilstak.sh:9000")
	checkerr(err)
	conn.Write([]byte("shalom"))
	time.Sleep(4 * time.Second)
	b := make([]byte, 1024)
	conn.Read(b)
	fmt.Println(string(b))
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

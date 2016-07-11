package main

import "net"
import "time"
import "fmt"

func main() {
	conn, err := net.Dial("tcp", "skilstak.sh:8382")
	if err != nil {
		panic(err)
	}
	fmt.Println("donepanik")
	fmt.Println("writting")
	_, err2 := conn.Write([]byte("shalmom"))
	if err2 != nil {
		panic(err)
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

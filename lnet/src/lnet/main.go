package main

import "net"
import "os"
import "fmt"
import "time"

//import "io"
import "bufio"

//import "github.com/pelletier/go-toml"

func main() {
	conn, err := net.Dial("tcp", "skilstak.sh:9000")
	checkerr(err)
	//writer
	bf := bufio.NewScanner(os.Stdin)
	go func() {
		for {
			bf.Scan()
			if bf.Text() != "" {
				var qw = []byte("msg = \"" + bf.Text() + "\"")
				fmt.Println(string(qw))
				conn.Write(qw)
			}
		}
	}()
	b := make([]byte, 1024)
	for {
		n, _ := conn.Read(b)
		if n > 0 {
			time.Sleep(1 * time.Millisecond)
			fmt.Println(string(b))
		}
	}
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

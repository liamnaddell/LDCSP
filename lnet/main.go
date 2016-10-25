package main

import "net"
import "os"
import "fmt"
import "bufio"

//import "github.com/pelletier/go-toml"

func main() {
	bf := bufio.NewScanner(os.Stdin)
	fmt.Println("what server do you want to connect to?\n>")
	bf.Scan()
	conn, err := net.Dial("tcp", bf.Text()+":9000")
	checkerr(err)
	//writer
	fmt.Println("what's your name")
	bf.Scan()
	var Name = bf.Text()
	go func() {
		for {
			bf.Scan()
			if bf.Text() != "" {
				var msgg = bf.Text()
				qw := []byte(`msg = "` + msgg + `"
				name = "` + Name + `"`)
				conn.Write(qw)
				msgg = ""
			}
		}
	}()

	b := make([]byte, 1024)
	for {
		n, _ := conn.Read(b)
		if n > 0 {
			var buf = make([]byte, n)
			var empb = make([]byte, 0)
			copy(buf, b)
			fmt.Println(string(buf))
			buf = empb
		}
	}
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

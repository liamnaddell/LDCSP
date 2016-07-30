package main

import "net"
import "os"
import "fmt"
import "bufio"

//import "github.com/pelletier/go-toml"

func main() {
	conn, err := net.Dial("tcp", "skilstak.sh:9000")
	checkerr(err)
	//writer
	bf := bufio.NewScanner(os.Stdin)
	fmt.Println("whats your name")
	bf.Scan()
	var Name = bf.Text()
	go func() {
		for {
			bf.Scan()
			if bf.Text() != "" {
				var msgg = bf.Text()
				qw := []byte(`msg = "` + msgg + `"
				name = "` + Name + `"`)
				println("debug:" + msgg)
				conn.Write(qw)
				msgg = ""
			}
		}
	}()

	//fix printing error mistake here
	b := make([]byte, 1024)
	for {
		n, _ := conn.Read(b)
		if n > 0 {

			fmt.Println(string(b))
		}
	}
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

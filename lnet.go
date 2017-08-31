package ldcsp

import "net"
import "os"
import "fmt"
import "bufio"

type msg struct {
	msg  string
	name string
}

func JoinServer() {
	bf := bufio.NewScanner(os.Stdin)
	fmt.Print("what server do you want to connect to?>")
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
				if err != nil {
					fmt.Println(err)
				}
				qw := []byte(fmt.Sprintf("%s←%s", msgg, Name))
				conn.Write(qw)
				msgg = ""
			}
		}
	}()
	b := make([]byte, 1024)
	var empb = make([]byte, 0)
	for {
		n, _ := conn.Read(b)
		if n != 0 {
			var buf = make([]byte, n)
			copy(buf, b)
			fmt.Println(string(buf))
			buf = empb
		}
	}
}

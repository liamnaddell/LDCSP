package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Print("\nimput a commit message: ")
	bf := bufio.NewScanner(os.Stdin)
	bf.Scan()
	mkfile(bf.Text())
}

func mkfile(msg string) {
	msgn := `
	
	comment=save
	[ ! -z "$*"] && comment="$*"
	
	git pull
	git add -A
	git commit -a -m "` + msg + `"
	git push
	rm lsa`
	f, _ := os.Create("lsa")
	f.Write([]byte(msgn))
	f.Close()
	cmd := exec.Command("chmod", "=rwx", "lsa")
	cmd.Start()
	cmd2 := exec.Command("/bin/sh", "lsa")
	cmd2.Start()
	time.Sleep(4 * time.Second)
}

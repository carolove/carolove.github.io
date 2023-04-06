package main

import (
	"log"
	"net"
	"path/filepath"
	"time"
)

var (
	TransDomainSocket = filepath.Join("../server", "unix.sock")
	Seperator         = byte('#')
)

func main() {
	var (
		unixConn net.Conn
		err      error
	)

	unixConn, err = net.DialTimeout("unix", TransDomainSocket, 1*time.Second)
	if err != nil {
		log.Printf("[warning]: isReconfigure connect failed: %s\n", err)
		return
	}
	defer unixConn.Close()

	buf := make([]byte, 1)
	n, _ := unixConn.Read(buf)
	if n != 1 {
		log.Printf("read unix socket failed!")
		return
	}
	log.Printf("read unix socket success!")
	return
}

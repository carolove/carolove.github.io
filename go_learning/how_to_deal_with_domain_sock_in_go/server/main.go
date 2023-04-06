package main

import (
	"log"
	"net"
	"path/filepath"
	"syscall"
)

var (
	TransDomainSocket = filepath.Join(".", "unix.sock")
	Seperator         = byte('#')
)

func main() {

	// unlink 删除已经存在的domain socket
	_ = syscall.Unlink(TransDomainSocket)

	l, err := net.Listen("unix", TransDomainSocket)
	if err != nil {
		log.Printf("ListenUnixSock faield %s\n", err)
		return
	}
	defer l.Close()

	// unix listener for unix domain socket
	ul := l.(*net.UnixListener)
	for {
		uc, err := ul.AcceptUnix()
		if err != nil {
			if ope, ok := err.(*net.OpError); ok && (ope.Op == "accept") {
				log.Printf("ListenUnixSock unix socket listener closed")
			} else {
				log.Printf("ListenUnixSock Accept error :%v\n", err)
			}
			return
		}
		_, err = uc.Write([]byte{0})
		if err != nil {
			log.Printf("ListenUnixSock Accept error :%v\n", err)
			continue
		}
		_ = uc.Close()
		log.Printf("ListenUnixSock success")
	}
}

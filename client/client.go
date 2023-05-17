package main

import (
	"log"
	"syscall"
)

const port int = 8080
var addr = [4]byte{127, 0, 0, 1}

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatal("couldn't create socket: ", err)
		return
	}
	sa := &syscall.SockaddrInet4{Addr:addr, Port: port}
	if err = syscall.Connect(fd, sa); err != nil {
		log.Fatal("couldn't connect to addr ", addr, err)
	}

	log.Println("Connected successfully")

	syscall.Close(fd)
}
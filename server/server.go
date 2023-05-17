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
		log.Fatal("error creating tcp socket: ", err)
		return
	}

	if err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		syscall.Close(fd)
		log.Fatal("error setting socket options: ", err)
		return
	}

	sa := &syscall.SockaddrInet4{Addr:addr, Port: port}
	if err = syscall.Bind(fd, sa); err != nil {
		log.Fatal("error binding socket to fd: ", err)
		return
	}

    if err = syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		log.Fatal("error listening on fd: ", err)
        return
    }

	log.Println("Server running on ", sa.Addr, ":", sa.Port, " with fd: ", fd)

	for true {
		conn_fd, conn_sa, conn_err := syscall.Accept(fd)
		if conn_err != nil {
			log.Fatal("error handling connection: ", conn_err)
		} else {
			handleConnection(conn_fd, conn_sa) 
		}
	}

}

func handleConnection(conn_fd int, conn_sa syscall.Sockaddr) {
	log.Println("Received connection with fd ", conn_fd, " and sockaddr ", conn_sa)
}
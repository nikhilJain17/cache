package main

import (
	"log"
	"syscall"
	"encoding/binary"
	"bytes"
	// "cache/utils"
)

const maxMessageSize int = 4096
const port int = 8080 
// arrays cannot be const in go
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

// Read all bytes from a socket
func readAll(conn_fd int, msgSize int, msg []byte) {
	// TODO
}

func handleConnection(conn_fd int, conn_sa syscall.Sockaddr) {
	log.Println("Received connection with fd ", conn_fd)

	var msgSizeArr = make([]byte, 4)  
	len, err := syscall.Read(conn_fd, msgSizeArr) 
	if len < 0 {
		log.Fatal("couldn't read msg size from fd")
		return
	}
	if err != nil {
		log.Fatal("couldn't read msg size from client: ", err)
		return
	}
	var msgSize int32
	buf := bytes.NewReader(msgSizeArr)
	err = binary.Read(buf, binary.LittleEndian, &msgSize) 
	if err != nil {
		log.Println("binary.Read failed:", err)
		return
	}
  
	log.Println("Received msg size from client:\n", msgSize)
	

	var input = make([]byte, msgSize)
	len, err = syscall.Read(conn_fd, input)
	if len < 0 {
		log.Fatal("couldn't read from fd")
		return
	}
	if err != nil {
		log.Fatal("couldn't read from client: ", err)
		return
	}
	log.Println("Received message from client:\n", string(input))
	syscall.Write(conn_fd, []byte("world"))
}
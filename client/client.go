package main

import (
	"log"
	"syscall"
	"encoding/binary"
)

const port int = 8080
const maxMessageSize int = 4096
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
	talkToServer(fd)
	syscall.Close(fd)
}

func talkToServer(fd int) {
	msgLen := make([]byte, 4)
	dummyMsg := "teritolay"
	dummyMsgArr := []byte(dummyMsg)
    binary.LittleEndian.PutUint32(msgLen, uint32(len(dummyMsgArr)))
	log.Println("msg -->", dummyMsg, dummyMsgArr, len(dummyMsgArr))

	syscall.Write(fd, msgLen)
	var response = make([]byte, maxMessageSize)
	syscall.Write(fd, dummyMsgArr)
	syscall.Read(fd, response)
	log.Println("Server message:\n", string(response))
}
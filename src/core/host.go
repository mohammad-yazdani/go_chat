package core

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"strconv"
)

const portNum = 9500

type Host struct {
	Ip   string
}

func (Host Host) receive(conn net.Conn, partition string) {
	connReader := bufio.NewReader(conn)
	message, connErr := connReader.ReadString('\n')
	if connErr != nil {
		log.Fatal("Error: ", connErr)
	}
	fmt.Print("Parition " + partition + ": " + message)
}

func (Host Host) Listen(conn net.Conn, partition string)  {
	for {
		Host.receive(conn, partition)
	}
}

func (Host Host) Run(freePartition int) {
	port := fmt.Sprintf("%d", portNum + freePartition)
	address := Host.Ip + ":" + port
	socket, socketErr := net.Listen("tcp", address)

	if socketErr != nil {
		log.Fatal("Error: ", socketErr)
	}
	fmt.Println("Listening on ", address)

	conn, refused := socket.Accept()
	if refused != nil {
		log.Fatal("Error: ", refused)
	}
	fmt.Println("New connection accepted")
	go Host.Listen(conn, strconv.Itoa(freePartition))

	Host.Run(freePartition + 1)
}
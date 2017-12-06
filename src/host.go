package src

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"os"
)

const port = "8080"

type Host struct {
	Ip   string
}

func (Host Host) Listen(conn net.Conn)  {
	connReader := bufio.NewReader(conn)

	message, connErr := connReader.ReadString('\n')
	if connErr != nil {
		log.Fatal("Error: ", connErr)
	}
	fmt.Println("New message: ", message)

	fmt.Print("Reply: ")
	stdReader := bufio.NewReader(os.Stdin)

	reply, sysErr := stdReader.ReadString('\n')
	if sysErr != nil {
		log.Fatal("Error: ", sysErr)
	}

	fmt.Fprint(conn, reply)
}

func (Host Host) Run() {
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

	for {
		Host.Listen(conn)
	}
}
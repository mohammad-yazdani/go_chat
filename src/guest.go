package src

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"os"
)

type Guest struct {
	Ip string
}

func (Guest Guest) Dial(conn net.Conn) {
	fmt.Print("Send: ")

	stdReader := bufio.NewReader(os.Stdin)
	reply, sysErr := stdReader.ReadString('\n')
	if sysErr != nil {
		log.Fatal("Error: ", sysErr)
	}

	fmt.Fprint(conn, reply)

	replyReader := bufio.NewReader(conn)
	reply, connErr := replyReader.ReadString('\n')

	if connErr != nil {
		log.Fatal("Error: ", connErr)
	}

	fmt.Println("Reply:", reply)
}

func (Guest Guest) Run() {
	address := Guest.Ip + ":" + port
	dial, dialErr := net.Dial("tcp", address)

	if dialErr != nil {
		log.Fatal("Error: ", dialErr)
	}
	for {
		Guest.Dial(dial)
	}
}

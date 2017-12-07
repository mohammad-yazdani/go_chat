package core

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

func (Guest Guest) send(conn net.Conn) {
	for {
		stdReader := bufio.NewReader(os.Stdin)

		reply, sysErr := stdReader.ReadString('\n')
		if sysErr != nil {
			log.Fatal("Error: ", sysErr)
		}
		fmt.Fprint(conn, reply)
	}
}

func (Guest Guest) receive(conn net.Conn) {
	replyReader := bufio.NewReader(conn)
	reply, connErr := replyReader.ReadString('\n')

	if connErr != nil {
		log.Fatal("Error: ", connErr)
	}
	fmt.Print("Host: " + reply)
}

func (Guest Guest) Dial(conn net.Conn) {
	for {
		go Guest.send(conn)
		Guest.receive(conn)
	}
}

func (Guest Guest) Run() {
	address := Guest.Ip + ":" + port
	dial, dialErr := net.Dial("tcp", address)

	if dialErr != nil {
		log.Fatal("Error: ", dialErr)
	}


	Guest.Dial(dial)
}

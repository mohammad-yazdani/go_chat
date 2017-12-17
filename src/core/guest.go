package core

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"os"
)

const port  = 9500

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

func (Guest Guest) Dial(conn net.Conn) {
	for {
		Guest.send(conn)
	}
}

func (Guest Guest) Run(partition int) {
	address := Guest.Ip + ":" + fmt.Sprintf("%d", port + partition)
	dial, dialErr := net.Dial("tcp", address)

	if dialErr != nil {
		log.Fatal("Error: ", dialErr)
	}


	Guest.Dial(dial)
}

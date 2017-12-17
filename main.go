package main

import (
	"fmt"
	"time"
	"os"
	"errors"
	endPoint "go_chat/src/core"
	"flag"
	"strconv"
	"log"
)

func getGreeting(hour int) (string, error)  {
	var message string
	if hour < 7 {
		err := errors.New("Too early, we're closed!")
		return message, err
	} else if hour < 12 {
		message = "Good Morning"
	} else if hour < 18 {
		message = "Good Afternoon"
	} else {
		message = "Good evening"
	}
	return message, nil
}

func main() {
	args := os.Args
	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(hourOfDay)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println("Hello, I am Gopher")
	}
	fmt.Println(greeting)

	isHost := flag.Bool("listen", false, "Listen on the given ip address")
	flag.Parse()

	if *isHost {
		ip := os.Args[2]
		host := endPoint.Host{Ip: ip}
		host.Run(0)
	} else {
		ip := os.Args[1]
		partition, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("Incorrect parition format")
		}
		guest := endPoint.Guest{Ip: ip}
		guest.Run(partition)
	}
}

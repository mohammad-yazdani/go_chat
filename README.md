# go_chat
Go chat is a command line chatting app, where two people can share a local network to chat with each other. One of them is a host and the other being the guest.

### COMING SOON!:
For features coming soon checkout other branches.

### Motivation:
This is the final course project for a GOlang course I took online with CodeSchool. I did not purchase the course however I did the final project. 
The code satisfies the same requirements as the course project, however I implemented it with an **object oriented approach**. This version also support **simultaneous** messaging where host and guest can send messages at the same time, which is done using the Golang's **multi-threaded** capablities.

### Build and Run
To run as host:
`go run main.go -listen localhost`

To run as guest:
`go run main.go localhost`

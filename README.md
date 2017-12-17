# Go Chat
Go chat is a command line chatting app, where two people can share a local network to chat with each other. One of them is a host and the other being the guest.

## Many-To-Many (Under developement, branch `many-to-many`)
This branch is scaling the original go chat to support chatting between n many user over an HTTP frontend. It will supports user to user and group chatting.

### Motivation:
This is the final course project for a GOlang course I took online with CodeSchool. I did not purchase the course however I did the final project. 
The code satisfies the same requirements as the course project, however I implemented it with an **object oriented approach**. This version also support **simultaneous** messaging where host and guest can send messages at the same time, which is done using the Golang's **multi-threaded** capablities.

### Updates:
**Dec 17 2017:** Partitioning added. Every guest is now a parition and new guest can be added by using an additional integer argument. After the controller is finished, controller will create paritions as traffic increases. Design documents will be added.

### Build and Run
To run as host:
`go run main.go -listen localhost`

To run a parition:
`go run main.go localhost partitionNumber`
Note: Start with parition 0 and increment. 0 cannot be followed by x > 1 and y cannot be followed by x > y + 1.
Basically: 0, 1, 2, 3, 4 and so on

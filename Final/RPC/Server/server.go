package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Client struct {
	ID   int
	time time.Duration
}

func registerClient(server *rpc.Server, c Client) {
	server.RegisterName("Client", c)
}

func main() {
	//client := new(Client)
	server := rpc.NewServer()
	//registerClient(server, client)
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("Listening in port 8080")
	server.Accept(l)
}

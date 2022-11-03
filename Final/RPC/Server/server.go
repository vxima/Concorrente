package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	ID   int
	time time.Duration
}
type Client struct {
	ID   int
	time time.Duration
}

func (c *Client) GiveClientID(args *Args, reply *int) error {
	*reply = args.ID
	return nil
}
func registerClient(server *rpc.Server, c Client) {
	rpc.Register(c)
}

func main() {
	client := new(Client)
	rpc.Register(client)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("Listening in port 8080")
	http.Serve(l, nil)
}

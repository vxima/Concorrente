package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Args struct {
	ID   int
	time time.Duration
}
type Client struct {
	ID   int
	time time.Duration
}

func NewClient(id int) Client {
	//incoming := client_incoming()
	tmp := client_time()

	//time.Sleep(incoming) // waits to make a new client
	c := Client{
		ID:   id,
		time: tmp,
	}
	return c

}
func client_time() time.Duration {
	rng := rand.Intn(10)
	return time.Duration(rng) * time.Second
}
func main() {
	fmt.Println("Clients start coming")
	//qntdClients := 1
	/*for i := 1; i <= qntdClients; i++ {
		client := NewClient(i)
	}*/
	conn, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connectiong:", err)
	}
	fmt.Println("Connection Successfully")
	//Synchronus call
	client := NewClient(1)
	args := &Args{client.ID, client.time}
	var reply int
	// call
	err = conn.Call("Client.GiveClientID", args, &reply)
	if err != nil {
		log.Fatal("ID error:", err)
	}
	fmt.Printf("Client info: {%d ,%d} , received= %d\n", args.ID, args.time, reply)

}

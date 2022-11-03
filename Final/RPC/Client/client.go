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
	//Create client and args

	var reply int
	// call method from server

	// Print reply from server
	// put in the queue
	for i := 0; i < 20; i++ {
		client := NewClient(i)
		args := &Args{client.ID, client.time}
		err = conn.Call("Client.NewClientRPC", args, &reply)
		if err != nil {
			log.Fatal("ID error:", err)
		}
	}

}

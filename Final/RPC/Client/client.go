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

func client_request(i int, reply int, conn *rpc.Client) {
	client := NewClient(i)
	args := &Args{client.ID, client.time}
	err := conn.Call("Client.NewClientRPC", args, &reply)
	if err != nil {
		log.Fatal("ID error:", err)
	}
	wg.Done()
}
func main() {
	// Connect with the server
	conn, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connectiong:", err)
	}
	fmt.Println("Connection Successfully")

	var reply int
	ClientsInDay := 20 //determines the quantity of clients that we go
	// call method from server
	fmt.Println("Clients start coming!")
	for i := 1; i <= ClientsInDay; i++ {
		wg.Add(1)
		go client_request(i, reply, conn)
	}
	wg.Wait()

}

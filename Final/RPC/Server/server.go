package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

// global variables
var wg sync.WaitGroup
var waiting Queue
var t_size int

// Arguments passed from client.go
type Args struct {
	ID   int
	time time.Duration
}

// Client struct and method
type Client struct {
	ID   int
	time time.Duration
}

func (c *Client) NewClientRPC(args *Args, reply *int) error {
	*reply = args.ID                            //reply is the response to the client.go
	cli := Client{ID: args.ID, time: args.time} //transforms the args passed in a client
	waiting.push(cli)                           // puts the client in the queue
	return nil
}

// queue algorithm
type Queue struct {
	mu    sync.Mutex
	queue []Client
	size  int
}

func (q *Queue) push(c Client) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, c)
	q.size += 1
}
func (q *Queue) pop() Client {
	q.mu.Lock()
	defer q.mu.Unlock()
	x := q.queue[0]
	q.queue = q.queue[1:]
	q.size -= 1
	return x
}

func client_consuming(c Client) { //sleep the goroutine to simulate the client consuming
	defer wg.Done()
	time.Sleep(c.time)
	fmt.Printf("Client %d has finished\n", c.ID)
	t_size -= 1 // client leaves the table

}

// Management of the table
func Managing() {
	//while(1)
	for {
		if t_size == 5 {
			fmt.Println("Full Table!")
			wg.Wait() //wait all clients to finish
			fmt.Println("Group of friends leaving")
		} else if waiting.size > 0 {
			wg.Add(1)
			c := waiting.pop() //client in the waiting queue go to the table
			fmt.Printf("Entering client %d\n", c.ID)
			t_size += 1            //increment the table size
			go client_consuming(c) //client starts consuming
		}
	}

}
func main() {
	// Create instance
	client := new(Client)
	// Register
	rpc.Register(client)
	// Make connection
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("Listening in port 8080")
	// Two goroutines running:
	// 1) http.Serve() reads the request from the clients

	// 2) Managing() is an infinity loop thats gets the clients in the waiting queue
	// and manages the table with the time consuming of each client
	go func() {
		http.Serve(l, nil)
	}()
	go Managing()
	select {} //this select is to run simultaneously both http.serve() and Managing()

}

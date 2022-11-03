package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

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
	qntdClients := 5
	for i := 1; i <= qntdClients; i++ {
		client := NewClient(i)
	}
}

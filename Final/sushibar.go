package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var t_size int

// queue algorithm
type Queue struct {
	mu    sync.Mutex
	queue []Client
	size  int
}

type Table struct {
	mu       sync.Mutex
	queue    []Client
	capacity int
	size     int
}

///estrutura
type Client struct {
	ID   int
	time time.Duration
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

// Auxiliar functions
func client_time() time.Duration {
	rng := rand.Intn(10)
	return time.Duration(rng) * time.Second
}

func client_incoming() time.Duration {
	rng := 5 + rand.Intn(10)
	return time.Duration(rng) * time.Second
}

func client_consuming(c Client) {
	defer wg.Done()
	time.Sleep(c.time)
	fmt.Printf("Client %d has finished\n", c.ID)
	t_size -= 1

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
func main() {
	//table_size := 5
	//occupied := 0
	table := Table{queue: make([]Client, 5), capacity: 5, size: 0}
	waiting := Queue{queue: make([]Client, 0), size: 0}
	//make clientes
	for i := 1; i <= 20; i++ {
		waiting.push(NewClient(i))
	}
	//5 goroutines para gerenciar as mesas
	for waiting.size != 0 {
		if t_size == table.capacity {
			fmt.Println("Full Table")
			wg.Wait()
			fmt.Println("Group of friends leaving")
		} else {
			wg.Add(1)
			c := waiting.pop()
			fmt.Printf("Entering client %d\n", c.ID)
			t_size += 1
			go client_consuming(c)
		}
	}
	wg.Wait()
}

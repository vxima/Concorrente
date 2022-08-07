package main

import (
	"fmt"
	"sync"
)

type MutexIntBuffer struct {
	mut      sync.Mutex
	capacity int
	buff     []int
}

func NewMutexIntBuffer(capacity int) MutexIntBuffer {
	return MutexIntBuffer{
		mut:      sync.Mutex{},
		capacity: capacity,
		buff:     []int{},
	}
}

func (s *MutexIntBuffer) produtor(n int) {
	s.mut.Lock()
	s.buff = append(s.buff, n)
	s.mut.Unlock()
}

func (s *MutexIntBuffer) consumidor() int {
	s.mut.Lock()
	valor := s.buff[0] //pega o primeiro valor do buffer
	s.buff = s.buff[1:]
	s.mut.Unlock()
	return valor
}

func main() {
	cap := 10   //tamanho da capacidade do buffer
	n_cons := 4 //numero de consumidores
	mutBuff := NewMutexIntBuffer(cap)
	go mutBuff.produtor(1)
	for i := 0; i < n_cons; i++ {
		value := mutBuff.consumidor()
		fmt.Println("Consumidor #%d recebeu %d", i, value)
	}
}

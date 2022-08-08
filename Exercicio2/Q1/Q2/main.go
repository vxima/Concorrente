package main

import (
	"fmt"
	"runtime"
	"sync"
)

const cap int = 10 //tamanho da capacidade do buffer
const n_cons = 4   //numero de consumidores

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

func (s *MutexIntBuffer) produzir(n int) {
	s.mut.Lock()
	for len(s.buff) == s.capacity {
		s.mut.Unlock()
		runtime.Gosched() //chama outro
		s.mut.Lock()

	}

	s.buff = append(s.buff, n)
	s.mut.Unlock()
}

func (s *MutexIntBuffer) consumir() int {
	s.mut.Lock()
	for len(s.buff) == 0 {
		s.mut.Unlock()
		runtime.Gosched()
		s.mut.Lock()
	}
	valor := s.buff[0] //pega o primeiro valor do buffer
	s.buff = s.buff[1:] //pegar tail
	s.mut.Unlock()

	return valor
}

func produtor(id string, value int, wg *sync.WaitGroup, mib *MutexIntBuffer) {
	defer wg.Done()
	for i := 0; i < cap; i++ {
		mib.produzir(i)
		fmt.Println(id, "produziu", i)
	}

}

func consumidor(id int, wg *sync.WaitGroup, mib *MutexIntBuffer) {
	defer wg.Done()
	e := mib.consumir()
	fmt.Println(id, "consumiu", e)
}
func main() {
	wg := sync.WaitGroup{}

	for 
	wg.Wait()

}

package main

import (
	"fmt"
	"sync"
)
type post struct{
	views int
	mu sync.Mutex
}

func(p *post) inc (w *sync.WaitGroup){
	defer func ()  {
		p.mu.Unlock()
		w.Done()
	}()
	p.mu.Lock()
	p.views++
	
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)
}
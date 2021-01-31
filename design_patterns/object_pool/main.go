package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

type connection struct {
	id string
}

var _ poolObject = (*connection)(nil)

func (c *connection) getID() string {
	return c.id
}

func main() {
	var wg sync.WaitGroup

	connections := make([]poolObject, 0)

	for i := 0; i < 3; i++ {
		c := &connection{
			id: strconv.Itoa(i),
		}
		connections = append(connections, c)
	}

	myPool, err := createPool(connections)
	if err != nil {
		log.Fatalf("could not create object pool: %v", err)
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			conn, loanErr := myPool.loan()
			if loanErr != nil {
				log.Fatalf("loan object error :%v", loanErr)
			}
			time.Sleep(1 * time.Second)
			myPool.done(conn)
		}(&wg)
	}

	wg.Wait()
}

package main

import (
	"math/rand"
	"sync"
	"time"
)

type stationManager struct {
	isStationAvailable bool
	lock               *sync.Mutex
	trainQueue         []train
}

var _ mediator = (*stationManager)(nil)

func newStationManager() *stationManager {
	return &stationManager{
		isStationAvailable: true,
		lock:               &sync.Mutex{},
	}
}

func (s *stationManager) canArrive(t train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isStationAvailable {
		s.isStationAvailable = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyAvailable() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isStationAvailable {
		s.isStationAvailable = true
	}
	if len(s.trainQueue) > 0 {
		nextTrain := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		nextTrain.permitArrival()
	}
}

func randBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func main() {
	var wg sync.WaitGroup
	sm := newStationManager()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		randVal := randBool()
		var aTrain train
		if randVal {
			aTrain = &passengerTrain{
				mediator: sm,
			}
		} else {
			aTrain = &goodsTrain{
				mediator: sm,
			}
		}

		go func(wg *sync.WaitGroup, t train) {
			defer wg.Done()
			t.requestArrival()
			landTime := rand.Intn(5-1) + 1
			time.Sleep(time.Duration(landTime) * time.Second)
			t.departure()
		}(&wg, aTrain)
	}

	wg.Wait()
}

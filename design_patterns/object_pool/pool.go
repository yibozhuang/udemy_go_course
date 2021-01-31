package main

import (
	"errors"
	"fmt"
	"sync"
)

type poolObject interface {
	getID() string
}

type pool struct {
	idle     []poolObject
	active   []poolObject
	capacity int
	lock     *sync.Mutex
}

func createPool(objects []poolObject) (*pool, error) {
	if len(objects) == 0 {
		return nil, errors.New("cannot create a pool of 0 length")
	}
	active := make([]poolObject, 0)
	pool := &pool{
		idle:     objects,
		active:   active,
		capacity: len(objects),
		lock:     new(sync.Mutex),
	}

	return pool, nil
}

func (p *pool) loan() (poolObject, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if len(p.idle) == 0 {
		return nil, errors.New("no idle objects are available")
	}

	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("allocated object with ID: %s\n", obj.getID())
	return obj, nil
}

func (p *pool) remove(obj poolObject) error {
	activeLength := len(p.active)
	for i, active := range p.active {
		if obj.getID() == active.getID() {
			p.active[activeLength-1], p.active[i] = p.active[i], p.active[activeLength-1]
			p.active = p.active[:activeLength-1]
			return nil
		}
	}
	return errors.New("object not found in pool")
}

func (p *pool) done(obj poolObject) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	err := p.remove(obj)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, obj)
	fmt.Printf("returned object with ID: %s\n", obj.getID())
	return nil
}

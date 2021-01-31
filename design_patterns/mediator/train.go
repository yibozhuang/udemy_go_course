package main

import (
	"fmt"
)

type mediator interface {
	canArrive(train) bool
	notifyAvailable()
}

type train interface {
	requestArrival()
	departure()
	permitArrival()
}

type passengerTrain struct {
	mediator mediator
}

var _ train = (*passengerTrain)(nil)

func (pt *passengerTrain) requestArrival() {
	if pt.mediator.canArrive(pt) {
		fmt.Println("Passenger train landing...")
	} else {
		fmt.Println("Passenger train waiting...")
	}
}

func (pt *passengerTrain) departure() {
	fmt.Println("Passenger train leaving...")
	pt.mediator.notifyAvailable()
}

func (pt *passengerTrain) permitArrival() {
	fmt.Println("Passenger train permitted to arrive...")
}

type goodsTrain struct {
	mediator mediator
}

var _ train = (*goodsTrain)(nil)

func (gt *goodsTrain) requestArrival() {
	if gt.mediator.canArrive(gt) {
		fmt.Println("Goods train landing...")
	} else {
		fmt.Println("Goods train waiting...")
	}
}

func (gt *goodsTrain) departure() {
	fmt.Println("Goods train leaving...")
	gt.mediator.notifyAvailable()
}

func (gt *goodsTrain) permitArrival() {
	fmt.Println("Goods train permitted to arrive...")
}

package main

import "fmt"

type orchestrator struct {
	builder Builder
}

func newOrchestrator(b Builder) *orchestrator {
	return &orchestrator{
		builder: b,
	}
}

func (o *orchestrator) setBuilder(b Builder) {
	o.builder = b
}

func (o *orchestrator) buildResidence() residence {
	o.builder.setWindow()
	o.builder.setDoor()
	o.builder.setFloorInfo()
	return o.builder.build()
}

func main() {
	house := getBuilder("house")
	condo := getBuilder("condo")

	director := newOrchestrator(house)
	r := director.buildResidence()
	fmt.Println(r.String())

	director.setBuilder(condo)
	r = director.buildResidence()
	fmt.Println(r.String())
}

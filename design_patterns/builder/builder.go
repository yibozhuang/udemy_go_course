package main

import "fmt"

type residence struct {
	window    string
	door      string
	floorInfo string
}

var _ fmt.Stringer = (*residence)(nil)

func (r *residence) String() string {
	return fmt.Sprintf("Residence info: \nWindow type: %s\nDoor type: %s\nFloor info: %s\n",
		r.window, r.door, r.floorInfo)
}

type Builder interface {
	setWindow()
	setDoor()
	setFloorInfo()
	build() residence
}

type condoBuilder struct {
	residence
}

var _ Builder = (*condoBuilder)(nil)

func newCondoBuilder() *condoBuilder {
	return &condoBuilder{}
}

func (c *condoBuilder) setWindow() {
	c.residence.window = "Glass Window"
}

func (c *condoBuilder) setDoor() {
	c.residence.door = "Steel Door"
}

func (c *condoBuilder) setFloorInfo() {
	c.residence.floorInfo = "14 Floors"
}

func (c *condoBuilder) build() residence {
	return c.residence
}

type houseBuilder struct {
	residence
}

var _ Builder = (*houseBuilder)(nil)

func newHouseBuilder() *houseBuilder {
	return &houseBuilder{}
}

func (c *houseBuilder) setWindow() {
	c.residence.window = "Wooden Window"
}

func (c *houseBuilder) setDoor() {
	c.residence.door = "Wooden Door"
}

func (c *houseBuilder) setFloorInfo() {
	c.residence.floorInfo = "1 Floor"
}

func (c *houseBuilder) build() residence {
	return c.residence
}

func getBuilder(builderType string) Builder {
	if builderType == "condo" {
		return newCondoBuilder()
	}
	if builderType == "house" {
		return newHouseBuilder()
	}
	return nil
}

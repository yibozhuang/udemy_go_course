package main

import (
	"errors"
)

type sportFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

type adidasShoe struct {
	shoe
}

type nikeShoe struct {
	shoe
}

type adidasShort struct {
	short
}

type nikeShort struct {
	short
}

type adidas struct{}

var _ sportFactory = (*adidas)(nil)

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) makeShort() iShort {
	return &adidasShort{
		short: short{
			logo: "adidas",
			size: 14,
		},
	}
}

type nike struct{}

var _ sportFactory = (*nike)(nil)

func (a *nike) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (a *nike) makeShort() iShort {
	return &adidasShort{
		short: short{
			logo: "nike",
			size: 14,
		},
	}
}

func getSportFactory(brand string) (sportFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, errors.New("unsupported brand")
}

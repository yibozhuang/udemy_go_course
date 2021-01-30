package main

import "fmt"

func main() {
	adidasFactory, _ := getSportFactory("adidas")
	nikeFactory, _ := getSportFactory("nike")

	nShoe := nikeFactory.makeShoe()
	aShoe := adidasFactory.makeShoe()

	printShoeDetails(nShoe)
	printShoeDetails(aShoe)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

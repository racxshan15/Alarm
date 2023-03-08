package main

import (
	"fmt"
)

func main() {
	var getal int
	fmt.Println("Voer een getal in:")
	fmt.Scan(&getal)
	forloop(getal)
}

func forloop(getal int) {
	for i := 0; i < getal; i++ {
		fmt.Println("alarm!", i)
	}
}

package main

import (
	"fmt"
)

func main() {
	var getal int
	fmt.Println("Voer een getal in:")
	fmt.Scan(&getal)
	if getal < 0 {
		fmt.Println("Geen NEGATIEVE getallen!! Voer een nieuwe getal in:")
		fmt.Scan(&getal)
	}
	for i := 0; i < getal; i++ {
		fmt.Println("alarm!", i)
	}
}

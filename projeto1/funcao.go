package main

import "fmt"

func main() {
	lista := []float64{10, 7, 8, 10, 10}
	total := 0.0
	for _, valor := range lista {
		total += valor
	}

	fmt.Println(total)
}

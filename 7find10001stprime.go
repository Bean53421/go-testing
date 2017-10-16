package main

import "fmt"
import "time"

func main() {
	// Time taken
	start := time.Now()

	length := 104999

	// Primes.
	primes := make([]int, length)
	for ind := range primes {
		primes[ind] = ind + 2
	}

	count := 0

	// Eulcidean seive.
	for ind, val := range primes {
		if val != 1 {
			count++
			if count == 10001 {
				fmt.Printf("Prime number %d: %d\n", count, val)
				break
			} else {
				for j := ind + val; j < length; j += val {
					if primes[j] != 1 {
						primes[j] = 1
					}
				}
			}
		}
	}
	fmt.Println("Time taken:", time.Now().Sub(start))
}

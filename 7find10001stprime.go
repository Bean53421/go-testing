package main

import "fmt"
import "time"

func main() {
	// Time taken
	start := time.Now()

	// Primes.
	primes := make([]int, 105000)
	for ind := range primes {
		primes[ind] = ind + 2
	}

	count := 0
	// Eulcidean seive.
	for i := 0; i < len(primes); i++ {
		if primes[i] != 1 {
			count++
			if count == 10001 {
				fmt.Printf("Prime number %d: %d\n", count, primes[i])
				break
			} else {
				for j := i + primes[i]; j < len(primes); j += primes[i] {
					if primes[j] != 1 {
						primes[j] = 1
					}
				}
			}
		}
	}
	fmt.Println("Time taken:", time.Now().Sub(start))
}
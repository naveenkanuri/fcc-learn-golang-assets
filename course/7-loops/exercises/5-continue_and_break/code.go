package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		if i == 2 {
			fmt.Printf("%v is prime\n", i)
			continue
		} else if i%2 == 0 {
			continue
		}
		isPrime := true
		for j := 3; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%v is prime\n", i)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(100)
	//test(20)
	//test(30)
}

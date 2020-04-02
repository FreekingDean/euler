package euler

import (
	"log"
)

func placeholder() {
	log.Println("never")
}

func P1(n int) int {
	return (mul(3, n-1) + mul(5, n-1)) - mul(15, n-1)
}

func mul(n, b int) int {
	f := b / n
	//log.Println(f)
	return sum(f) * n
}

func sum(n int) int {
	return ((n * n) + n) / 2
}

func P2(n int) int {
	sum := 0
	p := 1
	i := 1
	for i < n {
		//log.Println(i)
		if i%2 == 0 {
			sum += i
		}
		t := i
		i = p + i
		p = t
	}
	return sum
}

func prime(n int, primes []int) bool {
	for _, prime := range primes {
		if n%prime == 0 {
			return false
		}
	}
	return true
}

func P3(n int) int {
	factors := []int{}
	primes := []int{}
	np := 2
	for n != 1 {
		if np > n {
			log.Fatalf("Should not get here n:%d np:%d factors:%v", n, np, factors)
		}
		if n%np == 0 {
			factors = append(factors, np)
			n = n / np
		}
		if np == 2 {
			np = np + 1
		} else {
			np = np + 2
		}
		for !prime(np, primes) {
			np = np + 2
		}
		primes = append(primes, np)
	}
	return factors[len(factors)-1]
}

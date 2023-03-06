package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd(a int, n int) int {
	var large int
	var small int
	if a > n {
		large = a
		small = n
	} else {
		large = n
		small = a
	}
	last := large % small

	for last != 0 {
		large = small
		small = last

		// "next" iteration
		last = large % small
	}

	return small
}

func ptest(a int, n int) int { // test primality of value n using int a
	counter := 1 // test if counter starts at 0 or 1 using new pmod func
	val := gcd(a, n)

	for val == 1 {
		pmod(&a, counter, n)
		//fmt.Println("a:", a)
		if a == 1 {
			val = -1
			break
		}
		val = gcd(a-1, n)
		counter++
	}

	if val != n {
		return val
	} else {
		return -1
	}
}

func pmod(a *int, i int, n int) { //a^i mod n
	// fmt.Println("pmod")
	// fmt.Println("a:", *a)
	oldA := *a
	for I := 1; I < i; I++ {
		*a *= oldA
		*a %= n
		if *a < 0 {
			*a += n
		}
	}
	// fmt.Println("a:", *a)

} // returns a^i mod n

func millerRabin(a int, n int) bool { // return true if probably prime, false for composite, same as p-1, testing primality of n using a
	N := n
	n--
	k := n / 2
	n /= 2
	counter := 1
	for n%2 == 0 {
		k = n / 2
		n /= 2
		counter++
	} // have k, use N for original n
	pmod(&a, k, N)

	if a == N-1 || a&N == 1 {
		return true
	}
	pmod(&a, 2, N)
	for i := 1; i < counter; i++ { //loop from b1 to b_i-1
		if a%N == 1 {
			return false
		}
		if a%N == N-1 {
			return true
		}
		pmod(&a, 2, N)
	}
	if a%n == N-1 {
		return true
	} else {
		return false
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: test(ex gcd, p-1) n a")
		return
	}
	a, _ := strconv.Atoi(os.Args[3])
	n, _ := strconv.Atoi(os.Args[2])

	switch os.Args[1] {
	case "gcd": // breaks unneeded in go
		v := gcd(a, n)
		fmt.Println("gcd(", a, ",", n, ") =", v)
	case "p-1":
		v := ptest(a, n)
		if v == -1 {
			fmt.Println("Test inconclusive")
		} else {
			fmt.Println(n, "is composite, divisible by", v)
		}
	case "miller-rabin":
		if millerRabin(a, n) {
			fmt.Println(n, "is probably prime")
		} else {
			fmt.Println(n, "is composite")
		}
	default:
		fmt.Println("Unexpected test passed in, error")
	}
}

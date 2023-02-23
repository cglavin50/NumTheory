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

func ptest(a int, n int) int {
	counter := 1
	val := gcd(a, n)

	for val == 1 {
		oldA := a
		for i := 0; i < counter; i++ {
			a *= oldA
			a %= n
			if a < 0 {
				a += n
			}
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

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: test(ex gcd, p-1) a n")
		return
	}
	a, _ := strconv.Atoi(os.Args[3])
	n, _ := strconv.Atoi(os.Args[2])

	switch os.Args[1] {
	case "gcd": // breaks unneeded in go
		v := gcd(a, n)
		fmt.Println("gcd(", a, ",", n, ") = ", v)
	case "p-1":
		v := ptest(a, n)
		if v == -1 {
			fmt.Println("Test inconclusive")
		} else {
			fmt.Println(n, " is composite, divisible by ", v)
		}
	default:
		fmt.Println("Unexpected test passed in, error")
	}
}

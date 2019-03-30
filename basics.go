package main

import (
	"fmt"
	"math"
	"time"
)

func helloworld() {
	fmt.Printf("Hello World")
}

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short" //var f string = "short"
	fmt.Println(f)
}

const s string = "constant"

func constants() {
	fmt.Println(s)

	const n = 500000

	println(math.Sin(n))
}

func iterators() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 4; j <= 9; j++ {
		fmt.Println(j)
	}
}

func ifelse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
}

func switchcase() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}

func main() {
	helloworld()
	variables()
	constants()
	iterators()
	ifelse()
	switchcase()
}

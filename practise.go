package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

var arraysV [2]string

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x float64, n float64, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func switchE() {
	fmt.Printf("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}

func switchWeek() {
	fmt.Println("When's saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")

	}
}

func switchHour() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferEx() {
	defer fmt.Println("world")

	fmt.Println("Hello")
}

func stackDeferEx() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func pointer1() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func VertexEx() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
}

func VertexEx1() {
	vv := Vertex{1, 2}
	pp := &vv
	pp.X = 1e9
	fmt.Println(vv)
}

func VertexEx2() {
	fmt.Println(v1, p, v2, v3)
}

func arraysEx() {
	arraysV[0] = "Hello"
	arraysV[1] = "world"

	fmt.Println(arraysV[0], arraysV[1])

	fmt.Println(arraysV)

	primes := [6]int{22, 33, 44, 55, 6, 66}
	fmt.Println(primes)
}

func sliceEx() {
	primes := [6]int{1, 2, 3, 4, 5, 6}
	var s []int = primes[1:4]
	fmt.Println(s)
}

func slicesNameEx() {
	names := [4]string{
		"John",
		"Paul",
		"Georage",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:2]

	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiteralsEx() {
	q := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(q)

	r := []bool{true, false, true, false}
	fmt.Println(r)

	s := []struct {
		x int
		y bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, true},
	}
	fmt.Println(s)
}

func main() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 20),
		pow(3, 3, 20),
	)

	switchE()
	switchWeek()
	switchHour()
	deferEx()
	stackDeferEx()
	pointer1()
	VertexEx()
	VertexEx1()
	VertexEx2()
	arraysEx()
	sliceEx()
	slicesNameEx()
	sliceLiteralsEx()
}

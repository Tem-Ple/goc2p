package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type Sumer interface {
	MySum() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) MySum() float64 {
	if f == 0 {
		return float64(f)
	}
	return float64(f-f)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())

	var s Sumer
	S := MyFloat(344)
	s = S
	fmt.Println(s.MySum())
}

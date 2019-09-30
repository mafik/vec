package vec

import (
	"fmt"
	"testing"
)

func ExampleNew() {
	a := Vec{X: 1, Y: 2}
	b := New(1, 2)
	fmt.Println(a.Equal(b))
	// Output: true
}

func TestZero(t *testing.T) {
	if !Zero.Equal(Vec{0, 0}) {
		t.Fail()
	}
}

func ExampleVec_Add() {
	a, b := Vec{1, 2}, Vec{3, 5}
	fmt.Println(a.Add(b))
	// Output: {4 7}
}

func ExampleVec_Sub() {
	a, b := Vec{1, 2}, Vec{3, 5}
	fmt.Println(a.Sub(b))
	// Output: {-2 -3}
}

func ExampleVec_Div() {
	a, b := Vec{3, 5}, Vec{1, 2}
	fmt.Println(a.Div(b))
	// Output: {3 2.5}
}

func ExampleVec_Mul() {
	a, b := Vec{1, 2}, Vec{3, 5}
	fmt.Println(a.Mul(b))
	// Output: {3 10}
}

func ExampleVec_Len() {
	a := Vec{3, 4}
	fmt.Println(a.Len())
	// Output: 5
}

func ExampleVec_Neg() {
	a := Vec{-0.5, 100}
	fmt.Println(a.Neg())
	// Output: {0.5 -100}
}

func ExampleVec_Scale() {
	a := Vec{-1, 100}
	fmt.Println(a.Scale(-0.5))
	// Output: {0.5 -50}
}

func ExampleVec_ScaleTo() {
	a := Vec{30, 40}
	fmt.Println(a.ScaleTo(5))
	// Output: {3 4}
}

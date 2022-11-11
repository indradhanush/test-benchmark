package main

import "testing"

func BechmarkCalculate(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		Calculate(1, 2)
	}
}

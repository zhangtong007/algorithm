package main

import "testing"

func BenchmarkIsOdd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsOdd1(111111)
	}
}

func BenchmarkIsOdd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsOdd2(111111)
	}
}

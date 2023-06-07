package main

import "testing"

func BenchmarckHoge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		println("愛はあるんか")
	}
}

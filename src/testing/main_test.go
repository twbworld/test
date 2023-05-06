package main

import (
	"fmt"
	"testing"
)

//测试两者性能,命令: go test -v -run="none" -bench=.

func BenchmarkConcurrentAtomicAdd(b *testing.B) {
	b.ResetTimer()
	fmt.Println("===============")
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		ConcurrentAtomicAdd()
	}
}

func BenchmarkConcurrentMutexAdd(b *testing.B) {
	b.ResetTimer()
	fmt.Println("===============")
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		ConcurrentMutexAdd()
	}
}

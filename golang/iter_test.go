// https://github.com/drgrib/iter
/*
	Run using
	go test -bench=. -run=.
*/

package iter

import (
	bIter "github.com/bradfitz/iter"
	"testing"
)

//////////////////////////////////////////////
/// many loops
//////////////////////////////////////////////

const loopsMany = 1e6

func BenchmarkForMany(b *testing.B) {
	b.ReportAllocs()
	j := 0
	for i := 0; i < b.N; i++ {
		for j = 0; j < loopsMany; j++ {
			j = j
		}
	}
	_ = j
}

func BenchmarkDrgribIterMany(b *testing.B) {
	b.ReportAllocs()
	j := 0
	for i := 0; i < b.N; i++ {
		for j = range N(loopsMany) {
			j = j
		}
	}
	_ = j
}

func BenchmarkBradfitzIterMany(b *testing.B) {
	b.ReportAllocs()
	j := 0
	for i := 0; i < b.N; i++ {
		for j = range bIter.N(loopsMany) {
			j = j
		}
	}
	_ = j
}

//////////////////////////////////////////////
/// loops10
//////////////////////////////////////////////

const loops10 = 10

func BenchmarkFor10(b *testing.B) {
	b.ReportAllocs()
	j := 0
	for i := 0; i < b.N; i++ {
		for j = 0; j < loops10; j++ {
			j = j
		}
	}
	_ = j
}

func BenchmarkDrgribIter10(b *testing.B) {
	b.ReportAllocs()
	j := 0
	for i := 0; i < b.N; i++ {
		for j = range N(loops10) {
			j = j
		}
	}
	_ = j
}

func BenchmarkBradfitzIter10(b *testing.B) {
	b.ReportAllocs()
	j := 0
	for i := 0; i < b.N; i++ {
		for j = range bIter.N(loops10) {
			j = j
		}
	}
	_ = j
}

package client

import (
	"testing"
)

func BenchmarkCovidCasesGrpc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TestGetAllCovidCasesGrpc()
	}
}

func BenchmarkCovidCasesRest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TestRest("covid.json")
	}
}

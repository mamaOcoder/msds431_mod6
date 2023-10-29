package main

import (
	"fmt"
	"testing"
)

func BenchmarkRunSerial(b *testing.B) {

	houses, err := parseCSV("./boston.csv")
	if err != nil {
		fmt.Printf("Error parsing file: %s", err)
	}

	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			runSerial(houses)
		}
	}
}

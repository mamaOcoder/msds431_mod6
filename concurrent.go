package main

import (
	"sync"
)

func runConcurrent(houses []record) {
	// Cannot get multiple linear regression to work, so selecting single predictor variables to test
	// Create nox, rooms, crim, mv slices
	var nox, rooms, crim, mv []float64
	for _, h := range houses {
		nox = append(nox, h.nox)
		rooms = append(rooms, h.rooms)
		crim = append(crim, h.crim)
		mv = append(mv, h.mv) //response variable
	}

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		_ = runLinearRegression("Nox", nox, mv)
		wg.Done()
	}()

	go func() {
		_ = runLinearRegression("Rooms", rooms, mv)
		wg.Done()
	}()

	go func() {
		_ = runLinearRegression("Crim", crim, mv)
		wg.Done()
	}()

	wg.Wait()

}

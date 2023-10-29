package main

func runSerial(houses []record) {
	// Cannot get multiple linear regression to work, so selecting single predictor variables to test
	// Create nox, rooms, crim, mv slices
	var nox, rooms, crim, mv []float64
	for _, h := range houses {
		nox = append(nox, h.nox)
		rooms = append(rooms, h.rooms)
		crim = append(crim, h.crim)
		mv = append(mv, h.mv) //response variable
	}

	// Run linear regression model for each predictor serially
	_ = runLinearRegression("Nox", nox, mv)

	_ = runLinearRegression("Rooms", rooms, mv)

	_ = runLinearRegression("Crim", crim, mv)

}

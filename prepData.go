package main

// Function to split training and testing data given a slice of records
func splitTrainTestRecords(data []record) ([]record, []record, error) {
	// Split data for training and testing
	// Define the ratio of the dataset to use for training (e.g., 80% for training, 20% for testing)
	trainingRatio := 0.8
	// Calculate the number of samples for training and testing
	totalSamples := len(data)
	trainingSamples := int(float64(totalSamples) * trainingRatio)

	trainingData := data[:trainingSamples]
	testingData := data[trainingSamples:]

	return trainingData, testingData, nil
}

// Function to split response and predictors
func splitPredResponse(data []record) ([][]float64, []float64, error) {
	// Separate response and predictors
	var X [][]float64
	var Y []float64

	for _, h := range data {
		X = append(X, []float64{h.crim, h.zn, h.indus, h.chas, h.nox, h.rooms, h.age, h.dis, h.rad, h.tax, h.ptratio, h.lstat})
		Y = append(Y, h.mv)
	}

	return X, Y, nil
}

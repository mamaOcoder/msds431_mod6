package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gonum/stat"
)

type record struct {
	neighborhood string
	crim         float64
	zn           float64
	indus        float64
	chas         float64
	nox          float64
	rooms        float64
	age          float64
	dis          float64
	rad          float64
	tax          float64
	ptratio      float64
	lstat        float64
	mv           float64
}

// Read and parse CSV data
func parseCSV(fname string) ([]record, error) {
	// Open the CSV file
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Cannot open file: %v", err)
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	cr := csv.NewReader(file)

	// Define slice to store house records
	var houses []record

	// Read in all CSV data
	allData, err := cr.ReadAll()
	if err != nil {
		fmt.Printf("Cannot read data from file: %v", err)
		return nil, err
	}

	// Read and parse each line of the CSV file
	for i, row := range allData {
		// Skip the header line
		if i == 0 {
			continue
		}

		// Parse the values from the CSV line
		var house record

		house.neighborhood = row[0]
		house.crim, err = strconv.ParseFloat(row[1], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.zn, err = strconv.ParseFloat(row[2], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.indus, err = strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.chas, err = strconv.ParseFloat(row[4], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.nox, err = strconv.ParseFloat(row[5], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.rooms, err = strconv.ParseFloat(row[6], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.age, err = strconv.ParseFloat(row[7], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.dis, err = strconv.ParseFloat(row[8], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.rad, err = strconv.ParseFloat(row[9], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.tax, err = strconv.ParseFloat(row[10], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.ptratio, err = strconv.ParseFloat(row[11], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.lstat, err = strconv.ParseFloat(row[12], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}
		house.mv, err = strconv.ParseFloat(row[13], 64)
		if err != nil {
			fmt.Printf("Error parsing to number: %v\n", err)
			return nil, err
		}

		houses = append(houses, house)
	}

	return houses, nil
}

// Run Linear Regression model taking in data
//func runMultLinearRegression(houses []record) error {
//
//	trainingData, testingData, err := splitTrainTest(houses)
//	if err != nil {
//		fmt.Errorf("Error splitting data: %s", err)
//	}
//
//	// Split training data for multiple linear regression
//	Xtrain, Ytrain, err := splitPredResponse(trainingData)
//
//	// Perform multiple linear regression
//	coeff, err := stat.LinearRegression(Xtrain, Ytrain, nil, false)
//	if err != nil {
//		fmt.Errorf("Error running linear regression: %s", err)
//	}
//
//	// Separate response and predictors for testing data
//	Xtest, Ytest, err := splitPredResponse(testingData)
//
//	//fmt.Println("Response (mv):", Ytrain)
//	//fmt.Println("Predictor variables:", Xtrain)
//
//	return nil
//}

func runLinearRegression(predString string, pred []float64, resp []float64) error {
	// Split data for training and testing
	// Define the ratio of the dataset to use for training (e.g., 80% for training, 20% for testing)
	trainingRatio := 0.8
	// Calculate the number of samples for training and testing
	totalSamples := len(pred)
	trainingSamples := int(float64(totalSamples) * trainingRatio)

	Xtrain := pred[:trainingSamples]
	Xtest := pred[trainingSamples:]
	Ytrain := resp[:trainingSamples]
	Ytest := resp[trainingSamples:]

	// Perform linear regression
	alpha, beta := stat.LinearRegression(Xtrain, Ytrain, nil, false)
	r2 := stat.RSquared(Xtrain, Ytrain, nil, alpha, beta)

	var predictions []float64

	//y = alpha + beta*x
	for _, x := range Xtest {
		// Calculate the predicted 'mv' value using the linear regression model
		prediction := alpha + beta*x
		predictions = append(predictions, prediction)
	}

	// Calculate R2 value froms predicted values and test
	testR2 := stat.RSquaredFrom(predictions, Ytest, nil)

	fmt.Println(predString, " R2 value for train dataset :", r2)
	fmt.Println(predString, " R2 value for test dataset :", testR2)

	return nil
}

func main() {

	houses, err := parseCSV("./boston.csv")
	if err != nil {
		fmt.Printf("Error parsing file: %s", err)
	}

	for i := 0; i < 100; i++ {
		runSerial(houses)
	}

	for i := 0; i < 100; i++ {
		runConcurrent(houses)
	}

}

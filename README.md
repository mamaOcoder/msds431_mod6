# Week 6 Assignment: Exploring Concurrency

## Project Summary
This assignment uses concurrent programming in Go to analyze data from the [Boston Housing Study (1970)](http://lib.stat.cmu.edu/datasets/boston). This project utilizes the [gonum](https://github.com/gonum/gonum) package to test 3 linear regression models to predict median home value. Gonum/stat's package contains a built in LinearRegression function that only allows a single predictor value, so each of the 3 linear regression models built tests a different predictor feature. 

Note: My intial goal was to write code to test both multiple linear regression and another ML model such as Support Vector Machine or Decision Trees, however, I spent so much time trying to get multiple linear regression to work (unsuccessfully) that I ran out of time to implement another regression model. This code simply runs 3 linear regression models each using a different predictor. You will see some functions block-commented out as these were in pursuit of building the multiple linear regression model.

This assignment was meant to be a simple implementation of concurrency in Go, so this code does not perform operations, such as multi-fold cross-validation, bootstrap estimation or hyperparameter tuning that would traditionally be included in machine learning data analysis studies such as this. Given more time to implement, inclusion of computationally-intensive operations such as these could demostrate even more advantages to using concurrency.

## Benchmarking Results
Processing times were calcuted using the Go's test package. Both benchmark functions parse the CSV file before running the respective program 100 times. Results show that the concurrent version of the program is over 2x as fast.
> go test -bench=serial_test.go
```
PASS
ok      Mod6    0.534s
```
> go test -bench=concurrent_test.go
```
PASS
ok      Mod6    0.207s
```

## Files
### *main.go*
This file defines the parseCSV() and runLinearRegression() functions.  
parseCSV() takes in a string filename, parses each record in the CSV into a record object (type struct) and returns a slice of records.  
runLinearRegression() takes in a string for the predictor feature that is being used, a slice of values for the predictor and slice of response values. This function first splits the data into training (80%) and testing (20%) datasets, then runs gonum's stat.LinearRegression function on the training data which returns computed values for alpha and beta (y = alpha + beta*x). runLinearRegression() then uses these values to predict values for the test dataset and prints the R-squared values for both the training (stat.RSquared) and testing (stat.RSquaredFrom).

### *serial.go*
This file defines the runSerial() method. It takes in a slice of housing records, which it loops through to create 4 new slices- one for each of the 3 predictor values (nox, rooms, and crim) and the response values. It then calls runLinearRegression sequentially for each of the predictors.

### *serial_test.go*
This file runs the benchmark test for the serial program.

### *concurrent.go*
This file defines the runConcurrent() method. It takes in a slice of housing records, which it loops through to create 4 new slices- one for each of the 3 predictor values (nox, rooms, and crim) and the response values. It then calls runLinearRegression concurrently using sync.WaitGroup.

### *concurrent_test.go*
This file runs the benchmark test for the concurrent program.

### *prepData.go*
This file contains methods for prepping the data for multiple linear regression. Note that these methods are not used in this current iteration of the programs.

### *boston.csv* 
Input file for Boston housing data from Miller 1999.

### *mod6.exe*
Executable file of cross-compiled Go code for Mac/Windows.

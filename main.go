package main

// file structure
// report folder houses the models, functions and tests

import "surfapp/report"

func main() {

	// fetcher call the NewUrlFetcher method which is located in the report/fetch_report.go
	fetcher := report.NewUrlFetcher("http://noaa")
	// NoaaV2LineParser{} is a struct which implements the PareseLine interface from report/parse.go
	// Similar to the shape example, each shape has a struct (triangle, rectangle, circle), since there is an type Shape interface with the Area method. Each struct can call the methods within the Shape interface. In this sense  NoaaV2LineParser{} is like a triangle struct etc..
	lineParser := report.NoaaV2LineParser{}
	// NoaaMultilineParser{} is a struct etc..
	multiLineParser := report.NoaaMultilineParser{}
	//NoaaSurfReportCalculator{} is a struct etc..
	calculator := report.NoaaSurfReportCalculator{}
	reportRunner := report.ReportRunner{}
	// by convention an error value will be on the right and result will be on the left
	// r represents the result where reportRunner calls the method RunSurfReport and takes the defined parameters.
	//  Q : How is does peportRunner access a method? cannot find the interface declared for it..
	r, err := reportRunner.RunSurfReport(fetcher, multiLineParser, lineParser, calculator)
	//if err is not equal to nil, in relation to nil in an error nil is a valid value for one. Therefore a nil error represents no error.
	//panic is use to abort a function and returns an error value.
	if err != nil {
		panic(err)
	}
	// if there are no errors then we print the result for SwellPeriod
	println(r.SwellPeriod)
}

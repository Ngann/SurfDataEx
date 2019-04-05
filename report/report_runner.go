package report

//why do we have an empty struct?
// ReportRunner is an empty struct
type ReportRunner struct {
}

// ReportRunner has access to the ReportSurfReport which passes methods from the Fetch, Parse, Calculate interface and returns the SurfReport and error
func (ReportRunner) RunSurfReport(fetch Fetch, multiLineParser ParseLines, lineParser ParseLine, calculator Calculate) (SurfReport, error) {
	// data is the result which calls the GetData method which returns the string the response body string(body) or err
	data, err := fetch.GetData()
	// no error, it returns the SurfReport struct defined in surf_report.go
	if err != nil {
		return SurfReport{}, err
	}
	//lost me here
	buoyData, err := multiLineParser.ParseLines(data, lineParser)
	if err != nil {
		return SurfReport{}, err
	}
	// returns the calculated surf report base on the above buoyData
	// it calls the method CalculateSurfReport which takes a an array of BouyData struct and returns a SurfReport struct
	report := calculator.CalculateSurfReport(buoyData)
	return report, nil
}

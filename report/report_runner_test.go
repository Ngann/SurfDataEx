package report

import "testing"

func TestReportRunner_RunSurfReport(t *testing.T) {

	mockFetch := MockFetcher{
		data: sampleData,
	}
	lineParser := NoaaV2LineParser{}
	multiLineParser := NoaaMultilineParser{}
	calculator := NoaaSurfReportCalculator{}

	reportRunner := ReportRunner{}

	_, err := reportRunner.RunSurfReport(mockFetch, multiLineParser, lineParser, calculator)
	if err != nil {
		t.Error(err)
		return
	}

}

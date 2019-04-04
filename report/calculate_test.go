package report

import "testing"

func TestNoaaSurfReportCalculator_CalculateSurfReport(t *testing.T) {
	calculator := NoaaSurfReportCalculator{}

	report := calculator.CalculateSurfReport([]BouyData{
		{SwP: 10},
		{SwP: 20},
	})

	if report.SwellPeriod != 15.0 {
		t.Log("Expected period of 15.0, got", report.SwellPeriod)
		t.Fail()
	}
}

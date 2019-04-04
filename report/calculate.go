package report

type Calculate interface {
	CalculateSurfReport([]BouyData) SurfReport
}


type NoaaSurfReportCalculator struct {

}

func (NoaaSurfReportCalculator) CalculateSurfReport(data []BouyData) SurfReport {
	swellPeriod := float64(0)
	dataPoints := 0
	for _, dataPoint := range data {
		swellPeriod += dataPoint.SwP
		dataPoints++
	}

	averageSwellPeriod := swellPeriod/float64(dataPoints)

	return SurfReport{
		WaveSize:      0,
		WindDirection: "",
		SwellPeriod:   averageSwellPeriod,
		Score:         0,
	}
}

type MockSurfCalculator struct {
	Report SurfReport
}

func (c MockSurfCalculator) CalculateSurfReport([]BouyData) SurfReport {
	return c.Report
}

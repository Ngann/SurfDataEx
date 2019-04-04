package report

import "testing"

const sampleRow = `2019 04 02 16 00  1.8  1.8 14.8  0.2  3.4 WNW NNW      SWELL  9.3 286`

const sampleData = `#YY  MM DD hh mm WVHT  SwH  SwP  WWH  WWP SwD WWD  STEEPNESS  APD MWD
#yr  mo dy hr mn    m    m  sec    m  sec  -  degT     -      sec degT
2019 04 02 16 00  1.8  1.8 14.8  0.2  3.4 WNW NNW      SWELL  9.3 286
2019 04 02 15 00  1.7  1.7 16.0  0.2  3.4 WNW NNE      SWELL  9.4 288
2019 04 02 14 00  1.5  1.5 13.8  0.2  3.6 WNW NNW      SWELL  9.1 285
2019 04 02 13 00  1.7  1.7 17.4  0.2  3.8 WNW NNE      SWELL 10.1 288
2019 04 02 12 00  1.9  1.8 16.0  0.2  3.7 WNW NNW      SWELL 10.5 289`

func TestSpaceLineParser_Parse(t *testing.T) {
	lineParser := NoaaLineParser{}
	data, err := lineParser.Parse(sampleRow)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if data.Year != 2019 {
		t.Log("Expected year 2019, got", data.Year)
		t.Fail()
	}
}


func TestNoaaMultilineParser_ParseLines(t *testing.T) {
	multiLineParser := NoaaMultilineParser{}

	parsedLines, err := multiLineParser.ParseLines(sampleData, MockLineParser{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(parsedLines) != 5 {
		t.Log("Expected 5 lines, got", len(parsedLines))
		t.Fail()
	}
}

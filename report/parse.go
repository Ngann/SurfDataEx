package report

import (
	"strconv"
	"strings"
)

type ParseLine interface {
	Parse(line string) (BouyData, error)
}

type ParseLines interface {
	ParseLines(lines string, lineParser ParseLine) ([]BouyData, error)
}

type NoaaMultilineParser struct {
}

func (NoaaMultilineParser) ParseLines(lines string, lineParser ParseLine) ([]BouyData, error) {
	rows := strings.Split(lines, "\n")

	if len(rows) > 2 {
		rows = rows[2:]
	}
	var data []BouyData
	for _, row := range rows {
		bouyData, err := lineParser.Parse(row)
		if err != nil {
			return nil, err
		}
		data = append(data, bouyData)
	}
	return data, nil
}

type NoaaLineParser struct {
}

func (NoaaLineParser) Parse(line string) (BouyData, error) {
	components := strings.Fields(line)

	year, err := strconv.Atoi(components[0])
	if err != nil {
		return BouyData{}, err
	}
	return BouyData{
		Year: year,
	}, nil
}

type MockLineParser struct {
	MockData BouyData
}

func (p MockLineParser) Parse(line string) (BouyData, error) {
	return p.MockData, nil
}


type NoaaV2LineParser struct {
}

func (NoaaV2LineParser) Parse(line string) (BouyData, error) {
	components := strings.Fields(line)

	year, err := strconv.Atoi(components[0])
	if err != nil {
		return BouyData{}, err
	}
	return BouyData{
		Year: year,
	}, nil
}

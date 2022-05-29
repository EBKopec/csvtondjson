package domain

import (
	"gitlab.neoway.com.br/awesomeCSVtoNDJSON/internal/infrastructure/client/numbersAPI"
	"gitlab.neoway.com.br/awesomeCSVtoNDJSON/utils/misc"
)

type ServiceOutput struct {
	InputFile []InputFile
}

func NewServiceOutput(file []InputFile) *ServiceOutput {
	return &ServiceOutput{
		InputFile: file,
	}
}

//CreateOutput mount the output file
func (s *ServiceOutput) CreateOutput() ([]OutputFile, error) {
	var outputFile []OutputFile
	var eachLine OutputFile

	for _, line := range s.InputFile {
		ts, err := misc.ParseTime(line.Date)
		if err != nil {
			return nil, err
		}
		sourceIp, err := misc.ParseIP(line.SourceIp)
		if err != nil {
			return nil, err
		}
		urlParsed, err := misc.ParseURLx(line.TargetURL)
		if err != nil {
			return nil, err
		}
		url := *EnvVarsURL(urlParsed.Scheme, urlParsed.Host, urlParsed.Path, urlParsed.Opaque)

		size, err := misc.ParseStrToInt(line.TrafficSize)
		if err != nil {
			return nil, err
		}

		numberAPI := numbersAPI.NewNumbersAPI()
		note, err := Numbers.Request(numberAPI)
		if err != nil {
			return nil, err
		}
		eachLine = *EnvVarsOutputFile(ts, sourceIp, &url, size, note)
		outputFile = append(outputFile, eachLine)

	}

	return outputFile, nil
}

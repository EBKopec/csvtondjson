package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"gitlab.neoway.com.br/awesomeCSVtoNDJSON/domain"
	"log"
	"os"
)

var (
	version, build, date string
)

// TODO
// Additional requirements and notes
//---------------------------------
//- Implement it in Golang
//- The conversion tool has at least 2 arguments: input file name, output file name
//- The input file can be huge, the long processing tasks should be parallelized
//- Invalid lines should be displayed in the console
//- Use comments when and if necessary
//- Unit tests are welcome

func main() {

	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	serviceFile := domain.NewServiceInput(data)
	inputFile := domain.Input.InputFile2Json(serviceFile)

	serviceOutput := domain.NewServiceOutput(inputFile)
	outputFile, err := domain.Output.CreateOutput(serviceOutput)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(outputFile, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("JsonData %s\n\n", jsonData)
}

// Package main
//======================================================================================================
// File: cmd/service/main.go
// Authors: Everton Barbosa Kopec (Radar)
// Date: 30 may 2022
// Brief: Main module service - Conversion CSV to NDJSON
//
// COPYRIGHT © 2019 all rights reserved to Radar
//#======================================================================================================

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"gitlab.neoway.com.br/awesomeCSVtoNDJSON/domain"
	"log"
	"os"
)

// Additional requirements and notes
//---------------------------------
//- Implement it in Golang - DONE
//- The conversion tool has at least 2 arguments: input file name, output file name - DONE
//- The input file can be huge, the long processing tasks should be parallelized - DONE
//- Invalid lines should be displayed in the console - DONE
//- Use comments when and if necessary - DONE
//- Unit tests are welcome - DONE

func main() {

	//Args1: path to the csv file and Args2 path to destination file
	//The extension must be explicit
	source := os.Args[1]
	dest := os.Args[2]
	fmt.Printf("Source: %s\nDestination: %s\n", source, dest)
	f, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//Create service for input file
	serviceFile := domain.NewServiceInput(data)
	inputFile := domain.Input.InputFile2Json(serviceFile)

	//Create service por output file
	serviceOutput := domain.NewServiceOutput(inputFile)
	outputFile, outputErr := domain.Output.CreateOutput(serviceOutput)

	//Prepare error data if exists
	jsonDataErr, err := json.Marshal(outputErr)
	if err != nil {
		log.Fatal(err)
	}
	err = domain.Output.ErrOutput(serviceOutput, jsonDataErr)
	if err != nil {
		log.Fatal(err)
	}
	//Prepare json data to write an outputFile
	jsonData, err := json.Marshal(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	//Write the outputFile
	err = domain.Output.WriteOutput(serviceOutput, jsonData, dest)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

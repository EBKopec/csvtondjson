// Package domain
//======================================================================================================
// File: domain/outputFileService.go
// Authors: Everton Barbosa Kopec (Radar)
// Date: 30 may 2022
// Brief: OutputFileService domain module - Conversion CSV to NDJSON
//
// COPYRIGHT © 2019 all rights reserved to Radar
//#======================================================================================================
package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.neoway.com.br/awesomeCSVtoNDJSON/internal/infrastructure/client/numbersAPI"
	"gitlab.neoway.com.br/awesomeCSVtoNDJSON/utils/misc"
	"os"
	"sync"
)

type ServiceOutput struct {
	InputFile []InputFile
}

func NewServiceOutput(file []InputFile) *ServiceOutput {
	return &ServiceOutput{
		InputFile: file,
	}
}

//CreateOutput mount the output file and passing data for channels. Using 10 threads
func (s *ServiceOutput) CreateOutput() ([]OutputFile, []OutputFile) {
	var wg sync.WaitGroup
	outputFile := make([]OutputFile, 0)
	outputErr := make([]OutputFile, 0)
	threads := make(chan bool, 10)
	InFile := make(chan *InputFile, len(s.InputFile))
	OutFile := make(chan []OutputFile, 1)
	OutErr := make(chan []OutputFile, 1)

	for _, line := range s.InputFile {
		wg.Add(1)
		go deliveryData(line, &wg, &threads, InFile)
	}
	go parseData(InFile, OutFile, OutErr)

	wg.Wait()
	close(InFile)
	close(threads)

	outputErr = <-OutErr
	outputFile = <-OutFile

	close(OutFile)
	close(OutErr)

	return outputFile, outputErr
}

//WriteOutput write the output file
func (s *ServiceOutput) WriteOutput(data []byte, dest string) error {
	var lines []OutputFile
	var value = []byte("null")

	// return if data is empty
	if bytes.Equal(data, value) {
		return fmt.Errorf("there is no data to write\n")
	}
	f, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &lines)
	if err != nil {
		return err
	}

	for _, line := range lines {
		ln, err := json.Marshal(line)
		if err != nil {
			return err
		}

		_, err = f.WriteString(fmt.Sprintln(string(ln)))
		if err != nil {
			return err
		}
	}
	return nil
}

//ErrOutput print the output error in the console
func (s *ServiceOutput) ErrOutput(data []byte) error {
	var lines []OutputFile
	err := json.Unmarshal(data, &lines)
	if err != nil {
		return err
	}
	//Print in the console only if there are errors
	for _, line := range lines {
		ln, err := json.Marshal(line)
		if err != nil {
			return err
		}
		fmt.Println(string(ln))
	}
	return nil
}

//deliveryData control the threads for go routines
func deliveryData(data InputFile, wg *sync.WaitGroup, threads *chan bool, inFile chan *InputFile) {
	defer func() {
		wg.Done()
		<-*threads
	}()
	*threads <- true
	inFile <- &data
}

//parseData parse all data for go routines
func parseData(inFile chan *InputFile, outFile, outErr chan []OutputFile) {
	var file, fileError []OutputFile
	var ccErr bool

	for {
		data, open := <-inFile
		if open {
			ts, err := misc.ParseTime(data.Date)
			if err != nil {
				ccErr = true
			}
			sourceIp, err := misc.ParseIP(data.SourceIp)
			if err != nil {
				ccErr = true
			}
			urlParsed, err := misc.ParseURLx(data.TargetURL)
			if err != nil {
				ccErr = true
			}
			url := *EnvVarsURL(urlParsed.Scheme, urlParsed.Host, urlParsed.Path, urlParsed.Opaque)

			size, err := misc.ParseStrToInt(data.TrafficSize)
			if err != nil {
				ccErr = true
			}

			numberAPI := numbersAPI.NewNumbersAPI()
			note, err := Numbers.Request(numberAPI)
			if err != nil {
				ccErr = true
			}

			if ccErr {
				eachLineError := *EnvVarsOutputFile(ts, sourceIp, &url, size, note)
				fileError = append(fileError, eachLineError)
				ccErr = false
				continue
			} else {
				eachLine := *EnvVarsOutputFile(ts, sourceIp, &url, size, note)
				file = append(file, eachLine)
			}

		} else {
			ccErr = false
			outErr <- fileError
			outFile <- file
			break
		}
	}
}

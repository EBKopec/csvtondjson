package domain

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	inputFile, inputErrFile, outputFileW, messageFile string
)

func init() {
	inputFile = `2022-04-21T10:13:31Z,1.2.3.4,www.yahoo.com/abc,12000`
	inputErrFile = `2022-04-21T10:13:31Z,1.2.3.4.5,www.yahoo.com/abc,12000`
	outputFileW = `../utils/files/testWrite.ndjson`
	messageFile = `there is no data to write`

}
func TestCreateOutputFile_Success(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString(inputFile)

	data, err := csv.NewReader(&buffer).ReadAll()
	assert.NoError(t, err)
	assert.NotNil(t, data)

	serviceFile := NewServiceInput(data)
	assert.NotNil(t, serviceFile)
	inputFile := Input.InputFile2Json(serviceFile)
	assert.NotNil(t, inputFile)

	serviceOutput := NewServiceOutput(inputFile)
	assert.NotNil(t, serviceOutput)

	outputFile, outputErr := Output.CreateOutput(serviceOutput)
	assert.Nil(t, outputErr)
	assert.NotNil(t, outputFile)
}

func TestCreateOutputErr_Success(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString(inputErrFile)

	data, err := csv.NewReader(&buffer).ReadAll()
	assert.NoError(t, err)
	assert.NotNil(t, data)

	serviceFile := NewServiceInput(data)
	assert.NotNil(t, serviceFile)
	inputFile := Input.InputFile2Json(serviceFile)
	assert.NotNil(t, inputFile)

	serviceOutput := NewServiceOutput(inputFile)
	assert.NotNil(t, serviceOutput)

	outputFile, outputErr := Output.CreateOutput(serviceOutput)
	assert.NotNil(t, outputErr)
	assert.Nil(t, outputFile)
}

func TestErrOutput_Success(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString(inputErrFile)

	data, err := csv.NewReader(&buffer).ReadAll()
	assert.NoError(t, err)
	assert.NotNil(t, data)

	serviceFile := NewServiceInput(data)
	assert.NotNil(t, serviceFile)
	inputFile := Input.InputFile2Json(serviceFile)
	assert.NotNil(t, inputFile)

	serviceOutput := NewServiceOutput(inputFile)
	assert.NotNil(t, serviceOutput)

	outputFile, outputErr := Output.CreateOutput(serviceOutput)
	assert.NotNil(t, outputErr)
	assert.Nil(t, outputFile)

	jsonDataErr, err := json.Marshal(outputErr)
	assert.NoError(t, err)
	assert.NotNil(t, jsonDataErr)

	err = Output.ErrOutput(serviceOutput, jsonDataErr)
	assert.NoError(t, err)
}

func TestWriteOutput_Success(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString(inputFile)

	data, err := csv.NewReader(&buffer).ReadAll()
	assert.NoError(t, err)
	assert.NotNil(t, data)

	serviceFile := NewServiceInput(data)
	assert.NotNil(t, serviceFile)
	inputFile := Input.InputFile2Json(serviceFile)
	assert.NotNil(t, inputFile)

	serviceOutput := NewServiceOutput(inputFile)
	assert.NotNil(t, serviceOutput)

	outputFile, outputErr := Output.CreateOutput(serviceOutput)
	assert.Nil(t, outputErr)
	assert.NotNil(t, outputFile)

	jsonDataErr, err := json.Marshal(outputErr)
	assert.NoError(t, err)
	assert.NotNil(t, jsonDataErr)

	err = Output.ErrOutput(serviceOutput, jsonDataErr)
	assert.NoError(t, err)

	jsonData, err := json.Marshal(outputFile)
	assert.NoError(t, err)
	assert.NotNil(t, jsonData)

	err = Output.WriteOutput(serviceOutput, jsonData, outputFileW)
	assert.NoError(t, err)
}

func TestWriteOutput_Fail(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString(inputErrFile)

	data, err := csv.NewReader(&buffer).ReadAll()
	assert.NoError(t, err)
	assert.NotNil(t, data)

	serviceFile := NewServiceInput(data)
	assert.NotNil(t, serviceFile)
	inputFile := Input.InputFile2Json(serviceFile)
	assert.NotNil(t, inputFile)

	serviceOutput := NewServiceOutput(inputFile)
	assert.NotNil(t, serviceOutput)

	outputFile, outputErr := Output.CreateOutput(serviceOutput)
	assert.NotNil(t, outputErr)
	assert.Nil(t, outputFile)

	jsonDataErr, err := json.Marshal(outputErr)
	assert.NoError(t, err)
	assert.NotNil(t, jsonDataErr)

	err = Output.ErrOutput(serviceOutput, jsonDataErr)
	assert.NoError(t, err)

	jsonData, err := json.Marshal(outputFile)
	assert.NoError(t, err)
	assert.NotNil(t, jsonData)

	err = Output.WriteOutput(serviceOutput, jsonData, outputFileW)
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("%s\n", messageFile), err.Error())
}

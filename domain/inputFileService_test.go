package domain

import (
	"bytes"
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	file string
)

func init() {
	file = `2022-04-21T10:13:31Z,1.2.3.4,www.yahoo.com/abc,12000`

}

func TestServiceInput_InputFile2Json_Success(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString(file)

	data, err := csv.NewReader(&buffer).ReadAll()
	assert.NoError(t, err)
	assert.NotNil(t, data)

	serviceFile := NewServiceInput(data)
	assert.NotNil(t, serviceFile)
	inputFile := Input.InputFile2Json(serviceFile)
	assert.NotNil(t, inputFile)
}

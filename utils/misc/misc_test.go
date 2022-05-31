package misc

import (
	"github.com/stretchr/testify/assert"
	"net"
	"net/url"
	"testing"
)

var (
	testSourceIPParsed  = net.IP{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0x1, 0x2, 0x3, 0x4}
	testTargetURLParsed = &url.URL{
		Scheme: "http",
		Host:   "www.yahoo.com",
		Path:   "/abc",
		Opaque: "",
	}
)

const (
	testTrafficSizeInt       = 10
	testTrafficSizeStr       = "10"
	testTrafficSizeErr       = ""
	testTimestamp            = "2022-04-21T10:13:31Z"
	testTimestampErr         = "2022-04-21T"
	testTimestampEPOCH int64 = 1650536011
	testSourceIP             = "1.2.3.4"
	testSourceIPErr          = "1.2.3.256"
	testTargetURL            = "www.yahoo.com/abc"
	testTargetURLErr         = ""
)

func TestParseStrToInt_Success(t *testing.T) {
	number, err := ParseStrToInt(testTrafficSizeStr)
	assert.NoError(t, err)
	assert.Equal(t, testTrafficSizeInt, number)
}

func TestParseStrToInt_Fail(t *testing.T) {
	_, err := ParseStrToInt(testTrafficSizeErr)
	assert.Error(t, err)
}

func TestParseTime_Success(t *testing.T) {
	ts, err := ParseTime(testTimestamp)
	assert.NoError(t, err)
	assert.Equal(t, testTimestampEPOCH, ts)
}

func TestParseTime_Fail(t *testing.T) {
	_, err := ParseTime(testTimestampErr)
	assert.Error(t, err)
}

func TestParseIP_Success(t *testing.T) {
	sourceIp, err := ParseIP(testSourceIP)
	assert.NoError(t, err)
	assert.Equal(t, testSourceIPParsed, sourceIp)
}

func TestParseIP_Fail(t *testing.T) {
	_, err := ParseIP(testSourceIPErr)
	assert.Error(t, err)
}

func TestParseURLx_Success(t *testing.T) {
	urlParsed, err := ParseURLx(testTargetURL)
	assert.NoError(t, err)
	assert.Equal(t, testTargetURLParsed, urlParsed)
}
func TestParseURLx_Fail(t *testing.T) {
	_, err := ParseURLx(testTargetURLErr)
	assert.Error(t, err)
}

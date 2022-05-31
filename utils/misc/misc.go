// Package misc
//======================================================================================================
// File: utils/misc/misc.go
// Authors: Everton Barbosa Kopec (Radar)
// Date: 30 may 2022
// Brief: Miscellaneous for service application - Conversion CSV to NDJSON
//
// COPYRIGHT © 2019 all rights reserved to Radar
//#======================================================================================================

package misc

import (
	"fmt"
	"github.com/goware/urlx"
	"net"
	"net/url"
	"strconv"
	"time"
)

const (
	ISO8601 = "2006-01-02T15:04:05Z"
)

//ParseStrToInt parse string to int
func ParseStrToInt(s string) (int, error) {
	number, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return number, nil
}

//ParseTime parse string time to timestamp in EPOCH format
func ParseTime(s string) (int64, error) {
	resultTime, err := time.Parse(ISO8601, s)
	if err != nil {
		return 0, err
	}
	return resultTime.Unix(), nil
}

//ParseIP validate IPV4/IPV6
func ParseIP(s string) (net.IP, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return nil, fmt.Errorf("IP is not valid")
	}
	return ip, nil
}

//ParseURLx URLs
func ParseURLx(s string) (*url.URL, error) {
	urls, err := urlx.Parse(s)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

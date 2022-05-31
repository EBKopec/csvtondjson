// Package domain
//======================================================================================================
// File: domain/outputFileService.go
// Authors: Everton Barbosa Kopec (Radar)
// Date: 30 may 2022
// Brief: InputFile domain module - Conversion CSV to NDJSON
//
// COPYRIGHT © 2019 all rights reserved to Radar
//#======================================================================================================
package domain

type InputFile struct {
	Date        string `json:"date_iso"`
	SourceIp    string `json:"source_ip"`
	TargetURL   string `json:"target_url"`
	TrafficSize string `json:"traffic_size"`
}

func EnvVarsInput(date, sourceIP, targetURL, trafficSize string) *InputFile {
	return &InputFile{
		Date:        date,
		SourceIp:    sourceIP,
		TargetURL:   targetURL,
		TrafficSize: trafficSize,
	}
}

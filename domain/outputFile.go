// Package domain
//======================================================================================================
// File: domain/outputFile.go
// Authors: Everton Barbosa Kopec (Radar)
// Date: 30 may 2022
// Brief: OutputFile domain module - Conversion CSV to NDJSON
//
// COPYRIGHT © 2019 all rights reserved to Radar
//#======================================================================================================
package domain

import (
	"net"
)

type OutputFile struct {
	TimestampEpoch int64  `json:"ts"`
	SourceIP       net.IP `json:"source_ip"`
	URL            *URL   `json:"url"`
	Size           int    `json:"size"`
	Note           string `json:"note"`
}

func EnvVarsOutputFile(tsEpoch int64, sourceIp net.IP, URL *URL, size int, note string) *OutputFile {
	return &OutputFile{
		TimestampEpoch: tsEpoch,
		SourceIP:       sourceIp,
		URL:            URL,
		Size:           size,
		Note:           note,
	}
}

type URL struct {
	Scheme string `json:"Scheme"`
	Host   string `json:"Host"`
	Path   string `json:"Path"`
	Opaque string `json:"Opaque"`
}

func EnvVarsURL(scheme, host, path, opaque string) *URL {
	return &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
		Opaque: opaque,
	}
}

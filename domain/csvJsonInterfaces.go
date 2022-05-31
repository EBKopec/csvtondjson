// Package domain
//======================================================================================================
// File: domain/outputFileService.go
// Authors: Everton Barbosa Kopec (Radar)
// Date: 30 may 2022
// Brief: Functions interface handlers - Conversion CSV to NDJSON
//
// COPYRIGHT © 2019 all rights reserved to Radar
//#======================================================================================================
package domain

type Input interface {
	InputFile2Json() []InputFile
}

type Output interface {
	CreateOutput() ([]OutputFile, []OutputFile)
	WriteOutput(data []byte, dest string) error
	ErrOutput(data []byte) error
}

type Numbers interface {
	Request() (string, error)
}

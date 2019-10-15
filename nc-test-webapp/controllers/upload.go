package controllers

import (
	"github.com/mkotr/netcool-test-automation/nc-test-webapp/service"
)

type FileController struct {
	BaseController
	fileParser service.FileParser
}

func (f *FileController) ParseFile() {
	file, _, err := f.GetFile("file")

	if err != nil {
		f.Error(UNKNOWN, "An error occured while getting the file", err)
	}

	parser := service.FileParser{File: file}
	parsedResult, err := parser.ParseFile()

	if err != nil {
		f.Error(UNKNOWN, "An error occured while parsing the file", err)
	}

	f.Success(parsedResult, "Successfully parsed file.")
}

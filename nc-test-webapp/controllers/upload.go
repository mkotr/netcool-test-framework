package controllers

import (
	"github.com/astaxie/beego"
)

type FileController struct {
	BaseController
}

type (f *FileController) ParseFile() {
	file, _, err := f.GetFile("file")
}
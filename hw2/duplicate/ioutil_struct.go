package duplicate

import (
	"io/fs"
	"io/ioutil"
)

type ioutilStruct struct {
}

func (i ioutilStruct) ReadDir(nameDir string) ([]fs.FileInfo, error) {
	return ioutil.ReadDir(nameDir)
}

func (i ioutilStruct) ReadFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

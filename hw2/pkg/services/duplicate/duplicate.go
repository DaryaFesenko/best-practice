package duplicate

import (
	"io/fs"
	"io/ioutil"
	"os"

	"hw2/pkg/models"
	fa "hw2/pkg/services/fileaction"

	log "github.com/sirupsen/logrus"
)

type ioutilStruct struct {
}

func (i ioutilStruct) ReadDir(nameDir string) ([]fs.FileInfo, error) {
	return ioutil.ReadDir(nameDir)
}

func (i ioutilStruct) ReadFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

func GetDuplicateFile(pathDir string) ([]string, error) {
	l := log.WithField("FuncName", "GetDuplicateFile").WithField("path", pathDir)
	l.Debugf("run get duplicates")

	i := &ioutilStruct{}
	f := &fa.FileActions{FS: i}

	files, err := f.GetAllFiles(pathDir)
	listDuplicate := []string{}
	listOrigin := []models.FileInfo{}

	if err != nil {
		return listDuplicate, err
	}

	for _, file := range files.List {
		exist := false
		for _, val := range listOrigin {
			if val.FileName == file.FileName && val.Hash_md5 == file.Hash_md5 && val.Hash_sha256 == file.Hash_sha256 && !exist {
				listDuplicate = append(listDuplicate, file.Path)
				exist = true
			}
		}

		if !exist {
			item := models.FileInfo{FileName: file.FileName, Hash_md5: file.Hash_md5, Hash_sha256: file.Hash_sha256}
			listOrigin = append(listOrigin, item)
		}
	}

	return listDuplicate, nil
}

func RemoveDuplicate(duplicate []string) error {
	l := log.WithField("FuncName", "RemoveDuplicate")
	l.Debugf("run remove duplicates")

	for _, item := range duplicate {
		err := os.Remove(item)

		if err != nil {
			return err
		}
	}
	return nil
}

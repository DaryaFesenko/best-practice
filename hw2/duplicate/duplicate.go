package duplicate

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func GetDuplicateFile(pathDir string) ([]string, error) {
	l := log.WithField("FuncName", "GetDuplicateFile").WithField("path", pathDir)
	l.Debugf("run get duplicates")

	i := &ioutilStruct{}
	f := &FileActions{fs: i}

	files, err := f.getAllFiles(pathDir)
	listDuplicate := []string{}
	listOrigin := []fileInfo{}

	if err != nil {
		return listDuplicate, err
	}

	for _, file := range files.list {
		exist := false
		for _, val := range listOrigin {
			if val.fileName == file.fileName && val.hash_md5 == file.hash_md5 && val.hash_sha256 == file.hash_sha256 && !exist {
				listDuplicate = append(listDuplicate, file.path)
				exist = true
			}
		}

		if !exist {
			item := fileInfo{fileName: file.fileName, hash_md5: file.hash_md5, hash_sha256: file.hash_sha256}
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

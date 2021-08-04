package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hw2/pkg/models"
	"io/fs"
	"time"
)

const plug = 0

type FileInfo struct {
	name  string
	isDir bool
}

// Линтер gomnd - убрала магические цифры в константу
func (f *FileInfo) Name() string {
	return f.name
}

func (f *FileInfo) IsDir() bool {
	return f.isDir
}

func (f *FileInfo) Size() int64 {
	return int64(plug)
}

func (f *FileInfo) Mode() fs.FileMode {
	return fs.FileMode(plug)
}

func (f *FileInfo) ModTime() time.Time {
	return time.Now()
}

func (f *FileInfo) Sys() interface{} {
	return int64(plug)
}

func FillFiles(dirNames, fileNames []string) []fs.FileInfo {
	returns := make([]fs.FileInfo, 0)

	for _, val := range dirNames {
		tmp := &FileInfo{name: val, isDir: true}
		returns = append(returns, tmp)
	}

	for _, val := range fileNames {
		tmp := &FileInfo{name: val, isDir: false}
		returns = append(returns, tmp)
	}

	return returns
}

// Линтер gocriric - параметры типа стринг, можно объединить
// Линтер revive - исправила название переменных
func AddFileInfo(path, fileName string, list *models.FilesInfo) {
	b := make([]byte, 0)
	h1 := md5.New()
	h1.Write(b)
	hashMd5 := hex.EncodeToString(h1.Sum(nil))

	h2 := sha256.New()
	h2.Write(b)
	hashSha256 := hex.EncodeToString(h2.Sum(nil))

	list.AddItem(fileName, path, hashMd5, hashSha256)
}

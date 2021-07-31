package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hw2/pkg/models"
	"io/fs"
	"time"
)

type FileInfo struct {
	name  string
	isDir bool
}

func (f *FileInfo) Name() string {
	return f.name
}

func (f *FileInfo) IsDir() bool {
	return f.isDir
}

func (f *FileInfo) Size() int64 {
	return int64(5)
}

func (f *FileInfo) Mode() fs.FileMode {
	return fs.FileMode(6)
}
func (f *FileInfo) ModTime() time.Time {
	return time.Now()
}
func (f *FileInfo) Sys() interface{} {
	return int64(5)
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

func AddFileInfo(path string, fileName string, list *models.FilesInfo) {
	b := make([]byte, 0)
	h1 := md5.New()
	h1.Write(b)
	hash_md5 := hex.EncodeToString(h1.Sum(nil))

	h2 := sha256.New()
	h2.Write(b)
	hash_sha256 := hex.EncodeToString(h2.Sum(nil))

	list.AddItem(fileName, path, hash_md5, hash_sha256)
}

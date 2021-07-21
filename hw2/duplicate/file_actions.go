package duplicate

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"sync"

	log "github.com/sirupsen/logrus"
)

type FS interface {
	ReadDir(string) ([]fs.FileInfo, error)
	ReadFile(string) ([]byte, error)
}

type FileActions struct {
	fs FS
	m  sync.Mutex
	wg sync.WaitGroup

	fi FilesInfo
}

func (f *FileActions) getAllFiles(path string) (FilesInfo, error) {
	l := log.WithField("FuncName", "getAllFiles").WithField("path", path)
	l.Debugf("run get all files")

	f.fi = FilesInfo{}

	files, err := f.fs.ReadDir(path)

	if err != nil {
		return f.fi, fmt.Errorf("directory %s  does not open. error: %v", path, err)
	}

	f.readDirectory(path, files)

	return f.fi, nil
}

func (f *FileActions) readDirectory(path string, files []fs.FileInfo) error {
	l := log.WithField("FuncName", "readDirectory").WithField("path", path)
	for _, file := range files {
		newPath := path + "/" + file.Name()
		if !file.IsDir() {
			l.Debug("read file:", newPath)
			fileByte, err := f.fs.ReadFile(newPath)
			if err != nil {
				return fmt.Errorf("file: %s error: %s", path, err)
			}
			f.wg.Add(1)
			go f.addFileInfo(fileByte, newPath, file.Name())
		} else {
			l.Debug("read directory:", newPath)
			dir, err := f.fs.ReadDir(newPath)

			if err != nil {
				return fmt.Errorf("directory %s  does not open. err: %v", newPath, err)
			}

			f.readDirectory(newPath, dir)
		}
	}

	f.wg.Wait()
	return nil
}

func (f *FileActions) addFileInfo(file []byte, path string, fileName string) {
	defer f.wg.Done()
	h1 := md5.New()
	h1.Write(file)
	hash_md5 := hex.EncodeToString(h1.Sum(nil))

	h2 := sha256.New()
	h2.Write(file)
	hash_sha256 := hex.EncodeToString(h2.Sum(nil))

	f.m.Lock()
	f.fi.AddItem(fileName, path, hash_md5, hash_sha256)
	f.m.Unlock()
}

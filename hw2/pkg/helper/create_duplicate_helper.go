package helper

import (
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

func CreateDuplicateFile(path string) []string {
	list := make(map[string]string)
	listCopy := []string{}
	files, _ := ioutil.ReadDir(path)

	if _, err := os.Stat(path + "/copy"); err == nil {
		os.RemoveAll(path + "/copy")
	}
	// Линтер wsl - названия пременных с с ошибками должны быть разными (но в проекте реальном такого не видела)
	if errCopy := os.Mkdir(path+"/copy", os.ModePerm); errCopy != nil {
		log.Println(errCopy)
	}

	readDirectory(path, list, files)
	n := rand.Intn(len(list)-1) + 1

	for name, pathFile := range list {
		if err := copy(path+"/copy", pathFile, name); err != nil {
			log.Fatalf("cant't copy file: %s", err)
		}

		listCopy = append(listCopy, name)
		n--

		if n == 0 {
			break
		}
	}

	return listCopy
}

func readDirectory(path string, list map[string]string, files []fs.FileInfo) {
	for _, file := range files {
		newPath := path + "/" + file.Name()
		if !file.IsDir() {
			list[file.Name()] = newPath
		} else {
			dir, _ := ioutil.ReadDir(newPath)
			readDirectory(newPath, list, dir)
		}
	}
}

func copy(pathDir, path, name string) error {
	file, _ := os.Open(path)

	copyFile, _ := os.Create(pathDir + "/" + name)

	// Линтер errcheck  - не было проверки на ошибку
	_, err := io.Copy(copyFile, file)
	if err != nil {
		return err
	}

	file.Close()
	copyFile.Close()

	return nil
}

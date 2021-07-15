package additional

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
	err := os.Mkdir(path+"/copy", os.ModePerm)

	if err != nil {
		log.Println(err)
	}

	readDirectory(path, list, files)

	n := rand.Intn(len(list)-1) + 1

	for i := 0; i < n; i++ {
		for name, pathFile := range list {
			copy(path+"/copy", pathFile, name)
			listCopy = append(listCopy, path+"\\copy\\"+name)
			break
		}
	}

	return listCopy
}

func readDirectory(path string, list map[string]string, files []fs.FileInfo) {
	for _, file := range files {
		newPath := path + "\\" + file.Name()
		if !file.IsDir() {
			list[file.Name()] = newPath
		} else {
			dir, _ := ioutil.ReadDir(newPath)
			readDirectory(newPath, list, dir)
		}
	}
}

func copy(pathDir string, path string, name string) {
	file, _ := os.Open(path)

	copyFile, _ := os.Create(pathDir + "/" + name)

	io.Copy(copyFile, file)

	file.Close()
	copyFile.Close()
}

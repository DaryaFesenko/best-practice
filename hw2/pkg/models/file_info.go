package models

type FilesInfo struct {
	List []FileInfo
}

// Линтер revive - исправила названия переменных
type FileInfo struct {
	FileName   string
	HashMd5    string
	HashSha256 string
	Path       string
}

// Линтер gocriric - все параметры типа стринг, можно объединить
func (f *FilesInfo) AddItem(fileName, path, hashMd5, hashSha256 string) {
	item := FileInfo{
		FileName:   fileName,
		Path:       path,
		HashMd5:    hashMd5,
		HashSha256: hashSha256,
	}

	f.List = append(f.List, item)
}

func (f *FilesInfo) FindItemByPath(path string) bool {
	for _, val := range f.List {
		if val.Path == path {
			return true
		}
	}

	return false
}

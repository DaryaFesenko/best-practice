package models

type FilesInfo struct {
	List []FileInfo
}
type FileInfo struct {
	FileName    string
	Hash_md5    string
	Hash_sha256 string
	Path        string
}

func (f *FilesInfo) AddItem(fileName string, path string, hash_md5 string, hash_sha256 string) {
	item := FileInfo{
		FileName:    fileName,
		Path:        path,
		Hash_md5:    hash_md5,
		Hash_sha256: hash_sha256,
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

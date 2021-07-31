package duplicate

type FilesInfo struct {
	list []fileInfo
}

func (f *FilesInfo) AddItem(fileName string, path string, hash_md5 string, hash_sha256 string) {
	item := fileInfo{
		fileName:    fileName,
		path:        path,
		hash_md5:    hash_md5,
		hash_sha256: hash_sha256,
	}

	f.list = append(f.list, item)
}

func (f *FilesInfo) FindItemByPath(path string) bool {
	for _, val := range f.list {
		if val.path == path {
			return true
		}
	}

	return false
}

type fileInfo struct {
	fileName    string
	hash_md5    string
	hash_sha256 string
	path        string
}

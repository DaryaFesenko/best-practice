package duplicate

import (
	"hw2/additional"
	"hw2/duplicate/mocks"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// мокаю файловую систему, проверяю логику формирования объектов для поиска дубликатов
func TestDuplicate_GetAllFiles(t *testing.T) {
	fsMock := &mocks.FS{}
	fa := &FileActions{fs: fsMock}

	fsMock.On("ReadDir", "test_dir").Return(FillFiles([]string{"copy"}, []string{"aaaa", "gggg"}), nil)
	fsMock.On("ReadDir", "test_dir/copy").Return(FillFiles([]string{}, []string{"aaaa", "gggg"}), nil)

	b := make([]byte, 0)
	fsMock.On("ReadFile", "test_dir/aaaa").Return(b, nil)
	fsMock.On("ReadFile", "test_dir/gggg").Return(b, nil)
	fsMock.On("ReadFile", "test_dir/copy/aaaa").Return(b, nil)
	fsMock.On("ReadFile", "test_dir/copy/gggg").Return(b, nil)

	res, err := fa.getAllFiles("test_dir")

	list := FilesInfo{}

	AddFileInfo("test_dir/aaaa", "aaaa", &list)
	AddFileInfo("test_dir/gggg", "gggg", &list)
	AddFileInfo("test_dir/copy/aaaa", "aaaa", &list)
	AddFileInfo("test_dir/copy/gggg", "gggg", &list)

	require.NoError(t, err)

	for _, val := range res.list {
		if !list.FindItemByPath(val.path) {
			t.Fatal("item not found :", val.path)
		}
	}
}

func TestGetDuplicate(t *testing.T) {
	testCases := []struct {
		Name      string
		path      string
		duplicate []string
	}{
		{
			Name:      "no duplicate",
			path:      "./test_dir",
			duplicate: []string{},
		},
		{
			Name:      "no folder",
			path:      "./test_dir2",
			duplicate: []string{},
		},
		{
			Name:      "test1",
			path:      "/home/d/projects/gb/best-practice/hw2/test_dir",
			duplicate: []string{},
		},
		{
			Name:      "test2",
			path:      "/home/d/projects/gb/best-practice/hw2/test_dir",
			duplicate: []string{},
		},
		{
			Name:      "test3",
			path:      "/home/d/projects/gb/best-practice/hw2/test_dir",
			duplicate: []string{},
		},
	}

	out, _ := GetDuplicateFile(testCases[0].path)
	assert.Empty(t, out, testCases[0].duplicate)

	_, err := GetDuplicateFile(testCases[1].path)
	assert.NotEqual(t, err, nil)

	for i := 2; i < len(testCases); i++ {
		tt := testCases[i]

		tt.duplicate = additional.CreateDuplicateFile(tt.path)

		res, err := GetDuplicateFile(tt.path)

		assert.Equal(t, err, nil)

		out := make([]string, 0)
		for _, val := range res {
			lastIndex := strings.LastIndex(val, "/")

			out = append(out, val[lastIndex+1:])
		}

		sort.Strings(tt.duplicate)
		sort.Strings(out)
		assert.Equal(t, out, tt.duplicate)
	}
}

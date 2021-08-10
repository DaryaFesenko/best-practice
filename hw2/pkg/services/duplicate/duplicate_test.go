package duplicate

import (
	"sort"
	"strings"
	"testing"

	"best-practice/hw2/pkg/mocks"

	"best-practice/hw2/pkg/helper"
	"best-practice/hw2/pkg/models"
	fa "best-practice/hw2/pkg/services/fileaction"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// мокаю файловую систему, проверяю логику формирования объектов для поиска дубликатов
func TestDuplicate_GetAllFiles(t *testing.T) {
	fsMock := &mocks.FS{}
	fa := &fa.FileActions{FS: fsMock}

	test_dir := "test_dir"
	test_dir_copy := "test_dir/copy"
	ReadDir := "ReadDir"
	ReadFile := "ReadFile"

	fsMock.On(ReadDir, test_dir).Return(helper.FillFiles([]string{"copy"}, []string{"aaaa", "gggg"}), nil)
	fsMock.On(ReadDir, test_dir_copy).Return(helper.FillFiles([]string{}, []string{"aaaa", "gggg"}), nil)

	b := make([]byte, 0)
	fsMock.On(ReadFile, test_dir+"/aaaa").Return(b, nil)
	fsMock.On(ReadFile, test_dir+"/gggg").Return(b, nil)
	fsMock.On(ReadFile, test_dir_copy+"/aaaa").Return(b, nil)
	fsMock.On(ReadFile, test_dir_copy+"/gggg").Return(b, nil)

	res, err := fa.GetAllFiles(test_dir)

	list := models.FilesInfo{}

	helper.AddFileInfo(test_dir+"/aaaa", "aaaa", &list)
	helper.AddFileInfo(test_dir+"/gggg", "gggg", &list)
	helper.AddFileInfo(test_dir_copy+"/aaaa", "aaaa", &list)
	helper.AddFileInfo(test_dir_copy+"/gggg", "gggg", &list)

	require.NoError(t, err)

	for _, val := range res.List {
		if !list.FindItemByPath(val.Path) {
			t.Fatal("item not found :", val.Path)
		}
	}
}

func TestGetDuplicate(t *testing.T) {
	path := "/home/d/projects/gb/best-practice/hw2/test/test_dir"
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
			path:      path,
			duplicate: []string{},
		},
		{
			Name:      "test2",
			path:      path,
			duplicate: []string{},
		},
		{
			Name:      "test3",
			path:      path,
			duplicate: []string{},
		},
	}

	out, _ := GetDuplicateFile(testCases[0].path)
	assert.Empty(t, out, testCases[0].duplicate)

	_, err := GetDuplicateFile(testCases[1].path)
	assert.NotEqual(t, err, nil)

	for i := 2; i < len(testCases); i++ {
		tt := testCases[i]

		tt.duplicate = helper.CreateDuplicateFile(tt.path)

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

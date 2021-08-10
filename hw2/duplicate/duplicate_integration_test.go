// +build integration

package duplicate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// проверка интеграции файловой системы
func TestDuplicate_ReadDirectory(t *testing.T) {
	path := "/home/d/projects/gb/best-practice/hw2/test_integration"

	i := &ioutilStruct{}
	f := &FileActions{fs: i}

	expected := FilesInfo{}
	AddFileInfo("/home/d/projects/gb/best-practice/hw2/test_integration/copy/aaaa", "aaaa", &expected)

	// вынуждена тестить тот же метод, что и для мок, но входные данные такие,
	// что буду просто тестировать открытие одной папки и одного файла
	res, err := f.getAllFiles(path)

	require.NoError(t, err)
	require.Len(t, res.list, len(expected.list))
	require.Equal(t, res.list[0], expected.list[0])
}

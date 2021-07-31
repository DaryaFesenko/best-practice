// +build integration

package duplicate

import (
	"testing"

	"hw2/pkg/helper"

	"hw2/pkg/models"
	fa "hw2/pkg/services/fileaction"

	"github.com/stretchr/testify/require"
)

// проверка интеграции файловой системы
func TestDuplicate_ReadDirectory(t *testing.T) {
	path := "/home/d/projects/gb/best-practice/hw2/test/test_integration"

	i := &ioutilStruct{}
	f := &fa.FileActions{FS: i}

	expected := models.FilesInfo{}
	helper.AddFileInfo(path+"/copy/aaaa", "aaaa", &expected)

	// вынуждена тестить тот же метод, что и для мок, но входные данные такие,
	// что буду просто тестировать открытие одной папки и одного файла
	res, err := f.GetAllFiles(path)

	require.NoError(t, err)
	require.Len(t, res.List, len(expected.List))
	require.Equal(t, res.List[0], expected.List[0])
}

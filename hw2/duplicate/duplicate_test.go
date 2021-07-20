package duplicate

import (
	"hw2/additional"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

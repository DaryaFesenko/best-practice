package duplicate

import (
	"hw2/additional"
	"testing"
)

func TestGetDuplicate(t *testing.T) {
	testCases := []struct {
		Name      string
		path      string
		duplicate []string
	}{
		{
			Name:      "no duplicate",
			path:      "C:\\goDir",
			duplicate: []string{},
		},
		{
			Name:      "no folder",
			path:      "C:\\goDir2",
			duplicate: []string{},
		},
		{
			Name:      "test1",
			path:      "C:\\goDir",
			duplicate: []string{},
		},
		{
			Name:      "test2",
			path:      "C:\\goDir",
			duplicate: []string{},
		},
		{
			Name:      "test3",
			path:      "C:\\goDir",
			duplicate: []string{},
		},
	}

	t.Run(testCases[0].Name, func(t *testing.T) {
		out, _ := GetDuplicateFile(testCases[0].path)

		if len(out) > 0 {
			t.Fatalf("got %v, but want %v", out, testCases[0].duplicate)
		}
	})

	t.Run(testCases[1].Name, func(t *testing.T) {
		_, err := GetDuplicateFile(testCases[1].path)

		if err == nil {
			t.Fatalf("an error was expected: directory does not open")
		}
	})

	for i := 2; i < len(testCases); i++ {
		testCases[i].duplicate = additional.CreateDuplicateFile(testCases[i].path)
	}

	for i := 2; i < len(testCases); i++ {
		tt := testCases[i]
		t.Run(tt.Name, func(t *testing.T) {
			out, err := GetDuplicateFile(tt.path)

			if err != nil {
				t.Fatalf("ERROR: %v", err)
			}

			if len(out) == len(tt.duplicate) {
				for _, val := range tt.duplicate {
					exist := false
					for _, valOut := range out {
						if valOut == val {
							exist = true
							break
						}
					}

					if !exist {
						t.Fatalf("got %v, but want %v", out, tt.duplicate)
					}
				}
			}
		})
	}
}

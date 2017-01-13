package fileparser

import (
	"testing"
)

func TestNewFile(t *testing.T) {
	dir := "test"
	title := "title"
	ext := "txt"

	mock := NewFile(dir, title, ext)

	if mock.Dir != dir && mock.Title != title && mock.Ext != ext {
		t.Error("Fail struct init")
	}
}

func TestGetFilePath(t *testing.T) {
	for _, test := range []struct {
		file     File
		response string
	}{
		{File{[]byte{}, "path", "file", "ext"}, "path/file.ext"},
		{File{[]byte{}, "path", "file", ""}, "path/file"},
	} {
		if test.file.GetFilePath() != test.response {
			t.Error(
				"Expect", test.file.GetFilePath(),
				"got", test.response,
			)
		}
	}
}

package fileparser

import (
	"io/ioutil"
)

type File struct {
	Content []byte
	Dir     string
	Title   string
	Ext     string
}

func (f *File) GetFilePath() string {
	basePath := f.Dir + "/" + f.Title

	if f.Ext != "" {
		return basePath + "." + f.Ext
	}

	return basePath
}

func (f *File) SaveFile() error {
	path := f.GetFilePath()
	return ioutil.WriteFile(path, f.Content, 0777)
}

func LoadFile(dir string, title string, ext string) (File, error) {
	fileInstance := NewFile(dir, title, ext)
	Content, err := ioutil.ReadFile(fileInstance.GetFilePath())

	if err != nil {
		return File{}, err
	}

	fileInstance.Content = Content

	return fileInstance, nil
}

func NewFile(dir string, title string, ext string) File {
	return File{
		Dir:     dir,
		Title:   title,
		Ext:     ext,
		Content: []byte{},
	}
}

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

func (f *File) SaveFile() {
	path := f.GetFilePath()
	err := ioutil.WriteFile(path, f.Content, 0777)

	if err != nil {
		panic(err)
	}
}

func LoadFile(dir string, title string, ext string) File {
	fileInstance := NewFile(dir, title, ext)
	Content, err := ioutil.ReadFile(fileInstance.GetFilePath())

	if err != nil {
		panic(err)
	}

	fileInstance.Content = Content

	return fileInstance
}

func NewFile(dir string, title string, ext string) File {
	return File{
		Dir:     dir,
		Title:   title,
		Ext:     ext,
		Content: []byte{},
	}
}

package suricata

import (
	"bytes"
	"os"
	"path/filepath"
)

type File struct {
	FilePath   string
	buffer     *bytes.Buffer
	OutputFile *os.File
}

func NewFile() *File {
	f := &File{
		buffer: &bytes.Buffer{},
	}
	return f
}

func (file *File) SetFile(name string, path string) error {
	file.FilePath = filepath.Join(path, name)

	os.Remove(file.FilePath)
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	newFile, err := os.Create(file.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	file.OutputFile = newFile
	return nil
}
func (file *File) Write(chunk []byte) error {
	if file.OutputFile == nil {
		return nil
	}
	_, err := file.OutputFile.Write(chunk)
	return err
}

func (file *File) Close() error {
	return file.OutputFile.Close()
}

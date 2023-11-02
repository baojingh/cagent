package suricata

import (
	"bytes"
	"errors"
	"fmt"
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

// Create the path if not exist
func (file *File) SetFile(name string, path string) error {
	file.FilePath = filepath.Join(path, name)

	os.Mkdir(path, os.ModePerm)
	// If the file already exists, it is truncated.
	// If the file does not exist, it is created with mode 0666
	newFile, err := os.Create(file.FilePath)
	if err != nil {
		log.Errorf("Fail to create file %s, err: %v", file.FilePath, err)
		return err
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
	if file.OutputFile == nil {
		msg := fmt.Sprintf("File %s is nil", file.FilePath)
		log.Errorf(msg)
		return errors.New(msg)
	}
	return file.OutputFile.Close()
}

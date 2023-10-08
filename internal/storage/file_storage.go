package storage

import (
	"errors"
	"io"
	"os"
)

type FileStorage struct {
	fileName string
}

func NewFileStorage(fileName string) *FileStorage {
	return &FileStorage{
		fileName: fileName,
	}
}

func (f *FileStorage) Store(data []byte) error {
	if f.fileName == "" {
		return errors.New("file not specified")
	}

	file, err := os.Create(f.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileStorage) Restore() ([]byte, error) {
	if f.fileName == "" {
		return nil, errors.New("file not specified")
	}

	file, err := os.Open(f.fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}

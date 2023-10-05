package storage

import (
	"bufio"
	"bytes"
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

func (f *FileStorage) Store(data [][]byte) error {
	if f.fileName == "" {
		return errors.New("file not specified")
	}

	file, err := os.Create(f.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := range data {
		if _, err := writer.Write(data[i]); err != nil {
			return err
		}

		if err := writer.WriteByte('\n'); err != nil {
			return err
		}
	}

	return writer.Flush()
}

func (f *FileStorage) Restore() ([][]byte, error) {
	if f.fileName == "" {
		return nil, errors.New("file not specified")
	}

	file, err := os.Open(f.fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bb, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes.Split(bb, []byte("\n")), nil
}

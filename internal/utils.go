package internal

import (
	"fmt"
	"os"
)

func CreateDir(dirname string) error {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		return os.Mkdir(dirname, os.ModePerm)
	}
	return fmt.Errorf("Unable to create directory %s", dirname)
}

func MustCreateDir(dirpath string) {
	err := CreateDir(dirpath)
	if err != nil {
		panic(fmt.Errorf("Should have created directory: %s. Apparently it already exists", dirpath))
	}
}

func WriteToFile(filepath string, data string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(data)
	f.Sync()
	return nil
}

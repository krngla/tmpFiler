// tmpfiler is a package that provides a simple way to create and read temporary files
package tmpfiler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	subpath = "tmpfiler"
)

func OpenRead(name string) (string, error) {
	path := filepath.Join(os.TempDir(), subpath, name)
	file, err := os.OpenFile(path, os.O_RDONLY, 0444)
	if err != nil {
		return "", errors.New("failed to open file: " + err.Error())
	}
	output, err := bufio.NewReader(file).ReadString('\n')
	if err != nil {
		return "", errors.New("failed to read file: " + err.Error())
	}
	file.Close()
	return output, nil
}
func OpenWrite(name string, message string) (int, error) {
	path := filepath.Join(os.TempDir(), subpath, name)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return 0, errors.New("failed to open file: " + err.Error())
	}
	n, err := fmt.Fprint(file, message)
	if err != nil {
		return n, errors.New("failed to write to file: " + err.Error())
	}
	file.Close()
	return n, nil
}

func DeleteFile(name string) error {
	path := filepath.Join(os.TempDir(), subpath, name)
	err := os.Remove(path)
	if err != nil {
		return errors.New("failed to delete file: " + err.Error())
	}
	return nil
}

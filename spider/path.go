package spider

import (
	"os"
	"path/filepath"
)

func GetRunPath() (string, error) {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return path, err
}

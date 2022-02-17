package helpers

import (
	"path"
	"path/filepath"
	"runtime"
)

func GetRoot() string {
	_, basepath, _, _ := runtime.Caller(0)
	dirpath := path.Join(path.Dir(basepath))

	return filepath.Dir(dirpath)
}

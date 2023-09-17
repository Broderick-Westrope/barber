package internal

import (
	"path/filepath"
	"runtime"
)

func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
    return filepath.Join(filepath.Dir(b), "..")
}
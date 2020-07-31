package filepath

import (
	"os"
	"runtime"
)

func PathExists(path string) bool {
	f, err := os.Stat(path)
	if err == nil && f != nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// GetCurrentFilePath get current file path
func GetCurrentFilePath() string {
	_, filePath, _, _ := runtime.Caller(1)
	return filePath
}

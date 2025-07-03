package idebug

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type CallerData struct {
	PC       uintptr
	Filepath string
	Line     int
	OK       bool
}

func GetCallerPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

func GetCallerInfo() CallerData {
	res := CallerData{}
	res.PC, res.Filepath, res.Line, res.OK = runtime.Caller(1)
	return res
}

func GetProjectRootRelativePath() string {
	res := CallerData{}
	_, dir, _, _ := runtime.Caller(1)
	rel := "./"
	for {
		if _, err := os.Stat(filepath.Join(res.Filepath, "go.mod")); err == nil {
			break
		}

		// 检查是否已到达根目录
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			rel = ""
		}

		// 继续向上查找
		dir = parentDir
		rel = fmt.Sprintf("%s../", rel)
	}

	return rel
}

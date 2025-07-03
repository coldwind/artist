package idebug

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
	_, filename, _, _ := runtime.Caller(1)
	fmt.Println(filename)
	dir := filepath.Dir(filename)
	up := 0
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			break
		}

		// 检查是否已到达根目录
		parentDir := filepath.Dir(dir)
		fmt.Println(parentDir)
		if parentDir == dir {
			up = -1
			break
		}

		// 继续向上查找
		dir = parentDir
		up++

		if up > 100 {
			break
		}
	}

	return "./" + strings.Repeat("../", up)
}

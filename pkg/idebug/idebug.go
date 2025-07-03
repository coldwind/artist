package idebug

// idebug在主要用于在项目开发期间调试使用
// 方法在打包发布前需删除或不在打包文件内
// 例如在_test.go文件中使用

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

// 获取当前调用该函数的代码所在目录相对项目跟目录的位置
func GetProjectRootRelativePath() string {
	return GetProjectRootRelativePathBySkip(2)
}

// 跟据skip来获取调用代码相对项目跟目录的位置
func GetProjectRootRelativePathBySkip(skip int) string {
	_, filename, _, _ := runtime.Caller(skip)
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

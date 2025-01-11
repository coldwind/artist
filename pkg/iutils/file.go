package iutils

import (
	"fmt"
	"os"
	"path/filepath"
)

func MkdirOnNonExist(filePath string) error {
	// 获取文件所在目录
	dirPath := filepath.Dir(filePath)

	// 检查目录是否存在，如果不存在则创建
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
	}

	return nil
}

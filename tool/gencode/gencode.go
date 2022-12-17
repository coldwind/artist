package gencode

import (
	"fmt"
	"os"
	"time"

	"github.com/coldwind/artist/pkg/ihttp"
	"github.com/dablelv/go-huge-util/zip"
)

func Run(name, dst string) {
	// 获取远端代码
	codeByte, err := ihttp.Get("github.com/coldwind/artist/tool/gencode/zip/project_code.zip", nil, nil)
	if err != nil {
		fmt.Println("zip file not found.")
	}
	tempFileName := fmt.Sprintf("./%d.zip", time.Now().Unix())
	file, err := os.Create(tempFileName)
	if err != nil {
		fmt.Println("create temp zip file failure.")
	}
	file.Write(codeByte)
	file.Close()
	err = zip.Unzip(tempFileName, dst)
	if err != nil {
		fmt.Println("unzip failure.")
	}
}

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
	codeByte, err := ihttp.Get("", nil, nil)
	if err != nil {
		fmt.Println("zip file not found.")
		return
	}
	tempFileName := fmt.Sprintf("./%d.zip", time.Now().Unix())
	file, err := os.Create(tempFileName)
	if err != nil {
		fmt.Println("create temp zip file failure.")
		return
	}
	file.Write(codeByte)
	file.Close()
	err = zip.Unzip(tempFileName, dst)
	if err != nil {
		fmt.Println("unzip failure.")
		return
	}
}

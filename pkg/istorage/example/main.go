package main

import (
	"bytes"
	"fmt"

	"github.com/coldwind/artist/pkg/istorage"
)

func main() {
	s := istorage.New(istorage.Qiniu, &istorage.Config{
		SecretId:      "",
		SecretKey:     "",
		UseHTTPS:      false,
		UseCdnDomains: false,
	})
	b := []byte("hello, this is qiniu cloud")
	bio := bytes.NewReader(b)
	res, err := s.PutFromStream("yqyn", "avatar/1985/test.log", bio, int64(len(b)))
	fmt.Println(res, err)
}

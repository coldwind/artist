package main

import (
	"fmt"

	"github.com/coldwind/artist/pkg/isms"
)

func main() {
	s := isms.New(isms.QCloud, &isms.Config{
		SecretId:   "",
		SecretKey:  "",
		SdkAppId:   "",
		Sign:       "",
		TemplateId: "",
	})
	res, err := s.Send("+8613816207221", []string{"6666"})
	fmt.Println(res, err)
}

package main

import (
	"tool/excel2json"
	"tool/gencode"

	"github.com/desertbit/grumble"
)

func main() {
	var app *grumble.App = grumble.New(&grumble.Config{
		Name:        "tool",
		Description: "开发工具集",
	})

	app.AddCommand(&grumble.Command{
		Name: "excel2json",
		Help: "excel转json工具",
		Flags: func(f *grumble.Flags) {
			f.String("s", "src", "./", "excel配置文件目录")
			f.String("d", "dst", "./", "生成的json文件目录")

		},
		Run: func(c *grumble.Context) error {
			src := c.Flags.String("src")
			dst := c.Flags.String("dst")
			excel2json.Run(src, dst)
			return nil
		},
	})

	app.AddCommand(&grumble.Command{
		Name: "gencode",
		Help: "生成基础代码工具",
		Flags: func(f *grumble.Flags) {
			f.String("n", "name", "./", "项目名 go.mod中的名字")
			f.String("d", "dst", "./", "项目文件根目录")

		},
		Run: func(c *grumble.Context) error {
			dst := c.Flags.String("dst")
			name := c.Flags.String("name")
			gencode.Run(name, dst)
			return nil
		},
	})

	app.Run()
}

package main

import (
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
			Run(src, dst)
			return nil
		},
	})

	app.Run()
}

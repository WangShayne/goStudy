/*
	以tree列出文件夹目录
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func getDir(dirPath string, deep int) (err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	if deep == 1 {

		fmt.Printf("!---%s\n", filepath.Base(dirPath))
	}
	sep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			fmt.Printf("|")
			for i := 0; i < deep; i++ {
				fmt.Printf("    |")
			}
			fmt.Printf("----%s\n", fi.Name())
			getDir(dirPath+sep+fi.Name(), deep+1)
			continue
		}
		fmt.Printf("|")
		for i := 0; i < deep; i++ {
			fmt.Printf("    |")
		}
		fmt.Printf("----%s\n", fi.Name())
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "TREE"
	app.Usage = "tree all dir and file"
	app.Action = func(c *cli.Context) error {
		// 默认当前目录
		var dir string = "."
		if c.NArg() > 0 {
			dir = c.Args()[0]
		}
		// 获取目录
		getDir(dir, 1)
		return nil
	}
	app.Run(os.Args)
}

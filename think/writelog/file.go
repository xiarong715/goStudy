package main

import (
	"fmt"
	"os"
)

type FileOut struct {
	FileName string
}

func (fileout *FileOut) Write(msg string) {
	out, e := os.OpenFile(fileout.FileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if e != nil {
		fmt.Fprintf(os.Stdout, "%s%s%s", "create ", fileout.FileName, " file error.")
		return
	}
	defer out.Close()
	fmt.Fprintf(out, "%s", msg) // 格式化信息到文件中，也可用out.Write([]byte)
}

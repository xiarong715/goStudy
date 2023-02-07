package main

import "os"

func main() {
	to := TerminalOut{
		DeviceName: os.Stdout,
	}
	fo := FileOut{
		FileName: "./test.log",
	}
	Writelog(&to, "hello\n") // 写终端
	Writelog(&fo, "world\n") // 写文件
}

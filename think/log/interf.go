package main

type WriteLog interface {
	Write(string)
}

func Writelog(wl WriteLog, msg string) {
	wl.Write(msg)
}

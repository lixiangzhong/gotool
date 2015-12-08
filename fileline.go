package gotool

import (
	"runtime"
)

func FileLine() (string, int) {
	_, file, line, ok := runtime.Caller(1) //这里的1 为一次func封装，多一次就要加1
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return file, line
}

package gotool

import (
	"github.com/axgle/mahonia"
)

func GBKtoUTF8(gbk string) string {
	return mahonia.NewDecoder("GB18030").ConvertString(gbk)
}
func UTF8toGBK(utf8 string) string {
	return mahonia.NewEncoder("GB18030").ConvertString(utf8)
}

package util

import (
	"regexp"
	"strings"
)

// DelChinese 删除字符串中文
func DelChinese(s string) string {
	chiReg := regexp.MustCompile("[\u4e00-\u9fa5]")
	return chiReg.ReplaceAllString(s, "")
}

// DelChineseAndBrackets 删除字符串中文和括号
func DelChineseAndBrackets(s string) string {
	reg := regexp.MustCompile("[\u4e00-\u9fa5]|[(]|[)]")
	return reg.ReplaceAllString(s, "")
}

// DelSpace 删除字符串空格
func DelSpace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// GetNumberStr 只要数字
func GetNumberStr(s string) string {
	reg := regexp.MustCompile("[^0-9]")
	return reg.ReplaceAllString(s, "")
}

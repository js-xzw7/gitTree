package helper

import (
	//"external_api/enums"
	"strings"
	"unicode/utf8"
)

type StringHelper interface {
	ArrayReplace(sliceParams []string, index int, s string, new string, n int) string
	TrimLastChar(s string) string
}

type defaultStringHelper struct {
}

func NewStringHelper() defaultStringHelper {
	return defaultStringHelper{}
}

/**
 *  ArrayReplace
 *  @Description:切片批量替换
 *  @param splice 需要被批量替换的切片
 *  @param index  从切片中的第几个开始
 *  @param s      需要替换的文本
 *  @param new    替换换的文本
 *  @param n      替换几个，-1，默认全部替换
 *  @return string  返回
 */
func (h defaultStringHelper) ArrayReplace(sliceParams []string, index int, s string, new string, n int) string {
	sliceLen := len(sliceParams)
	if index+1 >= sliceLen {
		return s
	}
	//获取切片数据
	oldString := sliceParams[index]
	newString := strings.Replace(s, oldString, new, -1)
	index++
	if index == sliceLen {
		return newString
	}
	return h.ArrayReplace(sliceParams, index, newString, new, n)
}

/**
 *  SliceChangeKeyCase
 *  @Description:
 *  @receiver receiver
 */
//func (receiver *SliceHelper) ChangeKeyCase(params string, CaseType int) {
//
//	if enums.NumberEnumObject.CaseTypeUpper == CaseType {
//
//	}
//
//}

//删除最后一个元素
func (h defaultStringHelper) TrimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

package helper

import (
	"regexp"
	"strconv"
)

type NumberHelper interface {
	IsNum(s string) bool
	IsSciNum(num1, num2 string) bool
	IsDec(s string) bool
	IsInt(s string) bool
}

type defaultNumberHelper struct {
}

func NewNumberHelper() defaultNumberHelper {
	return defaultNumberHelper{}
}

/**
 *  IsNum
 *  @Description: 判断是不是数字类型：整数或者小数
 *  @param s
 *  @return bool
 */
func (num defaultNumberHelper) IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

/**
 *  isSciNum
 *  @Description: 是否为科学计数法
 *  @param num1
 *  @param num2
 *  @return bool
 */
func (h defaultNumberHelper) IsSciNum(num1, num2 string) bool {
	// e 前后字符串长度为0 是错误的
	if len(num1) == 0 || len(num2) == 0 {
		return false
	}
	// e 后面必须是整数，前面可以是整数或小数  4  +
	return (h.IsInt(num1) || h.IsDec(num1)) && h.IsInt(num2)
}

/**
 *  isDec
 *  @Description: 判断是否为小数
 *  @param s
 *  @return bool
 */
func (h defaultNumberHelper) IsDec(s string) bool {
	// eg: 11.15, -0.15, +10.15, 3., .15,
	// err: +. 0..
	match1, _ := regexp.MatchString(`^[\+-]?\d*\.\d+$`, s)
	match2, _ := regexp.MatchString(`^[\+-]?\d+\.\d*$`, s)
	return match1 || match2
}

/**
 *  isInt
 *  @Description: 判断是否为整数
 *  @param s
 *  @return bool
 */
func (h defaultNumberHelper) IsInt(s string) bool {
	match, _ := regexp.MatchString(`^[\+-]?\d+$`, s)
	return match
}

package helper

import (
	//"external_api/tool/Logger"
	"reflect"
	"strconv"
)

/**
*  SliceHelper
*  @Description:
 */
type SliceHelper interface {
	InArray(obj interface{}, target interface{}) bool
	ColumnBySlice(mapParams []map[string]interface{}, keyParams string) []interface{}
	InSliceString(keyParams string, sliceParams []string) bool
	GetIntersection(sliceParams1 []string, sliceParams2 []string) []string
	ToSliceStringByInterface(params []interface{}) []string
	ToSliceStringByInt(params []int) []string
	DifferenceSliceByString(sliceOne []string, sliceTwo []string) []string
}

type defaultSliceHelper struct {
}

func NewSliceHelper() defaultSliceHelper {
	return defaultSliceHelper{}
}

//查找字符是否在数组中
func (h defaultSliceHelper) InArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

/**
*  GetKeys
*  @Description: 获取map的key对应的
*  @receiver receiver
*  @param mapParams
*  @return []string
 */
func (h defaultSliceHelper) ColumnBySlice(mapParams []map[string]interface{}, keyParams string) []interface{} {
	slice := make([]interface{}, 0)
	for _, value := range mapParams {
		//判断key是否存在
		if val, ok := value[keyParams]; !ok {
			//Logger.LoggersHelperObject.PanicAndLog("ColumnBySlice.37", "此key不存在")
		} else {
			slice = append(slice, val)
		}
	}
	return slice
}

/**
 *  inSlice
 *  @Description: 判断字符串是否在字符串切片中
 *  @receiver receiver
 *  @param mapParams
 *  @param keyParams
 *  @return bool
 */
func (h defaultSliceHelper) InSliceString(keyParams string, sliceParams []string) bool {
	for _, e := range sliceParams {
		if e == keyParams {
			return true
		}
	}
	return false
}

/**
 *  GetIntersection
 *  @Description: 求交集
 *  @receiver receiver
 *  @param sliceParams1
 *  @param sliceParams2
 *  @return []string
 */
func (h defaultSliceHelper) GetIntersection(sliceParams1 []string, sliceParams2 []string) []string {
	intersection := make([]string, 0)
	//获取两个参数的长度
	len1 := len(sliceParams1)
	len2 := len(sliceParams2)
	if 0 == len1 || 0 == len2 {
		return intersection
	}
	tmpLen := 0
	tmpShort := 0
	if len2 >= len1 {
		tmpLen = len2
		tmpShort = len1
	} else {
		tmpLen = len1
		tmpShort = len2
	}
	lenSliceParams := make([]string, tmpLen)
	shortSliceParams := make([]string, tmpShort)
	if len2 >= len1 {
		copy(lenSliceParams, sliceParams2)
		copy(shortSliceParams, sliceParams1)
	} else {
		copy(lenSliceParams, sliceParams1)
		copy(shortSliceParams, sliceParams2)
	}
	tmpMap := make(map[string]string)
	for key, value := range shortSliceParams {
		tmpMap[value] = strconv.Itoa(key)
	}
	//判断key是否存在
	for i := 0; i < tmpLen; i++ {
		item := sliceParams1[i]
		if _, ok := tmpMap[item]; ok {
			intersection = append(intersection, item)
		}
	}
	return intersection
}

/**
 *  ToSliceString
 *  @Description: []interface{} 转为 []string
 *  @receiver receiver
 *  @param params
 *  @return []string
 */
func (h defaultSliceHelper) ToSliceStringByInterface(params []interface{}) []string {
	strings := make([]string, 0)
	for _, value := range params {
		strings = append(strings, InterfaceHelperObject.ToString(value))
	}
	return strings
}

/**
 *  ToSliceString
 *  @Description: []interface{} 转为 []string
 *  @receiver receiver
 *  @param params
 *  @return []string
 */
func (h defaultSliceHelper) ToSliceStringByInt(params []int) []string {
	strings := make([]string, 0)
	for _, value := range params {
		strings = append(strings, InterfaceHelperObject.ToString(value))
	}
	return strings
}

/**
 *  DifferenceSliceByString
 *  @Description: 如果在A中不在B中，如果在B中但是不在A中的数据集合
 *  @receiver h
 *  @param sliceOne
 *  @param sliceTwo
 *  @return []string
 */
func (h defaultSliceHelper) DifferenceSliceByString(sliceOne []string, sliceTwo []string) []string {
	differ := make([]string, 0)
	if len(sliceOne) == 0 && len(sliceTwo) == 0 {
		return differ
	}
	mapOne := make(map[string]string)
	mapTwo := make(map[string]string)
	for _, i2 := range sliceOne {
		mapOne[i2] = i2
	}
	for _, i3 := range sliceTwo {
		mapTwo[i3] = i3
	}
	for i, _ := range mapOne {
		_, isHas := mapTwo[i]
		if !isHas {
			differ = append(differ, i)
		}
	}
	for i, _ := range mapTwo {
		_, isHas := mapOne[i]
		if !isHas {
			differ = append(differ, i)
		}
	}
	return differ
}

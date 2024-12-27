package helper

/**
*  MapHelper
*  @Description:
 */
type MapHelper interface {
	GetKeys(mapParams map[string]interface{}) []string
	GetKeysByInterface(mapParams map[interface{}]interface{}) []string
	GetKeysByInt(mapParams map[int]int) []string
	DifferenceMapKeyByString(mapOne map[string]interface{}, mapTwo map[string]interface{}) map[string]interface{}
}

type defaultMapHelper struct {
}

func NewMapHelper() defaultMapHelper {
	return defaultMapHelper{}
}

/**
 *  GetKeys
 *  @Description: 获取map的key
 *  @receiver receiver
 *  @param mapParams
 *  @return []string
 */
func (h defaultMapHelper) GetKeys(mapParams map[string]interface{}) []string {
	slice := make([]string, 0)
	for i, _ := range mapParams {
		slice = append(slice, i)
	}
	return slice
}

/**
 *  GetKeys1
 *  @Description:
 *  @receiver receiver
 *  @param mapParams
 *  @return []string
 */
func (h defaultMapHelper) GetKeysByInterface(mapParams map[interface{}]interface{}) []string {
	slice := make([]string, 0)
	if mapParams == nil {
		return slice
	}
	for key1, _ := range mapParams {
		slice = append(slice, InterfaceHelperObject.ToString(key1))
	}
	return slice
}

/**
 *  GetKeysByInt
 *  @Description:
 *  @receiver receiver
 *  @param mapParams
 *  @return []string
 */
func (h defaultMapHelper) GetKeysByInt(mapParams map[int]int) []string {
	slice := make([]string, 0)
	if mapParams == nil {
		return slice
	}
	for key1, _ := range mapParams {
		slice = append(slice, InterfaceHelperObject.ToString(key1))
	}
	return slice
}

//func (receiver *MapHelper) name() {
//
//}

/**
 *  DifferenceMapKeyByString
 *  @Description: 如果在A中不在B中，如果在B中但是不在A中的数据集合
 *  @receiver h
 *  @param sliceOne
 *  @param sliceTwo
 *  @return []string
 */
func (h defaultMapHelper) DifferenceMapKeyByString(mapOne map[string]interface{}, mapTwo map[string]interface{}) map[string]interface{} {
	differ := make(map[string]interface{}, 0)
	if len(mapOne) == 0 && len(mapTwo) == 0 {
		return differ
	}
	for i, value := range mapOne {
		_, isHas := mapTwo[i]
		if !isHas {
			differ[i] = value
		}
	}
	for i, value := range mapTwo {
		_, isHas := mapOne[i]
		if !isHas {
			differ[i] = value
		}
	}
	return differ
}

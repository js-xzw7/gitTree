package helper

type ListHelper interface {
	ListByInterface(dataParams interface{}) []map[string]string
}
type defaultListHelper struct {
}

func NewListHelper() defaultListHelper {
	return defaultListHelper{}
}

/**
 *  ListByInterface
 *  @Description: list列表格式转换
 *  @receiver receiver
 *  @param params
 *  @return []map[string]string
 */
func (h defaultListHelper) ListByInterface(dataParams interface{}) []map[string]string {
	toString := InterfaceHelperObject.ToString(dataParams)
	if toString == "[]" || toString == "" {
		return make([]map[string]string, 0)
	}
	params := dataParams.([]interface{})
	tmp := make([]map[string]string, 0)
	for _, v := range params {
		//转换
		m := v.(map[string]interface{})
		tmpMap := make(map[string]string)
		for k2, v2 := range m {
			tmpMap[k2] = InterfaceHelperObject.ToString(v2)
		}
		tmp = append(tmp, tmpMap)
	}
	return tmp
}

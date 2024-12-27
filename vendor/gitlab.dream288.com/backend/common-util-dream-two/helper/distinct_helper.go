package helper

type DistinctHelp struct {
}

var DistinctHelpObject DistinctHelp

//
func (d *DistinctHelp) SortMapData(data []string) []string {
	result := []string{}
	for i := range data {
		flag := true
		for j := range result {
			if data[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, data[i])
		}
	}
	return result
}

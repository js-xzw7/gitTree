package helper

import "github.com/shopspring/decimal"

type FloatHelper interface {
	GetKeepDecimal(params float64, number int32) float64
}

type defaultFloatHelper struct {
}

func NewFloatHelper() defaultFloatHelper {
	return defaultFloatHelper{}
}

/**
 *  GetKeepDecimal
 *  @Description: 浮点数保留几位小数
 *  @receiver receiver
 *  @param params
 *  @param number
 *  @return float64
 */
func (receiver defaultFloatHelper) GetKeepDecimal(params float64, number int32) float64 {
	v1, _ := decimal.NewFromFloat(params).Round(number).Float64()
	return v1
}

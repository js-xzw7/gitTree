package helper

import (
	"github.com/google/uuid"
)

type SerialNumber interface {
}

type defaultSerialNumber struct {
}

func NewSerialNumber() defaultSerialNumber {
	return defaultSerialNumber{}
}

//生成Uuid通过时间
func (n defaultSerialNumber) GetUuidByTime() (string, error) {
	// V1 基于时间
	u1, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return u1.String(), err
}

//生成uuid通过随机数
func (n defaultSerialNumber) GetUuidByRandom() (string, error) {
	// V1 基于时间
	u1, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return u1.String(), err
}

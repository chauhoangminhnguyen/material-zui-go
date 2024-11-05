package mz_number

import (
	"fmt"
	"math"
)

type ZNumber float64

func ZuiNumber(value float64) *ZNumber {
	result := ZNumber(value)
	return &result
}

func (value *ZNumber) Get() float64 {
	return value.GetFloat64()
}

func (value *ZNumber) GetInt() int {
	return int(*value)
}

func (value *ZNumber) GetInt8() int8 {
	return int8(*value)
}

func (value *ZNumber) GetInt16() int16 {
	return int16(*value)
}

func (value *ZNumber) GetInt32() int32 {
	return int32(*value)
}

func (value *ZNumber) GetInt64() int64 {
	return int64(*value)
}

func (value *ZNumber) GetFloat32() float32 {
	return float32(*value)
}

func (value *ZNumber) GetFloat64() float64 {
	return float64(*value)
}

func (value *ZNumber) GetIntPart() *ZNumber {
	intPart, _ := splitNumber(value.Get())
	return ZuiNumber(float64(intPart))
}

func (value *ZNumber) GetFracPart() *ZNumber {
	_, fracPart := splitNumber(value.Get())
	return ZuiNumber(float64(fracPart))
}

func (value *ZNumber) ToString() string {
	return fmt.Sprintf("%v", *value)
}

func (value *ZNumber) Round(decimals int) *ZNumber {
	multiplier := math.Pow(10, float64(decimals))
	result := math.Round(value.Get()*multiplier) / multiplier
	return ZuiNumber(result)
}

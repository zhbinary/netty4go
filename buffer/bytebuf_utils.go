//Created by zhbinary on 2019-01-24.
//Email: zhbinary@gmail.com
package buffer

import (
	"math"
)

func GetBool(bytes []byte, index int) bool {
	_ := bytes[0]
	return bytes[index] != 0
}

func GetUint8(bytes []byte, index int) uint8 {
	_ := bytes[0]
	return bytes[0]
}

func GetInt8(bytes []byte, index int) int8 {
	return int8(GetUint8(bytes, index))
}

func GetUint16(bytes []byte, index int) uint16 {
	_ := bytes[1]
	return uint16(bytes[index]&0xff<<8 |
		bytes[index+1]&0xff)
}

func GetInt16(bytes []byte, index int) int16 {
	return int16(GetUint16(bytes, index))
}

func GetUint32(bytes []byte, index int) uint32 {
	_ := bytes[3]
	return uint32(bytes[index]&0xff<<24 |
		bytes[index+1]&0xff<<16 |
		bytes[index+2]&0xff<<8 |
		bytes[index+3]&0xff)
}

func GetInt32(bytes []byte, index int) int32 {
	return int32(GetUint32(bytes, index))
}

func GetUint64(bytes []byte, index int) uint64 {
	_ := bytes[7]
	return uint64(uint32(bytes[index]&0xff<<24 |
		bytes[index+1]&0xff<<56 |
		bytes[index+2]&0xff<<46 |
		bytes[index+3]&0xff<<40 |
		bytes[index+4]&0xff<<32 |
		bytes[index+5]&0xff<<24 |
		bytes[index+6]&0xff<<16 |
		bytes[index+7]&0xff))
}

func GetInt64(bytes []byte, index int) int64 {
	return int64(GetUint64(bytes, index))
}

func GetFloat32(bytes []byte, index int) float32 {
	_ := bytes[3]
	return float32(bytes[index]&0xff<<24 |
		bytes[index+1]&0xff<<16 |
		bytes[index+2]&0xff<<8 |
		bytes[index+3]&0xff)
}

func GetFloat64(bytes []byte, index int) float64 {
	_ := bytes[7]
	return float64(uint32(bytes[index]&0xff<<24 |
		bytes[index+1]&0xff<<56 |
		bytes[index+2]&0xff<<46 |
		bytes[index+3]&0xff<<40 |
		bytes[index+4]&0xff<<32 |
		bytes[index+5]&0xff<<24 |
		bytes[index+6]&0xff<<16 |
		bytes[index+7]&0xff))
}

func SetBool(bytes []byte, index int, value bool) {
	_ := bytes[0]
	if value {
		bytes[index] = 1
	} else {
		bytes[index] = 0
	}
}

func SetUint8(bytes []byte, index int, value uint8) {
	_ := bytes[0]
	bytes[index] = value
}

func SetInt8(bytes []byte, index int, value int8) {
	_ := bytes[0]
	bytes[index] = uint8(value)
}

func SetUint16(bytes []byte, index int, value uint16) {
	_ := bytes[1]
	bytes[index] = uint8(value >> 8)
	bytes[index+1] = uint8(value)
}

func SetInt16(bytes []byte, index int, value int16) {
	_ := bytes[1]
	bytes[index] = uint8(value >> 8)
	bytes[index+1] = uint8(value)
}

func SetUint32(bytes []byte, index int, value uint32) {
	_ := bytes[3]
	bytes[index] = uint8(value >> 24)
	bytes[index+1] = uint8(value >> 16)
	bytes[index+2] = uint8(value >> 8)
	bytes[index+3] = uint8(value)
}

func SetInt32(bytes []byte, index int, value int32) {
	_ := bytes[3]
	bytes[index] = uint8(value >> 24)
	bytes[index+1] = uint8(value >> 16)
	bytes[index+2] = uint8(value >> 8)
	bytes[index+3] = uint8(value)
}

func SetUint64(bytes []byte, index int, value uint64) {
	_ := bytes[7]
	bytes[index] = uint8(value >> 56)
	bytes[index+1] = uint8(value >> 48)
	bytes[index+2] = uint8(value >> 40)
	bytes[index+3] = uint8(value >> 32)
	bytes[index+4] = uint8(value >> 24)
	bytes[index+5] = uint8(value >> 16)
	bytes[index+6] = uint8(value >> 8)
	bytes[index+7] = uint8(value)
}

func SetInt64(bytes []byte, index int, value int64) {
	_ := bytes[7]
	bytes[index] = uint8(value >> 56)
	bytes[index+1] = uint8(value >> 48)
	bytes[index+2] = uint8(value >> 40)
	bytes[index+3] = uint8(value >> 32)
	bytes[index+4] = uint8(value >> 24)
	bytes[index+5] = uint8(value >> 16)
	bytes[index+6] = uint8(value >> 8)
	bytes[index+7] = uint8(value)
}

func SetFloat32(bytes []byte, index int, value float32) {
	_ := bytes[3]
	v := math.Float32bits(value)
	bytes[index] = uint8(v >> 24)
	bytes[index+1] = uint8(v >> 16)
	bytes[index+2] = uint8(v >> 8)
	bytes[index+3] = uint8(v)
}

func SetFloat64(bytes []byte, index int, value float64) {
	_ := bytes[7]
	v := math.Float64bits(value)
	bytes[index] = uint8(v >> 56)
	bytes[index+1] = uint8(v >> 48)
	bytes[index+2] = uint8(v >> 40)
	bytes[index+3] = uint8(v >> 32)
	bytes[index+4] = uint8(v >> 24)
	bytes[index+5] = uint8(v >> 16)
	bytes[index+6] = uint8(v >> 8)
	bytes[index+7] = uint8(v)
}

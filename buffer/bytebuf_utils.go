//Created by zhbinary on 2019-01-24.
//Email: zhbinary@gmail.com
package buffer

import (
	"encoding/binary"
	"math"
)

func GetBool(bytes []byte, index int) bool {
	return bytes[index] != 0
}

func GetUint8(bytes []byte, index int) uint8 {
	return bytes[index]
}

func GetInt8(bytes []byte, index int) int8 {
	return int8(GetUint8(bytes, index))
}

func GetUint16(bytes []byte, index int) uint16 {
	return binary.BigEndian.Uint16(bytes[index:])
}

func GetInt16(bytes []byte, index int) int16 {
	return int16(GetUint16(bytes, index))
}

func GetUint32(bytes []byte, index int) uint32 {
	return binary.BigEndian.Uint32(bytes[index:])
}

func GetInt32(bytes []byte, index int) int32 {
	return int32(GetUint32(bytes, index))
}

func GetUint64(bytes []byte, index int) uint64 {
	return binary.BigEndian.Uint64(bytes[index:])
}

func GetInt64(bytes []byte, index int) int64 {
	return int64(GetUint64(bytes, index))
}

func GetFloat32(bytes []byte, index int) float32 {
	return math.Float32frombits(GetUint32(bytes, index))
}

func GetFloat64(bytes []byte, index int) float64 {
	return math.Float64frombits(GetUint64(bytes, index))
}

func SetBool(bytes []byte, index int, value bool) {
	if value {
		bytes[index] = 1
	} else {
		bytes[index] = 0
	}
}

func SetUint8(bytes []byte, index int, value uint8) {
	bytes[index] = value
}

func SetInt8(bytes []byte, index int, value int8) {
	bytes[index] = uint8(value)
}

func SetUint16(bytes []byte, index int, value uint16) {
	binary.BigEndian.PutUint16(bytes[index:], value)
}

func SetInt16(bytes []byte, index int, value int16) {
	SetUint16(bytes, index, uint16(value))
}

func SetUint32(bytes []byte, index int, value uint32) {
	binary.BigEndian.PutUint32(bytes[index:], value)
}

func SetInt32(bytes []byte, index int, value int32) {
	SetUint32(bytes, index, uint32(value))
}

func SetUint64(bytes []byte, index int, value uint64) {
	binary.BigEndian.PutUint64(bytes[index:], value)
}

func SetInt64(bytes []byte, index int, value int64) {
	SetUint64(bytes, index, uint64(value))
}

func SetFloat32(bytes []byte, index int, value float32) {
	v := math.Float32bits(value)
	binary.BigEndian.PutUint32(bytes[index:], v)
}

func SetFloat64(bytes []byte, index int, value float64) {
	v := math.Float64bits(value)
	binary.BigEndian.PutUint64(bytes[index:], v)
}

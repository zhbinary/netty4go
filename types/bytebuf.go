//Created by zhbinary on 2019-01-21.
//Email: zhbinary@gmail.com
package types

type ByteBuf interface {
	Capacity() int
	MaxCapacity() int
	ReaderIndex() int
	SetReaderIndex(readerIndex int) (err error)
	WriterIndex() int
	SetWriterIndex(writerIndex int) (err error)
	ReadableBytes() int
	WritableBytes() int
	WritableArray() []byte
	ReadableArray() []byte
	HasArray() bool
	Array() []byte
	MaxWritableBytes() int
	IsReadable() bool
	IsWritable() bool
	Clear()
	MarkReaderIndex()
	ResetReaderIndex()
	MarkWriterIndex()
	ResetWriterIndex()
	DiscardReadBytes() (err error)
	EnsureWritable(minWritableBytes int) (err error)
	Duplicate() ByteBuf
	String() string

	GetBool(index int) (v bool, err error)
	GetUint8(index int) (v uint8, err error)
	GetInt8(index int) (v int8, err error)
	GetUint16(index int) (v uint16, err error)
	GetInt16(index int) (v int16, err error)
	GetUint32(index int) (v uint32, err error)
	GetInt32(index int) (v int32, err error)
	GetUint64(index int) (v uint64, err error)
	GetInt64(index int) (v int64, err error)
	GetFloat32(index int) (v float32, err error)
	GetFloat64(index int) (v float64, err error)
	GetBytes(index int, dst []byte, dstIndex int, length int) (err error)
	GetBytes0(index int, dst ByteBuf, dstIndex int, length int) (err error)

	SetBool(index int, value bool) (err error)
	SetUint8(index int, value uint8) (err error)
	SetInt8(index int, value int8) (err error)
	SetUint16(index int, value uint16) (err error)
	SetInt16(index int, value int16) (err error)
	SetUint32(index int, value uint32) (err error)
	SetInt32(index int, value int32) (err error)
	SetUint64(index int, value uint64) (err error)
	SetInt64(index int, value int64) (err error)
	SetFloat32(index int, value float32) (err error)
	SetFloat64(index int, value float64) (err error)
	SetBytes(index int, src []byte, srcIndex int, length int) (err error)
	SetBytes0(index int, src ByteBuf, srcIndex int, length int) (err error)
	SetZero(index int, length int) (err error)

	ReadBool() (v bool, err error)
	ReadUint8() (v uint8, err error)
	ReadInt8() (v int8, err error)
	ReadUint16() (v uint16, err error)
	ReadInt16() (v int16, err error)
	ReadUint32() (v uint32, err error)
	ReadInt32() (v int32, err error)
	ReadUint64() (v uint64, err error)
	ReadInt64() (v int64, err error)
	ReadFloat32() (v float32, err error)
	ReadFloat64() (v float64, err error)
	ReadBytes(dst []byte, dstIndex int, length int) (err error)
	ReadBytes0(dst ByteBuf, dstIndex int, length int) (err error)
	SkipBytes(length int) (err error)

	WriteBool(value bool) (err error)
	WriteUint8(value uint8) (err error)
	WriteInt8(value int8) (err error)
	WriteUint16(value uint16) (err error)
	WriteInt16(value int16) (err error)
	WriteUint32(value uint32) (err error)
	WriteInt32(value int32) (err error)
	WriteUint64(value uint64) (err error)
	WriteInt64(value int64) (err error)
	WriteFloat32(value float32) (err error)
	WriteFloat64(value float64) (err error)
	WriteBytes(src []byte, srcIndex int, length int) (err error)
	WriteBytes0(src ByteBuf, srcIndex int, length int) (err error)
	WriteZero(length int) (err error)
}

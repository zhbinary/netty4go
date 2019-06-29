//Created by zhbinary on 2019-01-24.
//Email: zhbinary@gmail.com
package buffer

import (
	"errors"
	"fmt"
	"github.com/zhbinary/heng/types"
)

type HeapByteBuf struct {
	buf               []byte
	readerIndex       int
	writerIndex       int
	markedReaderIndex int
	markedWriterIndex int
	maxCapacity       int
	allocator         types.ByteBufAllocator
}

func NewHeapBytebuf(capacity int) types.ByteBuf {
	return &HeapByteBuf{buf: make([]byte, capacity), maxCapacity: capacity}
}

func (this *HeapByteBuf) Capacity() int {
	return cap(this.buf)
}

func (this *HeapByteBuf) setCapacity(newCapacity int) (err error) {
	if newCapacity < 0 || newCapacity > this.maxCapacity {
		return errors.New(fmt.Sprintf("newCapacity:%d (expected: 0-%d) ", newCapacity, this.maxCapacity))
	}

	oldCapacity := cap(this.buf)
	if newCapacity > oldCapacity {
		newBuf := make([]byte, newCapacity)
		copy(newBuf, this.buf)
		this.buf = newBuf
	} else {

	}
	return
}

func (this *HeapByteBuf) MaxCapacity() int {
	return this.maxCapacity
}

func (this *HeapByteBuf) ReaderIndex() int {
	return this.readerIndex
}

func (this *HeapByteBuf) SetReaderIndex(readerIndex int) (err error) {
	if readerIndex < 0 || readerIndex > this.writerIndex {
		return errors.New(fmt.Sprintf("readerIndex: %d (expected: 0 <= readerIndex <= writerIndex(%d))", readerIndex, this.writerIndex))
	}
	this.readerIndex = readerIndex
	return
}

func (this *HeapByteBuf) WriterIndex() int {
	return this.writerIndex
}

func (this *HeapByteBuf) SetWriterIndex(writerIndex int) (err error) {
	if writerIndex < 0 || writerIndex > this.Capacity() {
		return errors.New(fmt.Sprintf("writerIndex: %d (expected: readerIndex(%d) <= writerIndex <= capacity(%d))", writerIndex, this.readerIndex, this.Capacity()))
	}
	this.writerIndex = writerIndex
	return
}

func (this *HeapByteBuf) ReadableBytes() int {
	return this.writerIndex - this.readerIndex
}

func (this *HeapByteBuf) WritableBytes() int {
	return this.Capacity() - this.writerIndex
}

func (this *HeapByteBuf) MaxWritableBytes() int {
	return this.maxCapacity - this.writerIndex
}

func (this *HeapByteBuf) IsReadable() bool {
	return this.writerIndex > this.readerIndex
}

func (this *HeapByteBuf) IsWritable() bool {
	return this.Capacity() > this.writerIndex
}

func (this *HeapByteBuf) Clear() {
	this.readerIndex = 0
	this.writerIndex = 0
}

func (this *HeapByteBuf) MarkReaderIndex() {
	this.markedReaderIndex = this.readerIndex
}

func (this *HeapByteBuf) ResetReaderIndex() {
	this.SetReaderIndex(this.markedReaderIndex)
}

func (this *HeapByteBuf) MarkWriterIndex() {
	this.markedWriterIndex = this.writerIndex
}

func (this *HeapByteBuf) ResetWriterIndex() {
	this.SetWriterIndex(this.markedWriterIndex)
}

func (this *HeapByteBuf) DiscardReadBytes() (err error) {
	if this.readerIndex == 0 {
		return
	}

	if this.readerIndex != this.writerIndex {
		err = this.SetBytes(0, this.buf, this.readerIndex, this.writerIndex-this.readerIndex)
		if err != nil {
			return err
		}
		this.adjustMarders(this.readerIndex)
		this.readerIndex = 0
		this.writerIndex -= this.readerIndex
	} else {
		this.adjustMarders(this.readerIndex)
		this.readerIndex = 0
		this.writerIndex = 0
	}
	return
}

func (this *HeapByteBuf) GetBool(index int) (b bool, err error) {
	if err = this.checkReadableBytes(1); err != nil {
		return
	}
	if err = this.checkIndex(index, 1); err != nil {
		return
	}
	return GetBool(this.buf, index), nil
}

func (this *HeapByteBuf) GetUint8(index int) (u uint8, err error) {
	if err = this.checkIndex(index, 1); err != nil {
		return
	}
	return GetUint8(this.buf, index), nil
}

func (this *HeapByteBuf) GetInt8(index int) (i int8, err error) {
	if err = this.checkIndex(index, 1); err != nil {
		return
	}
	return GetInt8(this.buf, index), nil
}

func (this *HeapByteBuf) GetUint16(index int) (u uint16, err error) {
	if err = this.checkIndex(index, 2); err != nil {
		return
	}
	return GetUint16(this.buf, index), nil
}

func (this *HeapByteBuf) GetInt16(index int) (i int16, err error) {
	if err = this.checkIndex(index, 2); err != nil {
		return
	}
	return GetInt16(this.buf, index), nil
}

func (this *HeapByteBuf) GetUint32(index int) (u uint32, err error) {
	if err = this.checkIndex(index, 4); err != nil {
		return
	}
	return GetUint32(this.buf, index), nil
}

func (this *HeapByteBuf) GetInt32(index int) (i int32, err error) {
	if err = this.checkIndex(index, 4); err != nil {
		return
	}
	return GetInt32(this.buf, index), nil
}

func (this *HeapByteBuf) GetUint64(index int) (v uint64, err error) {
	if err = this.checkIndex(index, 8); err != nil {
		return
	}
	return GetUint64(this.buf, index), nil
}

func (this *HeapByteBuf) GetInt64(index int) (i int64, err error) {
	if err = this.checkIndex(index, 8); err != nil {
		return
	}
	return GetInt64(this.buf, index), nil
}

func (this *HeapByteBuf) GetFloat32(index int) (f float32, err error) {
	if err = this.checkIndex(index, 4); err != nil {
		return
	}
	return GetFloat32(this.buf, index), nil
}

func (this *HeapByteBuf) GetFloat64(index int) (f float64, err error) {
	if err = this.checkIndex(index, 8); err != nil {
		return
	}
	return GetFloat64(this.buf, index), nil
}

func (this *HeapByteBuf) GetBytes(index int, dst []byte, dstIndex int, length int) (err error) {
	if err = this.checkDstIndex(index, length, dstIndex, len(dst)); err != nil {
		return
	}
	copy(dst[dstIndex:], this.buf[index:index+length])
	return
}

func (this *HeapByteBuf) GetBytes0(index int, dst types.ByteBuf, dstIndex int, length int) (err error) {
	if err = this.checkDstIndex(index, length, dstIndex, dst.Capacity()); err != nil {
		return
	}
	copy(dst.Array()[dstIndex:], this.buf[index:index+length])
	return
}

func (this *HeapByteBuf) SetBool(index int, value bool) (err error) {
	if err = this.checkIndex(index, 1); err != nil {
		return
	}
	SetBool(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetUint8(index int, value uint8) (err error) {
	if err = this.checkIndex(index, 1); err != nil {
		return
	}
	SetUint8(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetInt8(index int, value int8) (err error) {
	if err = this.checkIndex(index, 1); err != nil {
		return
	}
	SetInt8(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetUint16(index int, value uint16) (err error) {
	if err = this.checkIndex(index, 2); err != nil {
		return
	}
	SetUint16(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetInt16(index int, value int16) (err error) {
	if err = this.checkIndex(index, 2); err != nil {
		return
	}
	SetInt16(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetUint32(index int, value uint32) (err error) {
	if err = this.checkIndex(index, 4); err != nil {
		return
	}
	SetUint32(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetInt32(index int, value int32) (err error) {
	if err = this.checkIndex(index, 4); err != nil {
		return
	}
	SetInt32(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetUint64(index int, value uint64) (err error) {
	if err = this.checkIndex(index, 8); err != nil {
		return
	}
	SetUint64(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetInt64(index int, value int64) (err error) {
	if err = this.checkIndex(index, 8); err != nil {
		return
	}
	SetInt64(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetFloat32(index int, value float32) (err error) {
	if err = this.checkIndex(index, 4); err != nil {
		return
	}
	SetFloat32(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetFloat64(index int, value float64) (err error) {
	if err = this.checkIndex(index, 8); err != nil {
		return
	}
	SetFloat64(this.buf, index, value)
	return
}

func (this *HeapByteBuf) SetBytes(index int, src []byte, srcIndex int, length int) (err error) {
	if err = this.checkSrcIndex(index, length, srcIndex, len(src)); err != nil {
		return
	}
	copy(this.buf[index:], src[srcIndex:srcIndex+length])
	return
}

func (this *HeapByteBuf) SetBytes0(index int, src types.ByteBuf, srcIndex int, length int) (err error) {
	if err = this.checkSrcIndex(index, length, srcIndex, src.Capacity()); err != nil {
		return
	}
	copy(this.buf[index:], src.Array()[srcIndex:srcIndex+length])
	return
}

// todo zero
func (this *HeapByteBuf) SetZero(index int, length int) (err error) {
	if length == 0 {
		//return
	}
	if err = this.checkIndex(index, length); err != nil {
		return
	}
	SetUint8(this.buf, index, 0)
	return
}

func (this *HeapByteBuf) ReadBool() (v bool, err error) {
	if err = this.checkReadableBytes(1); err != nil {
		return
	}
	v = GetBool(this.buf, this.readerIndex)
	this.readerIndex++
	return
}

func (this *HeapByteBuf) ReadUint8() (v uint8, err error) {
	if err = this.checkReadableBytes(1); err != nil {
		return
	}
	v = GetUint8(this.buf, this.readerIndex)
	this.readerIndex += 1
	return
}

func (this *HeapByteBuf) ReadInt8() (v int8, err error) {
	if err = this.checkReadableBytes(1); err != nil {
		return
	}
	v = GetInt8(this.buf, this.readerIndex)
	this.readerIndex += 1
	return
}

func (this *HeapByteBuf) ReadUint16() (v uint16, err error) {
	if err = this.checkReadableBytes(2); err != nil {
		return
	}
	v = GetUint16(this.buf, this.readerIndex)
	this.readerIndex += 2
	return
}

func (this *HeapByteBuf) ReadInt16() (v int16, err error) {
	if err = this.checkReadableBytes(2); err != nil {
		return
	}
	v = GetInt16(this.buf, this.readerIndex)
	this.readerIndex += 2
	return
}

func (this *HeapByteBuf) ReadUint32() (v uint32, err error) {
	if err = this.checkReadableBytes(4); err != nil {
		return
	}
	v = GetUint32(this.buf, this.readerIndex)
	this.readerIndex += 4
	return
}

func (this *HeapByteBuf) ReadInt32() (v int32, err error) {
	if err = this.checkReadableBytes(4); err != nil {
		return
	}
	v = GetInt32(this.buf, this.readerIndex)
	this.readerIndex += 4
	return
}

func (this *HeapByteBuf) ReadUint64() (v uint64, err error) {
	if err = this.checkReadableBytes(8); err != nil {
		return
	}
	v = GetUint64(this.buf, this.readerIndex)
	this.readerIndex += 8
	return
}

func (this *HeapByteBuf) ReadInt64() (v int64, err error) {
	if err = this.checkReadableBytes(8); err != nil {
		return
	}
	v = GetInt64(this.buf, this.readerIndex)
	this.readerIndex += 8
	return
}

func (this *HeapByteBuf) ReadFloat32() (v float32, err error) {
	if err = this.checkReadableBytes(4); err != nil {
		return
	}
	v = GetFloat32(this.buf, this.readerIndex)
	this.readerIndex += 4
	return
}

func (this *HeapByteBuf) ReadFloat64() (v float64, err error) {
	if err = this.checkReadableBytes(8); err != nil {
		return
	}
	v = GetFloat64(this.buf, this.readerIndex)
	this.readerIndex += 8
	return
}

func (this *HeapByteBuf) ReadBytes(dst []byte, dstIndex int, length int) (err error) {
	if err = this.checkReadableBytes(length); err != nil {
		return
	}

	if err = this.GetBytes(this.readerIndex, dst, dstIndex, length); err != nil {
		return
	}

	this.readerIndex += length
	return
}

func (this *HeapByteBuf) ReadBytes0(dst types.ByteBuf, dstIndex int, length int) (err error) {
	if err = this.checkReadableBytes(length); err != nil {
		return
	}

	if err = this.GetBytes0(this.readerIndex, dst, dstIndex, length); err != nil {
		return
	}

	this.readerIndex += length
	return
}

func (this *HeapByteBuf) SkipBytes(length int) (err error) {
	if err = this.checkReadableBytes(length); err != nil {
		return
	}
	this.readerIndex += length
	return
}

func (this *HeapByteBuf) WriteBool(value bool) (err error) {
	if err = this.EnsureWritable(1); err != nil {
		return
	}
	SetBool(this.buf, this.writerIndex, value)
	this.writerIndex++
	return
}

func (this *HeapByteBuf) WriteUint8(value uint8) (err error) {
	if err = this.EnsureWritable(1); err != nil {
		return
	}
	SetUint8(this.buf, this.writerIndex, value)
	this.writerIndex++
	return
}

func (this *HeapByteBuf) WriteInt8(value int8) (err error) {
	if err = this.EnsureWritable(1); err != nil {
		return
	}
	SetInt8(this.buf, this.writerIndex, value)
	this.writerIndex++
	return
}

func (this *HeapByteBuf) WriteUint16(value uint16) (err error) {
	if err = this.EnsureWritable(2); err != nil {
		return
	}
	SetUint16(this.buf, this.writerIndex, value)
	this.writerIndex += 2
	return
}

func (this *HeapByteBuf) WriteInt16(value int16) (err error) {
	if err = this.EnsureWritable(2); err != nil {
		return
	}
	SetInt16(this.buf, this.writerIndex, value)
	this.writerIndex += 2
	return
}

func (this *HeapByteBuf) WriteUint32(value uint32) (err error) {
	if err = this.EnsureWritable(4); err != nil {
		return
	}
	SetUint32(this.buf, this.writerIndex, value)
	this.writerIndex += 4
	return
}

func (this *HeapByteBuf) WriteInt32(value int32) (err error) {
	if err = this.EnsureWritable(4); err != nil {
		return
	}
	SetInt32(this.buf, this.writerIndex, value)
	this.writerIndex += 4
	return
}

func (this *HeapByteBuf) WriteUint64(value uint64) (err error) {
	if err = this.EnsureWritable(8); err != nil {
		return
	}
	SetUint64(this.buf, this.writerIndex, value)
	this.writerIndex += 8
	return
}

func (this *HeapByteBuf) WriteInt64(value int64) (err error) {
	if err = this.EnsureWritable(8); err != nil {
		return
	}
	SetInt64(this.buf, this.writerIndex, value)
	this.writerIndex += 8
	return
}

func (this *HeapByteBuf) WriteFloat32(value float32) (err error) {
	if err = this.EnsureWritable(4); err != nil {
		return
	}
	SetFloat32(this.buf, this.writerIndex, value)
	this.writerIndex += 4
	return
}

func (this *HeapByteBuf) WriteFloat64(value float64) (err error) {
	if err = this.EnsureWritable(8); err != nil {
		return
	}
	SetFloat64(this.buf, this.writerIndex, value)
	this.writerIndex += 8
	return
}

func (this *HeapByteBuf) WriteBytes(src []byte, srcIndex int, length int) (err error) {
	if err = this.EnsureWritable(length); err != nil {
		return
	}
	if err = this.SetBytes(this.writerIndex, src, srcIndex, length); err != nil {
		return
	}
	this.writerIndex += length
	return
}

func (this *HeapByteBuf) WriteBytes0(src types.ByteBuf, srcIndex int, length int) (err error) {
	if err = this.EnsureWritable(length); err != nil {
		return
	}
	if err = this.SetBytes0(this.writerIndex, src, srcIndex, length); err != nil {
		return
	}
	this.writerIndex += length
	return
}

// todo Write zero
func (this *HeapByteBuf) WriteZero(length int) (err error) {
	panic("implement me")
}

func (this *HeapByteBuf) checkSrcIndex(index int, length int, srcIndex int, srcCapacity int) (err error) {
	if err = this.checkIndex(index, length); err != nil {
		return
	}
	if this.isOutOfBounds(srcIndex, length, srcCapacity) {
		err = errors.New(fmt.Sprintf("srcIndex: %d, length: %d (expected: range(0, %d))", srcIndex, length, srcCapacity))
	}
	return
}

func (this *HeapByteBuf) checkDstIndex(index int, length int, dstIndex int, dstCapacity int) (err error) {
	if err = this.checkIndex(index, length); err != nil {
		return
	}
	if this.isOutOfBounds(dstIndex, length, dstCapacity) {
		err = errors.New(fmt.Sprintf("dstIndex: %d, length: %d (expected: range(0, %d))", dstIndex, length, dstCapacity))
	}
	return
}

func (this *HeapByteBuf) checkIndex(index int, length int) (err error) {
	if this.isOutOfBounds(index, length, this.Capacity()) {
		return errors.New(fmt.Sprintf("index: %d, length: %d (expected: range(0, %d))", index, length, this.Capacity()))
	}
	return
}

func (this *HeapByteBuf) isOutOfBounds(index int, length int, capacity int) bool {
	return (index | length | (capacity - (index + length))) < 0
}

func (this *HeapByteBuf) adjustMarders(decrement int) {
	if this.markedReaderIndex <= decrement {
		this.markedReaderIndex = 0
		if this.markedWriterIndex <= decrement {
			this.markedWriterIndex = 0
		} else {
			this.markedWriterIndex -= decrement
		}
	} else {
		this.markedReaderIndex -= decrement
		this.markedWriterIndex -= this.markedWriterIndex
	}
}

func (this *HeapByteBuf) checkReadableBytes(minimumReadableBytes int) (err error) {
	if minimumReadableBytes < 0 {
		return errors.New(fmt.Sprintf("minimumReadableBytes: %d (expected: >= 0)", minimumReadableBytes))
	}
	if this.ReadableBytes() < minimumReadableBytes {
		return errors.New(fmt.Sprintf("readerIndex(%d) + length(%d) exceeds writerIndex(%d)", this.readerIndex, minimumReadableBytes, this.writerIndex))
	}
	return
}

func (this *HeapByteBuf) EnsureWritable(minWritableBytes int) (err error) {
	if minWritableBytes < 0 {
		return errors.New(fmt.Sprintf("minWritableBytes: %d (expected: >= 0)", minWritableBytes))
	}

	if minWritableBytes > this.maxCapacity-this.writerIndex {
		return errors.New(fmt.Sprintf("writerIndex(%d) + minWritableBytes(%d) exceeds maxCapacity(%d)", this.writerIndex, minWritableBytes, this.maxCapacity))
	}

	newCapacity := this.Capacity() * 2
	this.setCapacity(newCapacity)
	return
}

func (this *HeapByteBuf) alloc() types.ByteBufAllocator {
	return this.allocator
}

func (this *HeapByteBuf) WritableArray() []byte {
	if this.writerIndex == this.Capacity() {
		return nil
	}
	return this.buf[this.writerIndex:]
}
func (this *HeapByteBuf) ReadableArray() []byte {
	if this.readerIndex == this.writerIndex {
		return nil
	}
	return this.buf[this.readerIndex:this.writerIndex]
}

func (this *HeapByteBuf) Array() []byte {
	return this.buf
}

func (this *HeapByteBuf) HasArray() bool {
	return true
}

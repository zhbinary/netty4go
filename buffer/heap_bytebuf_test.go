//Created by zhbinary on 2019-11-09.
package buffer

import (
	"fmt"
	"testing"
)

func TestNewHeapBytebuf(t *testing.T) {
	size := 100
	buf := NewHeapBytebuf(size)

	buf.WriteUint8(byte(2))
	v, _ := buf.ReadUint8()
	if buf.Capacity() != size {
		t.Error("")
	}
	if buf.WriterIndex() != 1 {
		t.Error()
	}
	if buf.ReaderIndex() != 1 {
		t.Error()
	}
	if v != 2 {
		t.Error()
	}

	buf.WriteInt8(-1)
	x, _ := buf.ReadInt8()
	if buf.WriterIndex() != 2 {
		t.Error()
	}
	if buf.ReaderIndex() != 2 {
		t.Error()
	}
	if x != -1 {
		t.Error()
	}

	buf.WriteInt16(-3)
	i16, _ := buf.ReadInt16()
	if buf.WriterIndex() != 4 {
		t.Error()
	}
	if buf.ReaderIndex() != 4 {
		t.Error()
	}
	if i16 != -3 {
		t.Error()
	}

	buf.WriteUint16(77)
	u16, _ := buf.ReadUint16()
	if buf.WriterIndex() != 6 {
		t.Error()
	}
	if buf.ReaderIndex() != 6 {
		t.Error()
	}
	if u16 != 77 {
		t.Error()
	}

	buf.WriteInt32(-257)
	i32, _ := buf.ReadInt32()
	if buf.WriterIndex() != 10 {
		t.Error()
	}
	if buf.ReaderIndex() != 10 {
		t.Error()
	}
	if i32 != -257 {
		t.Error()
	}

	buf.WriteUint32(13)
	u32, _ := buf.ReadUint32()
	if buf.WriterIndex() != 14 {
		t.Error()
	}
	if buf.ReaderIndex() != 14 {
		t.Error()
	}
	if u32 != 13 {
		t.Error()
	}

	buf.WriteFloat32(53.76)
	f32, _ := buf.ReadFloat32()
	if buf.WriterIndex() != 18 {
		t.Error()
	}
	if buf.ReaderIndex() != 18 {
		t.Error()
	}
	if f32 != 53.76 {
		t.Error()
	}

	buf.WriteUint64(565656)
	u64, _ := buf.ReadUint64()
	if buf.WriterIndex() != 26 {
		t.Error()
	}
	if buf.ReaderIndex() != 26 {
		t.Error()
	}
	if u64 != 565656 {
		t.Error()
	}

	buf.WriteInt64(-9999)
	i64, _ := buf.ReadInt64()
	if buf.WriterIndex() != 34 {
		t.Error()
	}
	if buf.ReaderIndex() != 34 {
		t.Error()
	}
	if i64 != -9999 {
		t.Error()
	}

	buf.WriteFloat64(-33.77)
	f64, _ := buf.ReadFloat64()
	if buf.WriterIndex() != 42 {
		t.Error()
	}
	if buf.ReaderIndex() != 42 {
		t.Error()
	}
	if f64 != -33.77 {
		t.Error()
	}

	buf.WriteBool(false)
	b, _ := buf.ReadBool()
	if buf.WriterIndex() != 43 {
		t.Error()
	}
	if buf.ReaderIndex() != 43 {
		t.Error()
	}
	if b != false {
		t.Error()
	}

	bytes := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	buf.WriteBytes(bytes, 0, len(bytes))
	buffer := make([]byte, 10)
	buf.ReadBytes(buffer, 0, 10)
	if buf.WriterIndex() != 53 {
		t.Error()
	}
	if buf.ReaderIndex() != 53 {
		t.Error()
	}
	for n, b := range buffer {
		if int(b) != n {
			t.Error()
		}
	}

	buf.WriteBytes(bytes, 0, 7)
	if buf.WriterIndex() != 60 {
		t.Error()
	}

	bytes = make([]byte, 40)
	err := buf.WriteBytes(bytes, 0, len(bytes))
	if err != nil {
		t.Error(err)
	}
	if buf.WriterIndex() != 100 {
		t.Error()
	}

	err = buf.WriteBool(false)
	if err == nil {
		t.Error()
	} else {
		fmt.Println(err)
	}
}

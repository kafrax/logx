package logx

import (
	"sync"
	"io"
	"unsafe"
)
//pool 4 bytesbuffer
type bytesBuffer struct {
	B []byte
}

//no bb
var bbFree = sync.Pool{
	New: func() interface{} { return new(bytesBuffer) },
}

func (b *bytesBuffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.B)
	return int64(n), err
}

func (b *bytesBuffer) Bytes() []byte {
	return b.B
}

func (b *bytesBuffer) Write(p []byte) (int, error) {
	b.B = append(b.B, p...)
	return len(p), nil
}

func (b *bytesBuffer) WriteByte(c byte) error {
	b.B = append(b.B, c)
	return nil
}

func (b *bytesBuffer) WriteString(s string) (int, error) {
	b.B = append(b.B, s...)
	return len(s), nil
}

func (b *bytesBuffer) Set(p []byte) {
	b.B = append(b.B[:0], p...)
}

func (b *bytesBuffer) SetString(s string) {
	b.B = append(b.B[:0], s...)
}

func (b *bytesBuffer) String() string {
	return b2s(b.B)
}

func (b *bytesBuffer) Reset() {
	b.B = b.B[:0]
}

func s2b(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

package logx

import (
	"bytes"
	"io"
	"sync"
	"unsafe"
)

//pool for logMessage logMsg
type logMsg struct { B []byte }

//no bb
var logMsgFree = sync.Pool{
	New: func() interface{} { return new(logMsg) },
}

func (b *logMsg) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.B)
	return int64(n), err
}

func (b *logMsg) Bytes() []byte {
	return b.B
}

func (b *logMsg) Write(p []byte) (int, error) {
	b.B = append(b.B, p...)
	return len(p), nil
}

func (b *logMsg) WriteByte(c byte) error {
	b.B = append(b.B, c)
	return nil
}

func (b *logMsg) WriteString(s string) (int, error) {
	b.B = append(b.B, s...)
	return len(s), nil
}

func (b *logMsg) Set(p []byte) {
	b.B = append(b.B[:0], p...)
}

func (b *logMsg) SetString(s string) {
	b.B = append(b.B[:0], s...)
}

func (b *logMsg) String() string {
	return b2s(b.B)
}

func (b *logMsg) Reset() {
	b.B = b.B[:0]
}

func s2b(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//pool for stitching log data¬
var bpFree = sync.Pool{}

func bufferPoolGet() *bytes.Buffer {
	if buf := bpFree.Get(); buf != nil {
		return buf.(*bytes.Buffer)
	} else {
		return &bytes.Buffer{}
	}
}

func put(b *bytes.Buffer) { bpFree.Put(b) }

func bufferPoolFree(b *bytes.Buffer) {
	b.Reset()
	put(b)
}

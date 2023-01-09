package helper

import (
	"io"
)

// Rot13Reader ...
type Rot13Reader struct {
	r io.Reader
}

// NewRot13Reader Rot13Reader
func NewRot13Reader(s io.Reader) *Rot13Reader {
	return &Rot13Reader{s}
}

// 转换byte  前进13位/后退13位
func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b = b + 13
	case 'M' < b && b <= 'Z':
		b = b - 13
	case 'a' <= b && b <= 'm':
		b = b + 13
	case 'm' < b && b <= 'z':
		b = b - 13
	}
	return b
}

// 重写Read方法
func (mr Rot13Reader) Read(b []byte) (int, error) {
	n, e := mr.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, e
}

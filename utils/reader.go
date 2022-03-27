package utils

import (
	"bytes"
	"encoding/binary"
	"io"
	"wasm-jvm/logger"
)

type Reader struct {
	Src io.Reader
}

func CreateReader(src io.Reader) *Reader {
	return &Reader{src}
}

func CreateBytesReader(src []byte) *Reader {
	return &Reader{bytes.NewReader(src)}
}

// From https://zserge.com/posts/jvm/

func (r *Reader) U1() uint8  { return r.ReadBytes(1)[0] }
func (r *Reader) U2() uint16 { return binary.BigEndian.Uint16(r.ReadBytes(2)) }
func (r *Reader) U4() uint32 { return binary.BigEndian.Uint32(r.ReadBytes(4)) }
func (r *Reader) U8() uint64 { return binary.BigEndian.Uint64(r.ReadBytes(8)) }

func (r *Reader) ReadBytes(n int) []byte {
	bs := make([]byte, n)
	if _, err := io.ReadFull(r.Src, bs); err != nil {
		logger.Errorln("unexpected EOF: ", err.Error())
	}
	//logger.Infoln(n, "bytes >> ", bs)
	return bs
}

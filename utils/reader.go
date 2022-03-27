package utils

import (
	"encoding/binary"
	"io"
	"wasm-jvm/logger"
)

type Reader struct {
	Origin io.Reader
}

func CreateReader() *Reader {
	return &Reader{}
}

// From https://zserge.com/posts/jvm/
func (r *Reader) u1() uint8  { return r.readBytes(1)[0] }
func (r *Reader) u2() uint16 { return binary.BigEndian.Uint16(r.readBytes(2)) }
func (r *Reader) u4() uint32 { return binary.BigEndian.Uint32(r.readBytes(4)) }
func (r *Reader) u8() uint64 { return binary.BigEndian.Uint64(r.readBytes(8)) }

func (r *Reader) readBytes(n int) []byte {
	bs := make([]byte, n)
	if _, err := io.ReadFull(r.Origin, bs); err != nil {
		logger.Errorln("unexpected EOF: ", err.Error())
	}
	//logger.Infoln(n, "bytes >> ", bs)
	return bs
}

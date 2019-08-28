package bitcask

import (
	"encoding/binary"
	"github.com/shniu/gostuff/kvs/util"
)

const (
	// Hint Header Size
	// ts : flag : ksz : vsz : valueOffset : key
	//  4     2     4     4        8
	HintHeaderSize = 22
)

type HintHeader struct {
	ts      uint32
	flag    uint16
	ksz     uint32
	vsz     uint32
	vOffset uint64
}

type Hint struct {
	header *HintHeader
	key    []byte
}

func (hint *Hint) bytes() []byte {
	hb := EncodeHintHeader(hint.header)
	return util.BytesCombine(hb, hint.key)
}

func DecodeHintHeader(headerBytes []byte) *HintHeader {
	hh := &HintHeader{}
	hh.ts = binary.LittleEndian.Uint32(headerBytes[:4])
	hh.flag = binary.LittleEndian.Uint16(headerBytes[4:6])
	hh.ksz = binary.LittleEndian.Uint32(headerBytes[6:10])
	hh.vsz = binary.LittleEndian.Uint32(headerBytes[10:14])
	hh.vOffset = binary.LittleEndian.Uint64(headerBytes[14:])
	return hh
}

func EncodeHintHeader(hh *HintHeader) []byte {
	encodeByte := make([]byte, HintHeaderSize)
	binary.LittleEndian.PutUint32(encodeByte[:4], hh.ts)
	binary.LittleEndian.PutUint16(encodeByte[4:6], hh.flag)
	binary.LittleEndian.PutUint32(encodeByte[6:10], hh.ksz)
	binary.LittleEndian.PutUint32(encodeByte[10:14], hh.vsz)
	binary.LittleEndian.PutUint64(encodeByte[14:], hh.vOffset)
	return encodeByte
}

func (hh *HintHeader) deleted() bool {
	return hh.flag == 1
}

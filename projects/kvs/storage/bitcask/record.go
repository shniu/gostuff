package bitcask

import (
	"encoding/binary"
	"github.com/shniu/gostuff/projects/kvs/util"
)

const (
	RecordDeleteFlag = 1 << iota

	// Record Header Size
	// crc32 : ts : flag : ksz : vsz : key : value
	//   4      4     2     4     4     xx    xxx
	RecHeaderSize = 18
)

// Record in the data file
type Record struct {
	// crc32
	crc uint32
	// Timestamp, accurate to second, e.g. time.Now().Unix()
	ts uint32
	// flag, whether deleted or not: 0 not delete, 1 delete
	flag uint16
	// Key size
	ksz uint32
	// Value size
	vsz uint32
	// Key Bytes
	key []byte
	// Value Bytes
	value []byte

	header []byte
}

func (r *Record) Prepare() {
	header := r.encodeHeader()

	// Set header
	r.header = header
}

func (r *Record) encodeHeader() []byte {
	var header [RecHeaderSize]byte
	binary.LittleEndian.PutUint32(header[4:8], r.ts)
	binary.LittleEndian.PutUint16(header[8:10], r.flag)
	binary.LittleEndian.PutUint32(header[10:14], r.ksz)
	binary.LittleEndian.PutUint32(header[14:18], r.vsz)

	// Get crc
	data := util.BytesCombine(header[4:], r.key, r.value)
	r.crc = util.Crc(data)
	binary.LittleEndian.PutUint32(header[:4], r.crc)
	return header[:]
}

func (r *Record) decodeHeader(headers []byte) {
	r.crc = binary.LittleEndian.Uint32(headers[:4])
	r.ts = binary.LittleEndian.Uint32(headers[4:8])
	r.flag = binary.LittleEndian.Uint16(headers[8:10])
	r.ksz = binary.LittleEndian.Uint32(headers[10:14])
	r.vsz = binary.LittleEndian.Uint32(headers[14:18])
}

func (r *Record) kvSize() int {
	return int(r.ksz + r.vsz)
}

// Record total size
func (r *Record) Size() int {
	//sizeofCrc := unsafe.Sizeof(r.crc)
	//sizeofTs := unsafe.Sizeof(r.ts)
	//sizeofKsz := unsafe.Sizeof(r.ksz)
	//sizeofVsz := unsafe.Sizeof(r.vsz)
	return RecHeaderSize + len(r.key) + len(r.value)
}

// Get Bytes
func (r *Record) Bytes() []byte {
	return util.BytesCombine(r.header, r.key, r.value)
}

func (r *Record) DecodeBody(body []byte) {
	r.key = body[:r.ksz]
	r.value = body[r.ksz:]
}

func (r *Record) CheckCrc() bool {
	header := r.encodeHeader()
	crc2 := binary.LittleEndian.Uint32(header[:4])
	return r.crc == crc2
}

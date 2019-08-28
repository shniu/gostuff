package bitcask

import (
	"github.com/shniu/gostuff/kvs/util"
	"os"
	"sync"
	"time"
)

// Segment file
// The log-structured data file is divided into several segments with controllable size
type SegmentFile struct {
	// fp is a pointer file to a segment file
	fp *os.File

	// the segment file's id
	fileId      uint32
	writeOffset uint64

	// hint file pointer
	hintFp *os.File

	// is it an active file? there is only one active file in the system
	active bool
}

func (sf *SegmentFile) write(key string, value []byte, delete bool) (entry, error) {
	if !sf.active {
		return entry{}, nil
	}

	// Write to data file
	var keyBytes = []byte(key)
	rec := &Record{
		ts:    uint32(time.Now().Unix()),
		flag:  0,
		ksz:   uint32(len(keyBytes)),
		vsz:   uint32(len(value)),
		key:   keyBytes,
		value: value,
	}
	if delete {
		rec.flag |= RecordDeleteFlag
	}
	rec.Prepare()
	vOffset := sf.writeOffset + uint64(RecHeaderSize+rec.ksz)

	byteswrite, err := util.AppendWrite(sf.fp, rec.Bytes())
	// byteswrite, err := sf.fp.Write(rec.Bytes())
	if err != nil || byteswrite != rec.Size() {
		panic(err)
	}
	logger.Printf("The active file append write Record, Record size = %v, Record key = %s",
		rec.Size(), key)

	// Write to hint file
	hh := &HintHeader{
		ts:      rec.ts,
		flag:    rec.flag,
		ksz:     rec.ksz,
		vsz:     rec.vsz,
		vOffset: vOffset,
	}
	hint := &Hint{
		header: hh,
		key:    keyBytes,
	}
	byteswrite, err = util.AppendWrite(sf.hintFp, hint.bytes())
	// byteswrite, err = sf.hintFp.Write(hint.Bytes())
	if err != nil {
		panic(err)
	}
	logger.Printf("The hint file append write Record, hint size = %v, Record key = %s",
		byteswrite, key)

	sf.writeOffset += uint64(rec.Size())

	return entry{
		fileId:  sf.fileId,
		vsz:     rec.vsz,
		vOffset: vOffset,
		ts:      rec.ts,
		flag:    rec.flag,
	}, nil
}

func (sf *SegmentFile) Close() {
	sf.fp.Close()
	if sf.hintFp != nil {
		sf.hintFp.Close()
	}
}

func openSFile(filePath string, fileId uint32) (*SegmentFile, error) {
	fp, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &SegmentFile{
		fileId:      fileId,
		fp:          fp,
		hintFp:      nil,
		writeOffset: 0,
		active:      false,
	}, nil
}

// A set of segment files
type SegmentFileSet struct {
	// segment files
	segmentFileMap map[uint32]*SegmentFile

	// rw lock
	rwLock *sync.RWMutex
}

func (sfs *SegmentFileSet) add(activeSf *SegmentFile) {
	rwLock.Lock()
	defer rwLock.Unlock()

	sfs.segmentFileMap[activeSf.fileId] = activeSf
}

func (sfs *SegmentFileSet) Close() {
	rwLock.Lock()
	defer rwLock.Unlock()
	for _, sf := range sfs.segmentFileMap {
		sf.fp.Close()
		sf.hintFp.Close()
	}
}

func (sfs *SegmentFileSet) get(fileId uint32) *SegmentFile {
	sfs.rwLock.RLock()
	defer sfs.rwLock.RUnlock()
	sf := sfs.segmentFileMap[fileId]
	return sf
}

func (sfs *SegmentFileSet) put(sf *SegmentFile, fileId uint32) {
	sfs.rwLock.Lock()
	defer sfs.rwLock.Unlock()
	sfs.segmentFileMap[fileId] = sf
}

// New a segment file
func newSegmentFile(active bool) *SegmentFile {
	return &SegmentFile{
		active: active,
	}
}

// Init old file set
func initOldFileSet() *SegmentFileSet {
	return &SegmentFileSet{
		segmentFileMap: make(map[uint32]*SegmentFile),
		rwLock:         &sync.RWMutex{},
	}
}

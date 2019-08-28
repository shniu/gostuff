package bitcask

// The memory index of a value
// There is a key corresponding to it in memory, total 20 Bytes
type entry struct {
	// segment file id
	fileId uint32

	// The value size
	vsz uint32

	// The value position in the data file, also is valuePos
	vOffset uint64

	// Timestamp
	ts uint32

	// flag
	flag uint16
}

func NewEntry() {

}

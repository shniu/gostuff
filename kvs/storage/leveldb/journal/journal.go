package journal

import "io"

const BlockSize = 32 * 1024

type Dropper interface {
	Drop(err error)
}

//
// Reader
//

// Reader reads journals from an underlying io.Reader.
type Reader struct {
	// the underlying io.Reader
	r io.Reader
	// strict flag
	strict bool
	// checksum fla
	checksum bool
}

func NewReader(r io.Reader, dropper Dropper, strict, checksum bool) *Reader {
	return &Reader{
		r: r,
	}
}

//
// Writer
//

type Writer struct {
	// the underlying io.Writer
	w io.Writer
}

// Returns a new writer
func NewWriter(w io.Writer) *Writer {
	return &Writer{}
}

func (w *Writer) Next() (io.Writer, error) {
	return &singleWriter{w: w, seq: 0}, nil
}

type singleWriter struct {
	w   *Writer
	seq int
}

func (x singleWriter) Write(p []byte) (int, error) {
	return 0, nil
}

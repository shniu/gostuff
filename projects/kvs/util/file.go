package util

import (
	"github.com/shniu/gostuff/projects/kvs/options"
	"os"
	"strconv"
	"strings"
)

// GetLastHintFile
func GetLastHintFile(hintFps []*os.File, defaultLastFunc func(int) string) (uint32, *os.File) {

	lastId := -1
	var lastHintFp *os.File

	for _, fp := range hintFps {
		filename := fp.Name()
		startPos := strings.LastIndex(filename, "/") + 1
		endPos := strings.LastIndex(filename, ".hint")
		idx, _ := strconv.Atoi(filename[startPos:endPos])
		if lastId < idx {
			lastId = idx
			lastHintFp = fp
		}
	}

	// when lt 0, it means that's a newly created dir
	if lastId < 0 {
		lastId = 0
		lastHintFp, _ = os.OpenFile(defaultLastFunc(lastId), os.O_CREATE|os.O_RDONLY, os.ModePerm)
	}

	return uint32(lastId), lastHintFp
}

func WritePid(f *os.File, fid uint32) {
	pid := strings.Join([]string{
		strconv.Itoa(os.Getpid()),
		"\t",
		strconv.Itoa(int(fid)),
		options.Options.DataFileSuffix,
	}, "")

	_, _ = f.WriteAt([]byte(pid), 0)
}

func AppendWrite(f *os.File, data []byte) (int, error) {
	stat, err := f.Stat()
	if err != nil {
		return -1, err
	}

	return f.WriteAt(data, stat.Size())
}

func LockFileWhenStartup(lockfpath string) (*os.File, error) {
	return os.OpenFile(lockfpath, os.O_EXCL|os.O_CREATE|os.O_RDWR, os.ModePerm)
}

func CloseHintFp(fps []*os.File) {
	for _, fp := range fps {
		fp.Close()
	}
}

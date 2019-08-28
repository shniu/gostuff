package bitcask

// An empty interface may hold values of any type: https://tour.golang.org/methods/14
// defer and sync.Mutex: https://segmentfault.com/a/1190000006823652
// Reader and Writer: http://www.flysnow.org/2017/05/08/go-in-action-go-reader-writer.html
// Reading files in Go — an overview: https://kgrz.io/reading-files-in-go-an-overview.html
// https://varunpant.com/posts/reading-and-writing-binary-files-in-go-lang

// https://pdfs.semanticscholar.org/fd5e/cfc6dd8f3df358e0dbdef5b03d87742f7c19.pdf

// https://github.com/emluque/dscache
// https://github.com/peterbourgon/diskv
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/SUMMARY.md
// https://github.com/Winnerhust/Code-of-Book/blob/master/Large-Scale-Distributed-Storage-System/bitcask.py

import (
	"errors"
	"github.com/shniu/gostuff/kvs/log"
	"github.com/shniu/gostuff/kvs/options"
	"github.com/shniu/gostuff/kvs/util"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

var logger = log.Logger
var InvalidKeyError = errors.New("invalid key")
var defaultActiveFileMode = os.O_CREATE | os.O_RDWR
var defaultOldFileMode = os.O_RDONLY

//type Kvs interface {
//	Get(key string) []byte
//	Set(key string, value []byte) (bool, error)
//	Delete(key string) (bool, error)
//	Close()
//	Sync() error
//	Keys() []string
//	Merge()
//}

// Open kvs, if not exist, then create it, or else to init
func Open(opts *options.KvsOptions) (kvs *Kvsdb, err error) {
	kvs = new(Kvsdb)
	kvs.Opts = opts
	kvs.rwLock = &sync.RWMutex{}

	// lock file
	lockFilePath := kvs.Opts.GetLockFilePath()
	kvs.lockFp, err = util.LockFileWhenStartup(lockFilePath)
	if err != nil {
		return nil, err
	}

	// init kvs
	kvs.init()
	return kvs, nil
}

// Store API
type Kvsdb struct {
	// Options
	Opts *options.KvsOptions
	// The active segment file, writable
	activeFile *SegmentFile
	// The old segment files and hint files
	oldFileSet *SegmentFileSet
	// lock the home dir
	lockFp *os.File
	// Hash index
	keyDirs *KeyDirs
	// rw locker for get and set operations
	rwLock *sync.RWMutex
}

// Retrieve a value by key
func (kvs *Kvsdb) Get(key string) []byte {
	e := kvs.keyDirs.get(key)
	if e == nil {
		return nil
	}

	fileId := e.fileId
	sfp, err := kvs.getFileState(fileId)
	if err != nil {
		logger.Error(err)
		return nil
	}

	buf := make([]byte, e.vsz)
	bytesread, err := sfp.fp.ReadAt(buf, int64(e.vOffset))
	if err != nil || bytesread != int(e.vsz) {
		return nil
	}
	return buf
}

// Store a key and value
func (kvs *Kvsdb) Put(key string, value []byte) (bool, error) {
	return kvs.set(key, value, false)
}

//func (kvs *Kvsdb) Set(key string, value []byte) (bool, error) {
//	return kvs.set(key, value, false)
//}

// Delete a key
func (kvs *Kvsdb) Delete(key string) (bool, error) {
	return kvs.set(key, nil, true)
}

// Keys
func (kvs *Kvsdb) Keys() []string {
	kvs.rwLock.RLock()
	defer kvs.rwLock.RUnlock()
	return kvs.keyDirs.Keys()
}

// Merge
func (kvs *Kvsdb) Merge() (bool, error) {
	return true, nil
}

// Sync
func (kvs *Kvsdb) Sync() error {
	return nil
}

// Close
func (kvs *Kvsdb) Close() error {
	kvs.activeFile.Close()
	kvs.oldFileSet.Close()
	kvs.lockFp.Close()
	if kvs.lockFp != nil {
		_ = os.Remove(kvs.Opts.GetLockFilePath())
	}
	return nil
}

// Init kvs
func (kvs *Kvsdb) init() {
	kvs.oldFileSet = initOldFileSet()
	kvs.keyDirs = NewKeyDirs()

	// scan files
	hintFps, err := kvs.scan()
	if err != nil {
		os.Exit(1)
	}

	// parse hint files to rebuild key dir (memory hash index)
	kvs.parseHintFiles(hintFps)

	// get the last file id
	lastFileId, _ := util.GetLastHintFile(hintFps, kvs.Opts.GetFpFunc(false))
	kvs.setActiveFile(lastFileId)

	util.CloseHintFp(hintFps)
	util.WritePid(kvs.lockFp, lastFileId)
}

func (kvs *Kvsdb) scan() ([]*os.File, error) {
	infos, err := ioutil.ReadDir(kvs.Opts.DataDir)
	if err != nil {
		return nil, err
	}

	var hintFpaths []string
	for _, fi := range infos {
		if !fi.IsDir() && strings.Contains(fi.Name(), "hint") {
			hintFpaths = append(hintFpaths, path.Join(kvs.Opts.DataDir, fi.Name()))
		}
	}

	fps := make([]*os.File, 0)
	for _, fpath := range hintFpaths {
		fp, err := os.OpenFile(fpath, defaultOldFileMode, os.ModePerm)
		if err != nil {
			return nil, err
		}
		fps = append(fps, fp)
	}

	return fps, nil
}

// Building memory hash indexes
func (kvs *Kvsdb) parseHintFiles(hintFps []*os.File) {
	// hint file: ts,flag,ksz,vsz,vOffset,key
	b := make([]byte, HintHeaderSize, HintHeaderSize)

	for _, hfp := range hintFps {
		hName := hfp.Name()
		s := strings.LastIndex(hName, "/") + 1
		e := strings.LastIndex(hName, ".hint")
		fileId, _ := strconv.Atoi(hName[s:e])

		offset := int64(0)

		for {
			// Parse hint file header
			n, err := hfp.ReadAt(b, offset)
			offset += int64(n)

			if err == io.EOF {
				break
			}

			if (err != nil && err != io.EOF) || n != HintHeaderSize {
				panic(err)
			}

			// Hint header
			hh := DecodeHintHeader(b)
			//log.Println(hh)

			// Parse hint key
			keyByte := make([]byte, hh.ksz)
			n, err = hfp.ReadAt(keyByte, offset)
			if err == io.EOF {
				break
			}

			if (err != nil && err != io.EOF) || n != int(hh.ksz) {
				panic(err)
			}

			key := string(keyByte)
			offset += int64(hh.ksz)

			if hh.deleted() {
				kvs.keyDirs.del(key)
				continue
			}

			e := &entry{
				fileId:  uint32(fileId),
				vsz:     hh.vsz,
				vOffset: hh.vOffset,
				ts:      hh.ts,
				flag:    hh.flag,
			}

			// Put to keyDirs
			kvs.keyDirs.put(key, e)
		}
	}
}

// Write to active file and active hint file
func (kvs *Kvsdb) set(key string, value []byte, delete bool) (bool, error) {
	if !IsValidKey(key) {
		return false, InvalidKeyError
	}

	kvs.rwLock.Lock()
	defer kvs.rwLock.Unlock()

	kvs.checkWritable()

	// Write to active file
	e, err := kvs.activeFile.write(key, value, delete)
	if err != nil {
		kvs.rwLock.Unlock()
		return false, err
	}
	kvs.keyDirs.put(key, &e)
	return true, nil
}

func (kvs *Kvsdb) checkWritable() {
	// check active file size
	if kvs.activeFile.writeOffset >= kvs.Opts.MaxFileSize {
		newFileId := kvs.activeFile.fileId + 1

		logger.Infoln("Open a new data file and hint file, new fileId=", newFileId)

		// Close active file
		kvs.activeFile.Close()

		// Set a new active file
		kvs.setActiveFile(newFileId)

		// Write to lock file
		util.WritePid(kvs.lockFp, newFileId)
	}
}

func (kvs *Kvsdb) setActiveFile(fileId uint32) {
	dataFilename := kvs.Opts.GetDataFilePath(uint32(fileId))
	activeFp, err := os.OpenFile(dataFilename, defaultActiveFileMode, os.ModePerm) // |os.O_APPEND, use WriteAt
	if err != nil {
		panic(err)
	}
	stat, _ := activeFp.Stat()

	logger.Infoln("Set the active file, fileId=", fileId, " active file size=", stat.Size())

	kvs.activeFile = newSegmentFile(true)
	kvs.activeFile.fileId = uint32(fileId)
	kvs.activeFile.fp = activeFp

	kvs.activeFile.writeOffset = uint64(stat.Size())

	hintFilename := kvs.Opts.GetHintFilePath(uint32(fileId))
	hintFp, err := os.OpenFile(hintFilename, defaultActiveFileMode, os.ModePerm) // |os.O_APPEND, use WriteAt
	if err != nil {
		panic(err)
	}
	kvs.activeFile.hintFp = hintFp
}

func (kvs *Kvsdb) getFileState(fileId uint32) (*SegmentFile, error) {
	if fileId == kvs.activeFile.fileId {
		return kvs.activeFile, nil
	}

	// Look up in the old files
	sFile := kvs.oldFileSet.get(fileId)
	if sFile != nil {
		return sFile, nil
	}

	sFile, err := openSFile(kvs.Opts.GetDataFilePath(fileId), fileId)
	if err != nil {
		return nil, err
	}
	kvs.oldFileSet.put(sFile, fileId)
	return sFile, nil
}

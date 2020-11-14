package options

import (
	"fmt"
	"os"
	"path"
	"reflect"

	"github.com/shniu/gostuff/kvs/log"
)

var logger = log.Logger

const (
	Bitcask = iota
	Leveldb
)

// Todo: Sync Strategy, Hint File CRC Check, Merge Policy(https://docs.basho.com/riak/kv/2.0.0/setup/planning/backend/bitcask/#merge-policy)
var Options = &KvsOptions{
	HostName:       "LocalKvs",
	Listen:         "",
	Port:           3000,
	Home:           "/tmp/kvs",
	LockFilename:   "kvs.lock",
	DataFileSuffix: ".data",
	HintFileSuffix: ".hint",

	// 1 << 27 --> 128 MB
	// 1 << 28 --> 256 MB
	// 1 << 29 --> 512 MB
	// 1 << 30 --> 1024 MB
	// 1 << 31 --> 2048 MB
	// 1 << 26 --> 64 MB
	MaxFileSize: 1 << 26,
}

func init() {
	Options.init()
}

type KvsOptions struct {
	// Server options
	HostName string
	Listen   string
	Port     int

	// Backend engine type, ref const: Bitcask, Leveldb
	EngineType int

	// Storage engine options
	Home           string
	DataDir        string
	DataFileSuffix string
	HintFileSuffix string
	LockFilename   string
	MaxFileSize    uint64
}

func (o *KvsOptions) init() {
	dataDir := path.Join(o.Home, "data")
	// Mkdir for home dir
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		logger.Panicf("Fail to init home and data dir %s \n", dataDir)
	}

	o.DataDir = dataDir
}

// Get data file path
// e.g. if fileId == 0, {KVS_HOME}/data/000.data
func (o *KvsOptions) GetDataFilePath(fileId uint32) string {
	filename := fmt.Sprintf("%03d%s", fileId, o.DataFileSuffix)
	return path.Join(o.DataDir, filename)
}

func (o *KvsOptions) GetHintFilePath(fileId uint32) string {
	filename := fmt.Sprintf("%03d%s", fileId, o.HintFileSuffix)
	return path.Join(o.DataDir, filename)
}

func (o *KvsOptions) GetFpFunc(active bool) func(int) string {
	var suffix string
	if active {
		suffix = o.DataFileSuffix
	} else {
		suffix = o.HintFileSuffix
	}

	return func(fileId int) string {
		filename := fmt.Sprintf("%03d%s", fileId, suffix)
		return path.Join(o.DataDir, filename)
	}
}

func (o *KvsOptions) GetLockFilePath() string {
	return path.Join(o.Home, o.LockFilename)
}

// Print all of kvs options
func (o *KvsOptions) Print() {
	t := reflect.TypeOf(*o)
	v := reflect.ValueOf(*o)
	optionsFieldNum := v.NumField()

	for k := 0; k < optionsFieldNum; k++ {
		logger.Infof("\t%s: %v", t.Field(k).Name, v.Field(k).Interface())
	}
}

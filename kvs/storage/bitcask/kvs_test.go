package bitcask

import (
	"fmt"
	"github.com/shniu/gostuff/kvs/options"
	"github.com/stretchr/testify/assert"
	"testing"
)

var kvs *Kvsdb
var err error

func init() {
	opts := options.Options
	// Test Open
	kvs, err = Open(opts)
	if err != nil {
		//log.Panicln("Init kvs failed!", err)
	}
}

func TestKVStore_GetNil(t *testing.T) {
	v := kvs.Get("item:k1__")
	//log.Println(v == nil)
	assert.Equal(t, []byte(nil), v, "kvs.Get should return nil")
}

func TestKVStore_GetTheActualValue(t *testing.T) {
	expectedVal := "UnitTest"
	key := "__t__"

	helperSet(t, key, []byte(expectedVal))

	v := kvs.Get(key)
	assert.Equal(t, expectedVal, string(v), fmt.Sprintf("kvs.Get(\"%s\") should return UnitTest", key))
}

func TestKVStore_Delete(t *testing.T) {
	expectedVal := "DeleteTest"
	key := "__del__"

	helperSet(t, key, []byte(expectedVal))

	b, e := kvs.Delete(key)
	assert.NoErrorf(t, e, "kvs.Delete(k) should not errors")
	assert.True(t, b, "kvs.Delete(k) should return true")

	v := kvs.Get(key)
	assert.Equal(t, []byte(nil), v, "kvs.Get should return nil")
}

func helperSet(t *testing.T, key string, val []byte) {
	// b, e := kvs.Set(key, []byte(val))
	b, e := kvs.Put(key, []byte(val))
	assert.NoErrorf(t, e, "kvs.Set(k,v) should not errors")
	assert.True(t, b, "kvs.Set(k,v) should return true")
}

func TestNewKVStore(t *testing.T) {
	//kvs, err := Open("/tmp/kvs")
	//if kvs != nil && err == nil {
	//	t.Log("Test passed!")
	//}
	//
	//ts := time.Now().Unix()
	//t.Log(ts)
	//
	//crc32q := crc32.MakeTable(0xD5828281)
	//crc32V1 := crc32.Checksum([]byte("hello"), crc32q)
	//crc32V2 := crc32.ChecksumIEEE([]byte("hello"))
	//t.Log(crc32V1, crc32V2)
}

func TestKVStore_Set(t *testing.T) {
	//var buf Bytes.Buffer
	//buf.Write([]byte("Hello\n"))
	//_, _ = buf.WriteTo(os.Stdout)

	//var fn = "tmp.out"
	//f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	fmt.Printf("Open %s fialed", fn)
	//	os.Exit(1)
	//}
	//defer f.Close()
	//
	//buf.Write([]byte("ssssssss\n"))
	//f.Write(buf.Bytes())
	//buf.Reset()
	/*_, err = buf.WriteTo(f)
	if err != nil {
		fmt.Printf("Write to %s fialed", fn)
	}*/

	// Write binary
	/*for i := 0; i < 10; i++ {
		binary.Write(&buf, binary.LittleEndian, int32(i))
		f.Write(buf.Bytes())
	}*/

	//idx := &MemIdx{FileId: "000", ValueSz: 10, ValuePos: 20, Ts: 1000000}
	//fo, err := os.OpenFile("obj.bin", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//defer fo.Close()
	//
	//buffer := new(Bytes.Buffer)
	//
	//err = binary.Write(buffer, binary.LittleEndian, idx)
	//fmt.Println(buffer.Len(), buffer.Bytes())
	//f.Write(buf.Bytes())
	//f.Sync()
	//
	//fr, err := os.Open("obj.bin")
	//defer fr.Close()
	//dataBytes := make([]byte, unsafe.Sizeof(MemIdx{}))
	//idxR := MemIdx{}
	//n, err := fr.Read(dataBytes)
	//dataBytes = dataBytes[:n]
	//binary.Read(Bytes.NewBuffer(dataBytes), binary.LittleEndian, &idxR)
	//fmt.Println(idxR)

}

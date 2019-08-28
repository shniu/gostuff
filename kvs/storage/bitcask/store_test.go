package bitcask

import (
	"testing"
)

func TestTest(t *testing.T) {
	//Test()
}

func TestKvs(t *testing.T) {
	//var testKey = "item_1234"
	//var fp = "/tmp/kvs"

	//var kvs SimpleKvs
	//kvs = &SimpleKvsImpl{
	//	Keydir: make(map[string]MemIdx),
	//}
	//err := kvs.Open(fp)
	//if err != nil {
	//	fmt.Printf("Open %s failed!", fp)
	//}

	// Put
	//succeed, err := kvs.Put(testKey, "{\"name\": 123456}")

	//testVal, err := kvs.Get(testKey)
	//fmt.Print(testVal)
}

/*func createAppendOnlyFile(filePath string) *os.File {

	//fd,_:=os.OpenFile("a.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	//fd_time:=time.Now().Format("2006-01-02 15:04:05");
	//fd_content:=strings.Join([]string{"======",fd_time,"=====",str_content,"\n"},"")
	//buf:=[]byte(fd_content)
	//fd.Write(buf)
	//fd.Close()

	var fd *os.File
	var err errors
	_, err = os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		fd, err = os.Create(filePath)
	}

	fd, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	return fd
}

//logFd := createAppendOnlyFile("000.kvs")
//_, _ = logFd.WriteString("abc,{\"name\": \"cat\"}\n")
//_, _ = logFd.WriteString("key,{\"name\":\"kkk\"}\n")
//closeFd(logFd)
func closeFd(f *os.File) {
	if f != nil {
		defer f.Close()
	}
}*/

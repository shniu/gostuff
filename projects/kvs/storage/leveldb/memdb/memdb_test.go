package memdb

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

const MaxRandInt = 100000

// @link https://play.golang.org/p/ZdFpbahgC1
func randPutToSkipList(skipList *SkipList, num int) {
	// Change the default Source
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		// a non-negative pseudo-random number in [0,n) from the default Source
		r := rand.Intn(MaxRandInt)
		// strconv.Itoa(r)
		skipList.Put(r, fmt.Sprintf("%v", r))
	}
}

func TestSkipList_Put_Get_Succeed(t *testing.T) {
	t.Log("Test skiplist Get, Put, Delete")

	skipList := NewIntSkipList()
	// rand
	randPutToSkipList(skipList, 100)

	skipList.Put(77, "seven")
	skipList.Put(10, "one")
	skipList.Put(130, "thirteen")

	logger.Debug("SkipList.length:", skipList.length)
	logger.Debug("SkipList.level:", len(skipList.header.forwards))

	n := skipList.header
	logger.Debug(">>>>>>>>>>>>")
	for n != nil {
		logger.Debug(n)
		n = n.forwards[0]
	}
	logger.Debug(">>>>>>>>>>>>")

	val, ok := skipList.Get(77)
	t.Log("key=77, val=", val)
	assert.Equal(t, "seven", val)
	assert.True(t, ok)

	// Put exists key
	skipList.Put(77, "put again")
	val, ok = skipList.Get(77)
	t.Log("key=77, new val=", val)
	assert.Equal(t, "put again", val)
	assert.True(t, ok)

	// Delete
	valDeleted, ok := skipList.Delete(130)
	assert.True(t, ok)
	assert.Equal(t, "thirteen", valDeleted)

	valGot, ok := skipList.Get(130)
	assert.False(t, ok)
	assert.Equal(t, nil, valGot)
	t.Log("Delete test done")
}

func TestSkipList_Nil_key(t *testing.T) {
	skipList := NewIntSkipList()

	// put nil key
	assert.Panics(t, func() {
		//defer func() {
		//	if r := recover(); r == nil {
		//		t.Errorf("The code did not panic")
		//	}
		//}()

		defer func() {
			t.Log("Put nil key will raise panic")
		}()

		skipList.Put(nil, "nil key")
	})

	// Delete nil key
	assert.Panics(t, func() {
		defer func() {
			t.Log("Delete nil key will raise panic")
		}()
		skipList.Delete(nil)
	})
	t.Log("Test nil key done")
}

func TestSkipList_Iterator(t *testing.T) {
	var randElementCnt = 20

	skipList := NewIntSkipList()
	randPutToSkipList(skipList, randElementCnt)

	var count = 0
	iterator := skipList.Iterator()

	// Next
	for iterator.Next() {
		t.Log("key=", iterator.Key(), ", value=", iterator.Value())
		count++
	}
	assert.True(t, count <= randElementCnt)
	t.Log(">>>>>>>>>>>>>>>> E N D N E X T <<<<<<<<<<<<<<")

	// Previous
	for iterator.Previous() {
		t.Log("key=", iterator.Key(), ", value=", iterator.Value())
	}
	t.Log(">>>>>>>>>>>>>>>> E N D P R E V I O U S <<<<<<<<<<<<<<")

	// Seek
	skipList.Put(999, "000000000")
	ok := iterator.Seek(999)
	assert.True(t, ok)
	assert.Equal(t, "000000000", iterator.Value())
	t.Log(">>>>>>>>>>>>>>>> E N D S E E K <<<<<<<<<<<<<<")

	// Close
	iterator.Close()
	assert.True(t, iterator.Key() == nil)
	assert.True(t, iterator.Value() == nil)
	assert.False(t, iterator.Next())
	assert.False(t, iterator.Previous())
	t.Log(">>>>>>>>>>>>>>>> E N D C L O S E <<<<<<<<<<<<<<")
}

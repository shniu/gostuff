package memdb

// https://github.com/ryszard/goskiplist/blob/master/skiplist/skiplist.go
// https://en.wikipedia.org/wiki/Skip_list
// Implementation by skip list

import (
	"github.com/shniu/gostuff/kvs/log"
	"github.com/shniu/gostuff/kvs/util"
	"math/rand"
	"sync"
)

var logger = log.Logger

// in-memory key/value database
type DB struct {
	rwLock sync.RWMutex
	kvData []byte
}

func (db *DB) Get(key []byte) (value []byte, err error) {
	db.rwLock.RLock()
	defer db.rwLock.RUnlock()

	return
}

func (db *DB) Put(key, value []byte) error {

	return nil
}

const DefaultMaxLevel = 32
const p = 0.5

type node struct {
	// key value
	key, value interface{}
	// forward nodes
	forwards []*node
	// backward node
	backward *node
	// level of the node
	// maxLevel int
	// flag
	flag string
}

// Next node
func (n *node) next() *node {
	if len(n.forwards) == 0 {
		return nil
	}
	return n.forwards[0]
}

// Previous node
func (n *node) previous() *node {
	return n.backward
}

// Has next?
func (n *node) hasNext() bool {
	return n.next() != nil
}

// Has previous
func (n *node) hasPrevious() bool {
	return n.previous() != nil
}

// Iterator is an interface that you can use to iterate through the
// skip list (in its entirety or fragments).
type Iterator interface {
	// Next returns true if the iterator contains subsequent elements
	// and advances its state to the next element if that is possible.
	Next() (ok bool)

	// Previous returns true if the iterator contains previous elements
	// and rewinds its state to the previous element if that is possible.
	Previous() (ok bool)

	// Key returns the current key.
	Key() interface{}

	// Value returns the current value.
	Value() interface{}

	// Seek reduces iterative seek costs for searching forward into the Skip List
	// by remarking the range of keys over which it has scanned before. If the
	// requested key occurs prior to the point, the Skip List will start searching
	// as a safeguard.  It returns true if the key is within the known range of the list.
	Seek(key interface{}) (ok bool)

	// Close this iterator to reap resources associated with it.  While not
	// strictly required, it will provide extra hints for the garbage collector.
	Close()
}

// iter
type iter struct {
	current *node
	key     interface{}
	list    *SkipList
	value   interface{}
}

func (i *iter) Next() (ok bool) {
	if i.current == nil || !i.current.hasNext() {
		return false
	}

	i.current = i.current.next()
	i.key = i.current.key
	i.value = i.current.value
	return true
}

func (i iter) Key() interface{} {
	return i.key
}

func (i iter) Value() interface{} {
	return i.value
}

func (i *iter) Previous() (ok bool) {
	if i.current == nil || !i.current.hasPrevious() {
		return false
	}

	i.current = i.current.previous()
	i.key = i.current.key
	i.value = i.current.value
	return true
}

func (i *iter) Seek(key interface{}) (ok bool) {
	current := i.current
	list := i.list

	if current == nil {
		current = list.header
	}

	if current.key != nil && list.lessThan(key, current.key) {
		current = list.header
	}

	if current.backward == nil {
		current = list.header
	} else {
		current = current.backward
	}

	current = list.getPath(current, nil, key)
	if current == nil {
		return false
	}

	i.current = current
	i.key = current.key
	i.value = current.value
	return true
}

func (i *iter) Close() {
	i.key = nil
	i.value = nil
	i.current = nil
	i.list = nil
}

type rangeiter struct {
	iter
	upperLimit interface{}
	lowerLimit interface{}
}

func (i *rangeiter) Next() bool {
	if !i.current.hasNext() {
		return false
	}

	next := i.current.next()
	if !i.list.lessThan(next.key, i.upperLimit) {
		return false
	}

	i.current = i.current.next()
	i.key = i.current.key
	i.value = i.current.value
	return true
}

func (i *rangeiter) Previous() bool {
	// Todo
	return true
}

func (i *rangeiter) Seek() bool {
	// Todo
	return true
}

func (i *rangeiter) Close() {
	i.iter.Close()
	i.upperLimit = nil
	i.lowerLimit = nil
}

// Skip list that maintains an ordered collection of key-value pairs
type SkipList struct {
	// header
	header *node
	// footer
	footer *node
	// max level
	MaxLevel int
	// compare func
	lessThan func(l, r interface{}) bool
	// length
	length int
}

// level of the skip list
func (s *SkipList) level() int {
	return len(s.header.forwards) - 1
}

// Iterator returns an Iterator that will go through all elements s.
func (s *SkipList) Iterator() Iterator {
	return &iter{
		current: s.header,
		list:    s,
	}
}

// Put put the value associated with key in s.
func (s *SkipList) Put(key, value interface{}) {
	logger.Debug(">>>>>>>>>>>> in SkipList.Put <<<<<<<<<<<<<<<<")
	logger.Debug("key=", key, ", value=", value)
	s.checkKey(key)

	// Find candidate node
	update := make([]*node, s.level()+1, s.effectiveMaxLevel()+1)
	candidateNode := s.getPath(s.header, update, key)
	logger.Debug("update len: ", len(update), ", update[0]: ", update[0],
		", update[0].flag: ", update[0].flag, ", update[0].forwards[0]: ", update[0].forwards[0])

	// The key is exist, then update the value
	if candidateNode != nil && candidateNode.key == key {
		candidateNode.value = value
		logger.Debug("=========================")
		return
	}

	// Random Level
	newLevel := s.randomLevel()
	logger.Debug(">>>>>>>>> N E W L E V E L:", newLevel)
	if currentLevel := s.level(); newLevel > currentLevel {
		// there are no pointers for the higher levels in
		// update. Header should be there. Also add higher
		// level links to the header.
		for i := currentLevel; i <= newLevel; i++ {
			update = append(update, s.header)
			s.header.forwards = append(s.header.forwards, nil)
		}
	}

	// New node
	newNode := &node{
		key:      key,
		value:    value,
		flag:     "new node",
		forwards: make([]*node, newLevel+1, s.effectiveMaxLevel()+1),
	}

	if previous := update[0]; previous != nil && previous.key != nil {
		newNode.backward = previous
	}

	for i := 0; i <= newLevel; i++ {
		//logger.Debug("Before newNode: ", newNode)
		newNode.forwards[i] = update[i].forwards[i]
		update[i].forwards[i] = newNode
		//logger.Debug("After newNode: ", newNode)
	}
	s.length++

	if newNode.forwards[0] != nil {
		if newNode.forwards[0].backward != newNode {
			newNode.forwards[0].backward = newNode
		}
	}

	if s.footer == nil || s.lessThan(s.footer.key, key) {
		s.footer = newNode
	}

	logger.Debug(">>>>>>>>>>>> out SkipList.Put <<<<<<<<<<<<<<<<\n\n")
}

// Get get the the value associated with key in s.
// if not exists, return nil
func (s *SkipList) Get(key interface{}) (interface{}, bool) {
	candidateNode := s.getPath(s.header, nil, key)
	//for candidateNode != nil && s.lessThan(candidateNode.key, key) {
	//	candidateNode = candidateNode.forwards[0]
	//}
	if candidateNode != nil && candidateNode.key == key {
		return candidateNode.value, true
	}
	return nil, false
}

// Delete
func (s *SkipList) Delete(key interface{}) (interface{}, bool) {
	s.checkKey(key)

	// find the key in s
	// Find candidate node
	update := make([]*node, s.level()+1, s.effectiveMaxLevel()+1)
	candidateNode := s.getPath(s.header, update, key)

	if candidateNode == nil || candidateNode.key != key {
		return nil, false
	}

	// bingo, do delete
	previous := candidateNode.backward
	// candidateNode is the last node
	if s.footer == candidateNode {
		s.footer = previous
	}

	next := candidateNode.next()
	if next != nil {
		next.backward = previous
	}

	for i := 0; i <= s.level() && update[i].forwards[i] == candidateNode; i++ {
		update[i].forwards[i] = candidateNode.forwards[i]
	}

	// the top level's forward node is nil or not
	for s.level() > 0 && s.header.forwards[s.level()] == nil {
		s.header.forwards = s.header.forwards[:s.level()]
	}
	s.length--

	return candidateNode.value, true
}

// Seek
func (s *SkipList) Seek(key interface{}) Iterator {
	current := s.getPath(s.header, nil, key)
	if current == nil {
		return nil
	}

	return &iter{
		current: current,
		key:     current.key,
		value:   current.value,
		list:    s,
	}
}

func (s SkipList) checkKey(key interface{}) {
	if key == nil {
		panic("skiplist: nil keys are not supported")
	}
}

func (s SkipList) randomLevel() (n int) {
	for n = 0; n < s.effectiveMaxLevel() && rand.Float64() < p; n++ {
	}
	return
}

func (s *SkipList) effectiveMaxLevel() int {
	return util.MaxInt(s.level(), s.MaxLevel)
}

// getPath find candidate node
func (s *SkipList) getPath(current *node, update []*node, key interface{}) *node {
	depth := len(current.forwards) - 1

	for i := depth; i >= 0; i-- {
		for current.forwards[i] != nil && s.lessThan(current.forwards[i].key, key) {
			current = current.forwards[i]
		}
		if update != nil {
			update[i] = current
		}
	}

	return current.next()
}

// ---
func NewCustomSkipList(lessThan func(l, r interface{}) bool) *SkipList {
	return &SkipList{
		header: &node{
			flag:     "header",
			forwards: []*node{nil},
		},
		lessThan: lessThan,
		MaxLevel: DefaultMaxLevel,
	}
}

// Ordered is an interface which can be linearly ordered by the
// LessThan method, whereby this instance is deemed to be less than
// other. Additionally, Ordered instances should behave properly when
// compared using == and !=.
type Ordered interface {
	LessThan(other Ordered) bool
}

// New returns a new SkipList.
// Its keys must implement the Ordered interface.
func New() *SkipList {
	comparator := func(left, right interface{}) bool {
		return left.(Ordered).LessThan(right.(Ordered))
	}
	return NewCustomSkipList(comparator)
}

// returns a SkipList that accepts int keys.
func NewIntSkipList() *SkipList {
	return NewCustomSkipList(func(l, r interface{}) bool {
		return l.(int) < r.(int)
	})
}

// returns a SkipList that accepts string keys.
func NewStringSkipList() *SkipList {
	return NewCustomSkipList(func(l, r interface{}) bool {
		return l.(string) < r.(string)
	})
}

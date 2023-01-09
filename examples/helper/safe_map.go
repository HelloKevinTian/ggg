package helper

import (
	"sync"
)

// SafeMap 安全的Map
type SafeMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

// Put 存储操作
func (sm *SafeMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	sm.data[k] = v
}

// Get 获取操作
func (sm *SafeMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	return sm.data[k]
}

// Del 删除操作
func (sm *SafeMap) Del(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	delete(sm.data, k)
}

// Each 遍历Map，并且把遍历的值给回调函数，可以让调用者控制做任何事情
func (sm *SafeMap) Each(cb func(interface{}, interface{})) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	for k, v := range sm.data {
		cb(k, v)
	}
}

// NewSafeMap 生成初始化一个SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{
		rw:   new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}

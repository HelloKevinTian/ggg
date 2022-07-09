package singleton

import (
	"fmt"
	"sync"
)

type LazySingleton struct{}

func (l *LazySingleton) Echo() {
	fmt.Printf("%p ", l)
}

var (
	lazySingleton *LazySingleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *LazySingleton {
	once.Do(func() {
		lazySingleton = &LazySingleton{}
	})
	return lazySingleton
}

//------------------------------------------

type singleton struct{}

var instance *singleton
var lock1 sync.Mutex

// GetInstance
// 注意：这种写法会产生竞态
func GetInstance() *singleton {
	if instance == nil {
		lock1.Lock()
		if instance == nil {
			instance = new(singleton)
		}
		lock1.Unlock()
	}
	return instance
}

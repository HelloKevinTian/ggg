// 并发安全的懒汉单例模式
package singleton

import (
	"fmt"
	"sync"
	"time"
)

var (
	mongoSession sync.Map //推荐使用sync.Map
	// mongoSession = make(map[string]*MClient) //使用map会有竞态问题
	lock sync.Mutex
)

type MClient struct {
	url string
}

func getMongoSession(url string) *MClient {
	session, ok := mongoSession.Load(url)
	// session, ok := mongoSession[url]
	if !ok {
		lock.Lock()
		// session, ok = mongoSession[url]
		session, ok = mongoSession.Load(url)
		if !ok {
			fmt.Println("new mongodb client ", url)
			// fmt.Println(runtime.NumGoroutine())
			session = &MClient{url: url}

			// mongoSession[url] = session
			mongoSession.Store(url, session)
		}
		lock.Unlock()
	}

	return session.(*MClient)
}

func Test() {
	for i := 0; i < 100; i++ {
		go func(ii int) {
			getMongoSession("111")

		}(i)
	}

	for i := 0; i < 10; i++ {
		go func(ii int) {
			v := GetLazyInstance()
			v.Echo()
		}(i)
	}

	for i := 0; i < 10; i++ {
		go func(ii int) {
			GetInstance()
		}(i)
	}

	time.Sleep(3 * time.Second)

}

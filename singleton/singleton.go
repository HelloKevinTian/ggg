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
	for i := 0; i < 10000; i++ {
		go func(ii int) {
			// d.getCollect()
			getMongoSession("111")
			// getMongoSession("222")
			// fmt.Println(">>> ", ii, c.Age)

		}(i)
	}

	time.Sleep(3 * time.Second)
}

package some

import (
	"fmt"
	"sync"
	"time"
)

func BatchN() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan struct{}, 3)
	var m sync.Map
	ws := sync.WaitGroup{}
	ws.Add(len(arr))

	for _, v := range arr {
		c <- struct{}{}
		go delay(v, c, &m, &ws)
	}

	ws.Wait()

	m.Range(func(k, v interface{}) bool {
		fmt.Println(">>>>>>", k, v)
		return true
	})
}

func delay(num int, c <-chan struct{}, m *sync.Map, ws *sync.WaitGroup) {
	time.AfterFunc(time.Second*2, func() {
		defer func() {
			ws.Done()
			<-c
		}()
		fmt.Println(num)
		m.Store(num*100, num*10000)
	})
}

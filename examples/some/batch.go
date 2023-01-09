//批处理计算
package some

import (
	"fmt"
	"time"
)

func TestBatch() {
	s := []int{}
	sc := make(chan int, 100)
	limitCh := make(chan struct{}, 10)

	go func() {
		for i := 0; i < 100; i++ {
			limitCh <- struct{}{}
			go func(i int) {
				defer func() {
					<-limitCh
				}()
				time.Sleep(2 * time.Second)
				sc <- i
			}(i)
		}
	}()

	for v := range sc {
		s = append(s, v)
		fmt.Println("---- ", len(s))
		if len(s) >= 100 {
			close(sc)
		}
	}

	fmt.Println(s)
}

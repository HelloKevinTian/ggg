package utils

import (
	"fmt"
	"ggg/utils/cache"
	"ggg/utils/duration"
	"ggg/utils/sets"
	"time"
)

func Test() {
	//test duration
	s1 := duration.ShortHumanDuration(1951 * time.Second)
	s2 := duration.HumanDuration(28765 * time.Second)
	fmt.Println(s1, s2)

	//test sets.NewByte
	bsets := sets.NewByte()
	bsets.Insert(byte(11))
	bsets.Insert(byte(12))
	bsets1 := sets.NewByte()
	bsets1.Insert(byte(12))
	bsets1.Insert(byte(13))
	a1 := bsets.Intersection(bsets1)
	a2 := bsets.Union(bsets1)
	fmt.Println(a1, a2)

	//test sets.NewString
	ssets := sets.NewString()
	ssets.Insert("a")
	ssets.Insert("b")
	ssets1 := sets.NewString()
	ssets1.Insert("b")
	ssets1.Insert("c")
	a11 := ssets.Intersection(ssets1)
	a22 := ssets.Union(ssets1)
	fmt.Println(a11, a22)

	//test lrucache
	c := cache.NewLRUExpireCache(100)
	c.Add("foo", "bar", 10*time.Second)
	c.Add("age", 18, 2*time.Second)
	time.Sleep(2 * time.Second)
	k1, b1 := c.Get("foo")
	k2, b2 := c.Get("age")
	fmt.Println(k1, b1, k2, b2)
}

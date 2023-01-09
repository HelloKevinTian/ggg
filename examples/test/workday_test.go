//go test -benchmem -run=^$ -bench ^Benchmark* gcargo/test
// goos: darwin
// goarch: amd64
// pkg: gcargo/test
// cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
// BenchmarkFindWorkDay-8           9851278               123.9 ns/op             0 B/op          0 allocs/op
// BenchmarkGetBetweenDays-8         268262              4145 ns/op               0 B/op          0 allocs/op
// PASS
// ok      gcargo/test     2.510s
package test

import (
	"testing"
	"time"
)

//推荐写法
func FindWorkDay(s, e time.Time) (workDays int) {
	ctz, _ := time.LoadLocation("Asia/Shanghai")
	s = s.In(ctz)
	e = e.In(ctz)
	sWeek := s.Weekday()
	diffDay := int(e.Sub(s).Hours() / 24)
	for i := int(sWeek); i <= int(sWeek)+diffDay; i++ {
		w := time.Weekday(i % 7)
		if w != time.Sunday && w != time.Saturday {
			workDays++
		}
	}
	return
}

//不推荐写法
func GetBetweenDays(sdate, edate time.Time) int {
	d := 0
	if sdate.Weekday() != time.Saturday && sdate.Weekday() != time.Sunday {
		d++
	}
	for {
		sdate = sdate.AddDate(0, 0, 1)
		if edate.Before(sdate) {
			break
		}
		if sdate.Weekday() != time.Saturday && sdate.Weekday() != time.Sunday {
			d++
		}
	}
	return d
}

var s, _ = time.Parse("2006-01-02", "2021-11-01")
var e, _ = time.Parse("2006-01-02", "2021-12-30")

func BenchmarkFindWorkDay(b *testing.B) { //BenchmarkFindWorkDay-8   	42167194	        28.32 ns/op	       0 B/op	       0 allocs/op
	for i := 0; i < b.N; i++ {
		FindWorkDay(s, e)
	}
}

func BenchmarkGetBetweenDays(b *testing.B) { //BenchmarkGetBetweenDays-8   	17730579	        72.44 ns/op	       0 B/op	       0 allocs/op
	for i := 0; i < b.N; i++ {
		GetBetweenDays(s, e)
	}
}

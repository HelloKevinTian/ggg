package some

import "testing"

var s1 = []string{"a", "b", "c", "d", "a", "b", "c", "a", "b", "a"}
var s2 = []string{"a", "b", "c", "e", "a", "b", "c", "a", "b", "a"}

func BenchmarkIntersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intersect(s1, s2)
	}
}

func BenchmarkIntersect1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intersect1(s1, s2)
	}
}

func BenchmarkIntersect2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intersect2(s1, s2)
	}
}

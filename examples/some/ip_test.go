package some

import (
	"testing"

	"github.com/3th1nk/cidr"
)

// BenchmarkCIDRContains-8   	24931017	        48.00 ns/op	      16 B/op	       1 allocs/op
func BenchmarkCIDRContains(b *testing.B) {
	c, _ := cidr.Parse("192.168.1.0/28")
	for i := 0; i < b.N; i++ {
		// c.Contains("192.168.1.17")
		c.Contains("192.168.1.12")
	}
}

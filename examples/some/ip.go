package some

import (
	"fmt"

	"github.com/3th1nk/cidr"
)

// https://github.com/3th1nk/cidr
func TestCIDR() {
	// parses a network segment as a CIDR
	c, _ := cidr.Parse("192.168.1.0/28")
	fmt.Println("network:", c.Network(), "broadcast:", c.Broadcast(), "mask", c.Mask())

	// ip contains
	b1 := c.Contains("192.168.1.17")
	b2 := c.Contains("192.168.1.12")
	fmt.Println("192.168.1.17 Contains: ", b1)
	fmt.Println("192.168.1.12 Contains: ", b2)

	// ip range
	beginIP, endIP := c.IPRange()
	fmt.Println("ip range:", beginIP, endIP)

	// iterate through each ip
	fmt.Println("ip total:", c.IPCount())
	c.Each(func(ip string) bool {
		fmt.Println("Each\t", ip)
		return true
	})
	c.EachFrom("192.168.1.10", func(ip string) bool {
		fmt.Println("EachFrom\t", ip)
		return true
	})

	fmt.Println("subnet plan based on the subnets num:")
	cs, _ := c.SubNetting(cidr.MethodSubnetNum, 4)
	for _, c := range cs {
		fmt.Println("\t", c.CIDR())
	}

	fmt.Println("subnet plan based on the hosts num:")
	cs, _ = c.SubNetting(cidr.MethodHostNum, 4)
	for _, c := range cs {
		fmt.Println("\t", c.CIDR())
	}

	fmt.Println("merge network:")
	c, _ = cidr.SuperNetting([]string{
		"2001:db8::/66",
		"2001:db8:0:0:8000::/66",
		"2001:db8:0:0:4000::/66",
		"2001:db8:0:0:c000::/66",
	})
	fmt.Println("\t", c.CIDR())
}

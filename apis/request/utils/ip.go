package utils

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/brotherpowers/ipsubnet"
)

type Range struct {
	Start [4]int
	End   [4]int
}

func PingRange() {
	// range:=Range{Start:"0.0.0.0",End:"0.0.0.10"}
}

func NetworkRanges() []Range {
	ranges := getRanges()
	for i := range ranges {
		fmt.Printf("from %v to %v\n", ranges[i].Start, ranges[i].End)
	}
	return ranges
}

func getRanges() []Range {
	addresses := localAddresses()
	output := []Range{}
	for i := range addresses {
		if !regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`).MatchString(addresses[i]) {
			continue
		}
		ip, mask := parseAddress(addresses[i])
		sub := ipsubnet.SubnetCalculator(ip, mask)
		startend := sub.GetIPAddressRange()
		output = append(output, Range{parseIP(startend[0]), parseIP(startend[1])})
	}
	return output
}

func parseIP(input string) [4]int {
	o := strings.Split(input, ".")
	out := [4]int{}
	for i := range o {
		v, _ := strconv.Atoi(o[i])
		out[i] = v
	}
	return out
}

func parseAddress(input string) (string, int) {
	val := strings.Split(input, "/")
	mask, _ := strconv.Atoi(val[1])
	return val[0], mask
}

func localAddresses() []string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	output := []string{}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPNet:
				output = append(output, fmt.Sprintf("%+v", v))
			}

		}
	}
	return output
}

package ipfind

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/brotherpowers/ipsubnet"
)

type Range struct {
	Start string
	End   string
}

func rangesMain() { getRanges() }

func FindHosts() ([]string, error) {
	ranges := getRanges()
	for i := range ranges {
		fmt.Printf("from %s to %s\n", ranges[i].Start, ranges[i].End)
	}
	return nil, nil
}

func getRanges() []Range {
	addresses := rangesLocalAddresses()
	output := []Range{}
	for i := range addresses {
		if !regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`).MatchString(addresses[i]) {
			continue
		}
		ip, mask := parseAddress(addresses[i])
		sub := ipsubnet.SubnetCalculator(ip, mask)
		fmt.Printf("%#v\n", sub)
		startend := sub.GetIPAddressRange()
		fmt.Printf("%#v\n", startend)
		// output = append(output, Range{startend[0], startend[1]})
	}
	return output
}

func parseAddress(input string) (string, int) {
	val := strings.Split(input, "/")
	mask, _ := strconv.Atoi(val[1])
	return val[0], mask
}

func rangesLocalAddresses() []string {
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

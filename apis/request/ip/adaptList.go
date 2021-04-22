package ipfind

import (
	"fmt"
	"net"
)

func adaptListLocalAddresses() []string {
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

func adaptListMain() {
	for _, i := range adaptListLocalAddresses() {
		fmt.Printf("%+v\n", i)
	}
}

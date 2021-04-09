package ipfind

import (
	"fmt"
	"net"
)

func ownIPMain() {
	ip, ipnet, err := net.ParseCIDR("172.17.0.1/16")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ip:%+v\n", ip)
	fmt.Printf("ipnet:%+v\n", ipnet)
}

func t() {
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			fmt.Println(ip)
		}
	}
}

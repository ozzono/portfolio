package ipfind

import (
	"fmt"

	"github.com/brotherpowers/ipsubnet"
)

func subnetMain() {

	sub := ipsubnet.SubnetCalculator("172.17.0.1", 16)
	fmt.Printf("%#v\n", sub.GetIPAddressRange())
}

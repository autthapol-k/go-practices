package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	parts := make([]string, len(ip))
	for i, b := range ip {
		// parts[i] = fmt.Sprintf("%d", b)
		parts[i] = strconv.Itoa(int(b))
	}
	return strings.Join(parts, ".")
}

func TryExcerciseStringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

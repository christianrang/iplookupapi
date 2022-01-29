package main

import (
	"net"
)

func validateIP(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	return false
}

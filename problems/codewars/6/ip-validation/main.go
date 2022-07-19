package main

import (
	"net"
)

func Is_valid_ip(ip string) bool {
	if v := net.ParseIP(ip); v == nil {
		return false
	}
	return true
}

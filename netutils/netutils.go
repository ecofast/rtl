// Copyright 2016~2017 ecofast(无尽愿). All rights reserved.
// Use of this source code is governed by a BSD-style license.

// Package netutils implements some useful network utility functions.
package netutils

import (
	"net"
	"strings"

	. "github.com/ecofast/rtl/sysutils"
)

// [fe80::a4fe:15b9:8b02:8122%本地连接]:62787 ==> fe80::a4fe:15b9:8b02:8122
// 192.168.1.122:5000 ==> 192.168.1.122
func ipFromAddrStr(s string) string {
	i := strings.LastIndex(s, ":")
	s = s[:i]
	l := len(s)
	if s[0] == '[' && s[l-1] == ']' {
		s = s[1 : l-1]
		i = strings.LastIndex(s, "%")
		if i > 0 {
			s = s[0:i]
		}
	}
	return s
}

func portFromAddrStr(s string) uint16 {
	i := strings.LastIndex(s, ":")
	return uint16(StrToInt(s[i+1:]))
}

func IPFromNetAddr(addr net.Addr) string {
	s := addr.String()
	return ipFromAddrStr(s)
}

func PortFromNetAddr(addr net.Addr) uint16 {
	s := addr.String()
	return portFromAddrStr(s)
}

func IPPortFromNetAddr(addr net.Addr) (string, uint16) {
	s := addr.String()
	return ipFromAddrStr(s), portFromAddrStr(s)
}

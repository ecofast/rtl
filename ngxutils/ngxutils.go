// Copyright 2017~2018 The NginX Authors / ecofast(无尽愿). All rights reserved.
// Use of this source code is governed by a BSD-style license.

// Package ngxutils rewrite some useful utility functions from Nginx in Golang.
package ngxutils

/*
 * 寻找大于等于 d 的第一个 a 的倍数(a 须是 2 的幂次)
 * #define ngx_align(d, a)     (((d) + (a - 1)) & ~(a - 1))
 */
func NgxAlign(d int, a int) int {
	return (d + (a - 1)) & (^(a - 1))
}

// 取出 v 的低 5 位的值, 相当于取模 32
func NgxMod32(v int) int {
	return v & 0x1f
}

// 获得一个数字的二进制里最低位的 1 的位置
func NgxLowestOneBit(v int) int {
	return v & (^(v - 1))
}

// 获得一个数字的二进制里最高位的 1 的位置
func NgxHighestOneBit(v int) int {
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32

	return v - (v >> 1)
}

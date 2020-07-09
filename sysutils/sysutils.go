// Copyright 2016~2017 ecofast(无尽愿). All rights reserved.
// Use of this source code is governed by a BSD-style license.

// Package sysutils implements some useful system utility functions.
package sysutils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func BoolToStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func StrToBool(s string) bool {
	if ret, err := strconv.ParseBool(s); err == nil {
		return ret
	}
	return false
}

func IntToStr(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func UIntToStr(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

func StrToInt(s string) int {
	if ret, err := strconv.Atoi(s); err == nil {
		return ret
	}
	panic(fmt.Sprintf("Cannot convert %s to int!", s))
}

func StrToIntDef(s string, defaultValue int) int {
	if ret, err := strconv.Atoi(s); err == nil {
		return ret
	}
	return defaultValue
}

func StrToUInt(s string) uint {
	if ret, err := strconv.ParseUint(s, 10, 0); err == nil {
		return uint(ret)
	}
	panic(fmt.Sprintf("Cannot convert %s to uint!", s))
}

func StrToUIntDef(s string, defaultValue uint) uint {
	if ret, err := strconv.ParseUint(s, 10, 0); err == nil {
		return uint(ret)
	}
	return defaultValue
}

func FloatToStr(f float32) string {
	return DoubleToStrPrec(float64(f), 7)
}

func DoubleToStr(f float64) string {
	return DoubleToStrPrec(f, 16)
}

func FloatToStrPrec(f float32, prec int) string {
	return DoubleToStrPrec(float64(f), prec)
}

func DoubleToStrPrec(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

func StrToFloat(s string) float32 {
	return float32(StrToDouble(s))
}

func StrToDouble(s string) float64 {
	if ret, err := strconv.ParseFloat(s, 64); err == nil {
		return ret
	}
	panic(fmt.Sprintf("Cannot convert %s to floating-point number!", s))
}

func BytesToUInt8(bs []byte) uint8 {
	return uint8(bs[0])
}

func BytesToUInt16(bs []byte) uint16 {
	return binary.LittleEndian.Uint16(bs)
}

func BytesToUInt32(bs []byte) uint32 {
	return binary.LittleEndian.Uint32(bs)
}

func UInt16ToBytes(v uint16) []byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, v)
	return bs
}

func Int16ToBytes(v int16) []byte {
	return UInt16ToBytes(uint16(v))
}

func UInt32ToBytes(v uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, v)
	return bs
}

func Int32ToBytes(v int32) []byte {
	return UInt32ToBytes(uint32(v))
}

func UInt64ToBytes(v uint64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, v)
	return bs
}

func Int64ToBytes(v int64) []byte {
	return UInt64ToBytes(uint64(v))
}

func GetApplicationPath() string {
	path := filepath.Dir(os.Args[0])
	return path + string(os.PathSeparator)
}

func DirectoryExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err == nil && fileInfo.IsDir() {
		return true
	}
	return false
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func CreateFile(fileName string) bool {
	_, err := os.Create(fileName)
	if err == nil {
		return true
	}
	return false
}

func IncludeTrailingBackslash(path string) string {
	if !strings.HasSuffix(path, string(os.PathSeparator)) {
		return path + string(os.PathSeparator)
	}
	return path
}

func ChangeFileExt(fileName, ext string) string {
	i := strings.LastIndex(fileName, ".")
	if i > 0 {
		ret := fileName[:i]
		return ret + ext
	}
	return fileName + ext
}

func GetFileExt(fileName string) string {
	i := strings.LastIndex(fileName, ".")
	if i > 0 {
		return fileName[i:]
	}
	return ""
}

func StructToBytes(data interface{}) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data)
	return buf.Bytes()
}

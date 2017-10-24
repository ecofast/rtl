// Copyright 2017 ecofast. All rights reserved.
// Use of this source code is governed by a BSD-style license.

// Package iniutils implements some useful ini-related utility functions.
package iniutils

import (
	"github.com/ecofast/rtl/inifiles"
)

func IniReadString(fileName, section, ident, defaultValue string) string {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	return inifile.ReadString(section, ident, defaultValue)
}

func IniWriteString(fileName, section, ident, value string) {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	inifile.WriteString(section, ident, value)
}

func IniReadInt(fileName, section, ident string, defaultValue int) int {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	return inifile.ReadInt(section, ident, defaultValue)
}

func IniWriteInt(fileName, section, ident string, value int) {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	inifile.WriteInt(section, ident, value)
}

func IniSectionExists(fileName, section string) bool {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	return inifile.SectionExists(section)
}

func IniReadSectionIdents(fileName, section string) []string {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	return inifile.ReadSectionIdents(section)
}

func IniReadSections(fileName string) []string {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	return inifile.ReadSections()
}

func IniReadSectionValues(fileName, section string) []string {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	return inifile.ReadSectionValues(section)
}

func IniEraseSection(fileName, section string) {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	inifile.EraseSection(section)
}

func IniIdentExists(fileName, section, ident string) bool {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	return inifile.IdentExists(section, ident)
}

func IniDeleteIdent(fileName, section, ident string) {
	inifile := inifiles.New(fileName, false)
	defer inifile.Close()
	inifile.DeleteIdent(section, ident)
}

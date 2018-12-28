// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package initfunc

import (
	"errors"
	"io"
	"strconv"
	"strings"
	"unicode"

	"github.com/allenhaozi/alog/writers/rotate"
)

var invalidByteQuantityError = errors.New("byte quantity must be a positive integer with a unit of measurement like M, MB, MiB, G, GiB, or GB")

const (
	BYTE = 1 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
)

// Rotate 的初始化函数。
func Rotate(args map[string]string) (io.Writer, error) {
	format, found := args["filename"]
	if !found {
		return nil, argNotFoundErr("rotate", "filename")
	}

	dir, found := args["dir"]
	if !found {
		return nil, argNotFoundErr("rotate", "dir")
	}

	sizeStr, found := args["size"]
	if !found {
		return nil, argNotFoundErr("rotate", "size")
	}

	size, err := toByte(sizeStr)
	if err != nil {
		return nil, err
	}

	return rotate.New(format, dir, size)
}

// ToBytes parses a string formatted by ByteSize as bytes. Note binary-prefixed and SI prefixed units both mean a base-2 units
// KB = K = KiB = 1024
// MB = M = MiB = 1024 * K
// GB = G = GiB = 1024 * M
// TB = T = TiB = 1024 * G
func toByte(s string) (int64, error) {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	i := strings.IndexFunc(s, unicode.IsLetter)

	if i == -1 {
		return 0, invalidByteQuantityError
	}

	bytesString, multiple := s[:i], s[i:]
	bytes, err := strconv.ParseFloat(bytesString, 64)
	if err != nil || bytes <= 0 {
		return 0, invalidByteQuantityError
	}

	switch multiple {
	case "T", "TB", "TIB":
		return int64(bytes * TERABYTE), nil
	case "G", "GB", "GIB":
		return int64(bytes * GIGABYTE), nil
	case "M", "MB", "MIB":
		return int64(bytes * MEGABYTE), nil
	case "K", "KB", "KIB":
		return int64(bytes * KILOBYTE), nil
	case "B":
		return int64(bytes), nil
	default:
		return 0, invalidByteQuantityError
	}
}

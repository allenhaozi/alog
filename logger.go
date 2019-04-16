// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package alog

import (
	"errors"
	"io"
	"log"
	"strings"

	"github.com/allenhaozi/alog/writers"
)

var flagMap = map[string]int{
	"none":              0,
	"log.ldate":         log.Ldate,
	"log.ltime":         log.Ltime,
	"log.lmicroseconds": log.Lmicroseconds,
	"log.llongfile":     log.Llongfile,
	"log.lshortfile":    log.Lshortfile,
	"log.lstdflags":     log.LstdFlags,
}

type logger struct {
	flush writers.Flusher // log is container of io.Writer
	log   *log.Logger     // make sure log is not nil
}

func (l *logger) set(w io.Writer, prefix string, flag int) {
	if w == nil {
		l.flush = nil
		l.log = discard
		return
	}

	if f, ok := w.(writers.Flusher); ok {
		l.flush = f
	}
	l.log = log.New(w, prefix, flag)
}

func loggerInitializer(args map[string]string) (io.Writer, error) {
	return writers.NewContainer(), nil
}

// 将 log.Ldate|log.Ltime 的字符串转换成正确的值
func parseFlag(flagStr string) (int, error) {
	flagStr = strings.TrimSpace(flagStr)
	if len(flagStr) == 0 {
		return 0, nil
	}

	strs := strings.Split(flagStr, "|")
	ret := 0

	for _, str := range strs {
		str = strings.ToLower(strings.TrimSpace(str))
		flag, found := flagMap[str]
		if !found {
			return 0, errors.New("无效的 flag:" + str)
		}
		ret |= flag
	}

	return ret, nil
}

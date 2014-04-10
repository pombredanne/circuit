// Copyright 2013 The Go Circuit Project
// Use of this source code is governed by the license for
// The Go Circuit Project, found in the LICENSE file.
//
// Authors:
//   2014 Petar Maymounkov <p@gocircuit.org>

package file

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/gocircuit/circuit/kit/iomisc"
	"github.com/gocircuit/circuit/kit/fs/rh"
)

func NewErrorFile() *ErrorFile {
	return &ErrorFile{}
}

type ErrorFile struct {
	sync.Mutex
	msg string
}

func (f *ErrorFile) Clear() {
	f.Set("")
}

func (f *ErrorFile) Set(msg string) {
	f.Lock()
	defer f.Unlock()
	f.msg = msg
}

func (f *ErrorFile) Setf(format string, arg ...interface{}) {
	f.Set(fmt.Sprintf(format, arg...))
}

func (f *ErrorFile) Get() string {
	f.Lock()
	defer f.Unlock()
	return f.msg
}

func (f *ErrorFile) Perm() rh.Perm {
	return 0444 // r--r--r--
}

func (f *ErrorFile) Open(rh.Flag, rh.Intr) (rh.FID, error) {
	return NewOpenReaderFile(iomisc.ReaderNopCloser(bytes.NewBufferString(f.Get()))), nil
}

func (f *ErrorFile) Remove() error {
	return rh.ErrPerm
}

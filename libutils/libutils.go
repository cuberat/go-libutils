// BSD 2-Clause License
//
// Copyright (c) 2018-2019 Don Owens <don@regexguy.com>.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// The libutils package provides various utilities for working in Go.
//
// Installation
//
//   go get github.com/cuberat/go-libutils/libutils
package libutils

import (
    "errors"
    "fmt"
    // "io"
    "os"
    "path"
    "runtime"
)

var (
    UnknownSuffix error = errors.New("Unknown suffix")
    VarintNotEnoughBytes error = errors.New("Not enough bytes in varint")
)

const (
    Version = "1.03"
)

// Like fmt.Errorf(), except adds the (base) file name and line number to the
// beginning of the error message in the format `[%s:%d] `.
func Errorf(fmt_str string, args ...interface{}) error {
    _, file_name, line, ok := runtime.Caller(1)

    err_str := fmt.Sprintf(fmt_str, args...)

    if ok {
        file_name = path.Base(file_name)
        err_str = fmt.Sprintf("[%s:%d] ", file_name, line) + err_str

    }

    return errors.New(err_str)
}

// Like fmt.Errorf(), except adds the full file name and line number to the
// beginning of the error message in the format `[%s:%d] `.
func ErrorfLong(fmt_str string, args ...interface{}) error {
    _, file_name, line, ok := runtime.Caller(1)

    err_str := fmt.Sprintf(fmt_str, args...)

    if ok {
        err_str = fmt.Sprintf("[%s:%d] ", file_name, line) + err_str

    }

    return errors.New(err_str)
}

func find_exec(file string) (string, error) {
    dirs := []string{"/bin", "/usr/bin", "/usr/local/bin"}

    for _, dir := range dirs {
        path := fmt.Sprintf("%s/%s", dir, file)
        _, err := os.Stat(path)
        if err == nil {
            return path, nil
        }
    }

    return "", fmt.Errorf("couldn't find executable %s", file)
}

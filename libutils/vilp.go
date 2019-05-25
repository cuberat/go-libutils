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

package libutils

import (
    "bufio"
    "fmt"
    "io"
)

// Returns a byte slice with a varint length prefix followed by the provided
// byte slice.
func BytesToVILP(data []byte) ([]byte) {
    prefix := EncodeVarint(uint64(len(data)))
    val := make([]byte, 0, len(prefix) + len(data))
    val = append(val, prefix...)
    val = append(val, data...)

    return val
}

// VILPWriter is used to write length-prefixed strings to an io.Writer
type VILPWriter struct {
    w io.Writer
}

// Writes the provided string as a length-prefixed string to the
// underlying io.Writer
func (plw *VILPWriter)WriteString(s string) (int, error) {
    prefix := EncodeVarint(uint64(len(s)))
    return io.WriteString(plw.w, string(prefix) + s)
}

// Writes the provided bytes as a length-prefixed string to the
// underlying io.Writer
func (plw *VILPWriter)Write(p []byte) (int, error) {
    return plw.w.Write(BytesToVILP(p))
}

// Returns a new VILPWriter. VILPWriter implements the
// io.Writer interface, in addition to the WriteString method.
func NewVILPWriter(w io.Writer) (*VILPWriter) {
    plw := new(VILPWriter)
    plw.w = w

    return plw
}

// Opens the provided file and returns a *VILPWriter created using the
// resulting file handle. Call close_func() to close the underlying file handle.
func NewVILPWriterF(file_path string) (*VILPWriter, CloseFunc,
    error) {
    w, close_func, err := OpenFileW(file_path)
    if err != nil {
        return nil, nil, err
    }

    return NewVILPWriter(w), close_func, nil
}

// Returns a bufio.Scanner that scans varint length-prefixed strings from the
// provided io.Reader.
func NewVILPScanner(r io.Reader) (*bufio.Scanner) {
    scanner := bufio.NewScanner(r)
    scanner.Split(ScannerVILPScan)

    return scanner
}

// Returns a bufio.Scanner that scans varint length-prefixed strings from the
// provided file. Call close_func() to close the underlying file handle.
func NewVILPScannerF(file_path string) (*bufio.Scanner,
    CloseFunc, error) {

    r, close_func, err := OpenFileRO(file_path)
    if err != nil {
        return nil, nil, err
    }

    scanner := NewVILPScanner(r)

    return scanner, close_func, nil
}

// A bufio.SplitFunc that reads length-prefixed strings from a reader.
func ScannerVILPScan(data []byte, at_eof bool) (int, []byte, error) {
    if len(data) == 0 {
        return 0, nil, nil
    }

    prefix_len, cnt, err := DecodeVarint(data)
    if err != nil {
        if err == VarintNotEnoughBytes {
            if at_eof {
                return len(data), nil, fmt.Errorf("invalid format")
            }
        }
        return 0, nil, err
    }

    needed_len := prefix_len + uint64(cnt)
    if uint64(len(data)) < needed_len {
        return 0, nil, nil
    }

    return int(needed_len), data[cnt:needed_len], nil
}

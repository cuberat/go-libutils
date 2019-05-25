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

// PrefixLenWriter is used to write length-prefixed strings to an io.Writer
//
// Deprecated: use VILPWriter and its corresponding methods.
type PrefixLenWriter struct {
    w io.Writer
}

// Writes the provided string as a length-prefixed string to the
// underlying io.Writer. This uses 32-bit integers for the length prefix.
//
// Deprecated: use VILPWriter and its corresponding methods.
func (plw *PrefixLenWriter)WriteString(s string) (int, error) {
    return io.WriteString(plw.w, s)
}

// Writes the provided bytes as a length-prefixed string to the
// underlying io.Writer. This uses 32-bit integers for the length prefix.
//
// Deprecated: use VILPWriter and its corresponding methods.
func (plw *PrefixLenWriter)Write(p []byte) (int, error) {
    prefix_len := uint(len(p))
    len_bytes := make([]byte, 4)

    len_bytes[0] = byte((prefix_len >> 24) & 0xff)
    len_bytes[1] = byte((prefix_len >> 16) & 0xff)
    len_bytes[2] = byte((prefix_len >> 8) & 0xff)
    len_bytes[3] = byte(prefix_len & 0xff)

    n, err := plw.w.Write(len_bytes)
    if err != nil {
        return n, err
    }

    n2, err := plw.w.Write(p)

    return n + n2, err
}

// Returns a new PrefixLenWriter. PrefixLenWriter implements the
// io.Writer interface, in addition to the WriteString method.
//
// Deprecated: use VILPWriter and its corresponding methods.
func NewPrefixLenWriter(w io.Writer) (*PrefixLenWriter) {
    plw := new(PrefixLenWriter)
    plw.w = w

    return plw
}

// Returns a bufio.Scanner that scans length-prefixed strings from the
// provided io.Reader.
//
// Deprecated: use NewVILPScanner and varint length-prefixed files.
func NewPrefixLenScanner(r io.Reader) (*bufio.Scanner) {
    scanner := bufio.NewScanner(r)
    scanner.Split(ScannerPrefixLenScan)

    return scanner
}

// A bufio.SplitFunc that reads length-prefixed strings from a reader
//
// Deprecated: use NewVILPScanner and varint length-prefixed files.
func ScannerPrefixLenScan(data []byte, at_eof bool) (int, []byte, error) {
    if len(data) < 4 {
        if at_eof {
            return len(data), nil, fmt.Errorf("invalid format")
        }

        return 0, nil, nil
    }

    prefix_len := uint(0)
    prefix_len += (uint(data[0]) << 24) + (uint(data[1]) << 16) +
        (uint(data[2]) << 8) + (uint(data[3]))

    needed_len := prefix_len + 4
    if uint(len(data)) < needed_len {
        return 0, nil, nil
    }

    return int(needed_len), data[4:needed_len], nil
}

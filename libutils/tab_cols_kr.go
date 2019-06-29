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
    "bytes"
    "fmt"
    "io"
)

// Implements the KeyedRecordEncoder and KeyedRecordDecoder interfaces specified
// by `github.com/cuberat/go-libutils/libutils`.
//
//
// This codec serializes and deserializes keyed records where the key is the
// first tab-separated column, and the value is a slice of byte slices
// containing the remaining columns. This is not TSV, where tab bytes can be
// escaped, etc. It is a simple split on tab bytes. That is,
//
//    <key>[tab]<val1>[tab]<val2>...
type TabColsKRCodec struct {

}

// Returns a new TabColsKRCodec
func NewTabColsKRCodec () (*TabColsKRCodec) {
    return new(TabColsKRCodec)
}

// Splits the record, returning the key and the serialized value data
// structure.
func (krc *TabColsKRCodec) SplitKV(wire_data []byte) ([]byte, []byte,
    error) {

    kv := bytes.SplitN(wire_data, []byte{'\t'}, 2)
    key := kv[0]

    if len(kv) < 2 {
        return key, make([]byte, 0, 0), nil
    }

    return key, kv[1], nil
}

// Deserializes the value.
func (krc *TabColsKRCodec) UnmarshalVal(val_bytes []byte) (interface{},
    error) {

    return bytes.Split(val_bytes, []byte{'\t'}), nil
}

// Joins the key and value bytes, returning the serialized record.
func (krc *TabColsKRCodec) JoinKV(key, val []byte) ([]byte, error) {
    return bytes.Join([][]byte{key, val}, []byte{'\t'}), nil
}

// Serializes the value data structure.
func (krc *TabColsKRCodec) MarshalVal(data interface{}) ([]byte, error) {
    val_slice, ok := data.([][]byte)
    if !ok {
        return nil, fmt.Errorf("MarshalVal(): expected [][]byte, got %T", data)
    }

    return bytes.Join(val_slice, []byte{'\t'}), nil
}

// Returns true so that if this codec is used for both encoder and decoder,
// unnecessary re-serialization can be avoided.
//
// This allows for lazy encoding. That is, if the raw record bytes that were
// read in do not need to change, they can be written back out as-is, instead of
// actually re-encoding.
func (krc *TabColsKRCodec) CodecSame() bool {
    return true
}

// Implements the `libutils.KeyedRecordWriter` interface from
// `github.com/cuberat/go-libutils/libutils`. Records look like
//
//    <key>[tab]<val1>[tab]<val2>...[newline]
type TabColsKRWriter struct {
    encoder KeyedRecordEncoder
    writer io.Writer
}

// Returns a new TabColsKRWriter.
func NewTabColsKRWriter (w io.Writer) (*TabColsKRWriter) {
    krw := new(TabColsKRWriter)
    krw.writer = w
    krw.encoder = NewTabColsKRCodec()

    return krw
}

// The Write method for the KeyedRecordWriter interface
func (krw *TabColsKRWriter) Write(rec *KeyedRecord) (int, error) {
    rec_out_bytes, err := rec.RecordBytesOut(krw.encoder)
    if err != nil {
        return 0, err
    }

    return fmt.Fprintf(krw.writer, "%s\n", rec_out_bytes)
}

// Implements the `libutils.KeyedRecordScanner` interface from
// `github.com/cuberat/go-libutils/libutils`.
// This is a scanner for tab-delimited keyed records with column-based values.
// Records are expected to look like
//
//    <key>[tab]<val1>[tab]<val2>...[newline]
type TabColsKRScanner struct {
    decoder KeyedRecordDecoder
    scanner *bufio.Scanner
}

// Return a new TabColsKRScanner
func NewTabColsKRScanner(r io.Reader) (*TabColsKRScanner) {
    krs := new(TabColsKRScanner)
    krs.scanner = bufio.NewScanner(r)
    krs.decoder = NewTabColsKRCodec()

    return krs
}

// Advances the scanner to the next record. It returns false when the scan
// stops, either by reaching the end of the input or an error.
func (krs *TabColsKRScanner) Scan() bool {
    return krs.scanner.Scan()
}

// Returns the most recent serialized record generated by a call to Scan().
func (krs *TabColsKRScanner) Record() (*KeyedRecord) {
    wire_data := krs.scanner.Bytes()
    wire_data_copy := make([]byte, len(wire_data))
    copy(wire_data_copy, wire_data)

    return NewKeyedRecordFromBytes(wire_data_copy, krs.decoder)
}

// Returns the first non-EOF error that was encountered by the Scanner.
func (krs *TabColsKRScanner) Err() error {
    return krs.scanner.Err()
}

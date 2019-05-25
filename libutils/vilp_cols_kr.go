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

// Implements the KeyedRecordEncoder and KeyedRecordDecoder interfaces specified
// by `github.com/cuberat/go-libutils/libutils`.
//
// This codec serializes and deserializes keyed records where the is
// length-prefixed using a varint. The value immediately follows and consists of
// length-prefixed strings represented an array/list of strings. That is,
//
//    <key_len><key><val1_len><val1><val2_len><val2>...
type VILPColsKRCodec struct {

}

func NewVILPColsKRCodec () (*VILPColsKRCodec) {
    return new(VILPColsKRCodec)
}

// Splits the record, returning the key and the serialized value data
// structure.
func (krc *VILPColsKRCodec) SplitKV(wire_data []byte) ([]byte, []byte,
    error) {
    key_len, vi_len, err := DecodeVarint(wire_data)
    if err != nil {
        return nil, nil, fmt.Errorf("couldn't decode varint in SplitKV(): %s",
            err)
    }

    if uint64(len(wire_data)) < uint64(vi_len) + key_len {
        return []byte{}, []byte{},
        fmt.Errorf("wire_data too short in SplitKV()")
    }

    data := wire_data[vi_len:]

    key := data[:key_len]
    val := data[key_len:]

    return key, val, nil
}

// Deserializes the value.
func (krc *VILPColsKRCodec) UnmarshalVal(val_bytes []byte) (interface{},
    error) {
    rest_val := val_bytes

    values := make([][]byte, 0)
    for len(rest_val) > 0 {
        val_len, vi_len, err := DecodeVarint(rest_val)
        if err != nil {
            return nil, fmt.Errorf("bad varint while parsing value")
        }
        if uint64(len(rest_val)) < uint64(vi_len) + val_len {
            return nil, fmt.Errorf("data to short while parsing value")
        }
        rest_val = rest_val[vi_len:]
        val := rest_val[:val_len]
        values = append(values, val)
        rest_val = rest_val[val_len:]
    }

    return values, nil
}

// Joins the key and value bytes, returning the serialized record.
func (krc *VILPColsKRCodec) JoinKV(key, val []byte) ([]byte, error) {
    key_len_bytes := EncodeVarint(uint64(len(key)))
    rec := make([]byte, 0, len(key_len_bytes) + len(key) + len(val))
    rec = append(rec, key_len_bytes...)
    rec = append(rec, key...)
    rec = append(rec, val...)

    return rec, nil
}

// Serializes the value data structure.
func (krc *VILPColsKRCodec) MarshalVal(data interface{}) ([]byte, error) {
    val_slice, ok := data.([][]byte)
    if !ok {
        return nil, fmt.Errorf("MarshalVal(): expected [][]byte, got %T", data)
    }

    val_bytes := make([]byte, 0)
    for _, val := range val_slice {
        prefix := EncodeVarint(uint64(len(val)))
        val_bytes = append(val_bytes, prefix...)
        val_bytes = append(val_bytes, val...)
    }

    return val_bytes, nil
}

// Returns true so that if this codec is used for both encoder and decoder,
// unnecessary re-serialization can be avoided.
//
// This allows for lazy encoding. That is, if the raw record bytes that were
// read in do not need to change, they can be written back out as-is, instead of
// actually re-encoding.
func (krc *VILPColsKRCodec) CodecSame() bool {
    return true
}

// Implements the `libutils.KeyedRecordWriter` interface from
// `github.com/cuberat/go-libutils/libutils`.
type VILPColsKRWriter struct {
    encoder KeyedRecordEncoder
    vilp_writer *VILPWriter
}

// Returns a new VILPColsKRWriter.
func NewVILPColsKRWriter (w io.Writer) (*VILPColsKRWriter) {
    krw := new(VILPColsKRWriter)
    krw.vilp_writer = NewVILPWriter(w)
    krw.encoder = NewVILPColsKRCodec()

    return krw
}

func (krw *VILPColsKRWriter) Write(rec *KeyedRecord) (int, error) {
    rec_out_bytes, err := rec.RecordBytesOut(krw.encoder)
    if err != nil {
        return 0, err
    }

    return krw.vilp_writer.Write(rec_out_bytes)
}

// Scanner for varint length-prefixed keyed records with column-based values.
// That is, the value is an array/list of strings, each one prefixed with a
// length.
type VILPColsKRScanner struct {
    decoder KeyedRecordDecoder
    scanner *bufio.Scanner
}

// Return a new VILPColsKRScanner
func NewVILPColsKRScanner(r io.Reader) (*VILPColsKRScanner) {
    krs := new(VILPColsKRScanner)
    krs.scanner = NewVILPScanner(r)
    krs.decoder = NewVILPColsKRCodec()

    return krs
}

// Advances the scanner to the next record. It returns false when the scan
// stops, either by reaching the end of the input or an error.
func (krs *VILPColsKRScanner) Scan() bool {
    return krs.scanner.Scan()
}

// Returns the most recent serialized record generated by a call to Scan().
func (krs *VILPColsKRScanner) Record() (*KeyedRecord) {
    wire_data := krs.scanner.Bytes()
    wire_data_copy := make([]byte, len(wire_data))
    copy(wire_data_copy, wire_data)

    return NewKeyedRecordFromBytes(wire_data_copy, krs.decoder)
}

// Returns the first non-EOF error that was encountered by the Scanner.
func (krs *VILPColsKRScanner) Err() error {
    return krs.scanner.Err()
}

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
    "fmt"
)

// Decodes a varint (as used in protobuffers) into a uint64.
// See https://developers.google.com/protocol-buffers/docs/encoding#varints
// for the specification.
func DecodeVarint(data_in []byte) (uint64, int, error) {
    cnt := 0
    max_cnt := 10 // for 64-bit int

    val := uint64(0)
    for i, b := range data_in {
        cnt++
        if cnt > max_cnt {
            return 0, 0, fmt.Errorf("invalid varint encoding for 64-bit int")
        }

        val += ((uint64(b) & 0x7f) << uint(7 * i))

        if (b & 0x80) == 0 {
            break
        }
    }

    return val, cnt, nil
}

// Encodes a uint64 as a varint (as used in protobuffers).
// See https://developers.google.com/protocol-buffers/docs/encoding#varints
// for the specification.
func EncodeVarint(int_in uint64) ([]byte) {
    data := make([]byte, 0, 10)
    last_non_zero := 0
    for i := 0; i < 10; i++ {
        b := (int_in >> uint(i * 7)) & 0x7f
        if b != 0 {
            last_non_zero = i
        }
        b |= 0x80
        data = append(data, byte(b))
    }
    data[last_non_zero] &= 0x7f

    return data[:last_non_zero + 1]
}

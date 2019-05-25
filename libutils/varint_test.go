package libutils_test

import (
    "bytes"
    libutils "github.com/cuberat/go-libutils/libutils"
    "testing"
)

func TestVarintEncode(t *testing.T) {
    in_vals := []uint64{1, 2, 15, 16, 30, 32, 62, 64, 100, 200, 555,
        1024, 1234567, 23456057, 4611686018427387939, 18446744073709551615}
    encoded_vals :=
        [][]byte{
            []byte{1},
            []byte{2},
            []byte{15},
            []byte{16},
            []byte{30},
            []byte{32},
            []byte{62},
            []byte{64},
            []byte{100},
            []byte{200,1},
            []byte{171,4},
            []byte{128,8},
            []byte{135,173,75},
            []byte{185,210,151,11},
            []byte{163,128,128,128,128,128,128,128,64},
            []byte{255,255,255,255,255,255,255,255,255,1},
        }

    for i, in_val := range in_vals {
        vi := libutils.EncodeVarint(in_val)
        if ! bytes.Equal(vi, encoded_vals[i]) {
            t.Errorf("Encoded varint incorrect for %d: got %#v, expected %#v",
                in_val, vi, encoded_vals[i])
        }
    }
}

func TestVarintDecode(t *testing.T) {
    in_vals := []uint64{1, 2, 15, 16, 30, 32, 62, 64, 100, 200, 555,
        1024, 1234567, 23456057, 4611686018427387939, 18446744073709551615}
    encoded_vals :=
        [][]byte{
            []byte{1},
            []byte{2},
            []byte{15},
            []byte{16},
            []byte{30},
            []byte{32},
            []byte{62},
            []byte{64},
            []byte{100},
            []byte{200,1},
            []byte{171,4},
            []byte{128,8},
            []byte{135,173,75},
            []byte{185,210,151,11},
            []byte{163,128,128,128,128,128,128,128,64},
            []byte{255,255,255,255,255,255,255,255,255,1},
        }

    for i, enc_val := range encoded_vals {
        int_val, _, err := libutils.DecodeVarint(enc_val)
        if err != nil {
            t.Errorf("Error decoding encoded var int %#v: %s", enc_val, err)
        }
        if int_val != in_vals[i] {
            t.Errorf("Decoded varint incorrect for %#v: got %d, expected %d",
                enc_val, int_val, in_vals[i])
        }
    }
}

package libutils_test

import (
    "bytes"
    libutils "github.com/cuberat/go-libutils/libutils"
    "testing"
)

func TestVILPWriter(t *testing.T) {
    test_strings := []string{"foo", "bar", "something longer", "stuff 😜"}
    out_bytes := []byte{0x3, 0x66, 0x6f, 0x6f, 0x3, 0x62, 0x61, 0x72, 0x10,
        0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x6c, 0x6f,
        0x6e, 0x67, 0x65, 0x72, 0xa, 0x73, 0x74, 0x75, 0x66, 0x66, 0x20, 0xf0,
        0x9f, 0x98, 0x9c}

    out_fh := new(bytes.Buffer)

    w := libutils.NewVILPWriter(out_fh)
    for _, str := range test_strings {
        w.Write([]byte(str))
    }

    if ! bytes.Equal(out_fh.Bytes(), out_bytes) {
        t.Errorf("Error writing VILP strings:\ngot:\n%#v\nexpected:\n%#v",
            out_fh.Bytes(), out_bytes)
    }
}

func TestVILPScanner(t *testing.T) {
    test_strings := []string{"foo", "bar", "something longer", "stuff 😜"}
    enc_bytes := []byte{0x3, 0x66, 0x6f, 0x6f, 0x3, 0x62, 0x61, 0x72, 0x10,
        0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x6c, 0x6f,
        0x6e, 0x67, 0x65, 0x72, 0xa, 0x73, 0x74, 0x75, 0x66, 0x66, 0x20, 0xf0,
        0x9f, 0x98, 0x9c}

    r := bytes.NewBuffer(enc_bytes)
    scanner := libutils.NewVILPScanner(r)

    i := 0
    for scanner.Scan() {
        b := scanner.Bytes()
        if string(b) != test_strings[i] {
            t.Errorf("Error with VILPScanner: got %#v, expected %#v\n",
                b, []byte(test_strings[i]))
        }
        i++
    }
}

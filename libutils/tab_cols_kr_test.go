package libutils_test

import (
    "bytes"
    libutils "github.com/cuberat/go-libutils/libutils"
    "testing"
)

func TestTabColsScanner(t *testing.T) {
    wire_data := []byte("foo\tbar\ted\tdead\tbeef\n")
    key_expected := []byte("foo")
    val_expected := [][]byte{[]byte("bar"), []byte("ed"), []byte("dead"),
        []byte("beef")}

    r := bytes.NewBuffer(wire_data)
    scanner := libutils.NewTabColsKRScanner(r)

    if !scanner.Scan() {
        t.Errorf("no records to scan")
        return
    }

    kr := scanner.Record()
    key, err := kr.Key()
    if err != nil {
        t.Errorf("couldn't get key: %s", err)
        return
    }

    if !bytes.Equal(key, key_expected) {
        t.Errorf("incorrect key: expected %q, got %q", key_expected, key)
        return
    }

    val_int, err := kr.Val()
    if err != nil {
        t.Errorf("couldn't get val: %s", err)
        return
    }

    val, ok := val_int.([][]byte)
    if !ok {
        t.Errorf("val not a [][]byte")
        return
    }

    if len(val) != len(val_expected) {
        t.Errorf("val and val_expected not the same length: %d vs %d",
            len(val), len(val_expected))
        return
    }

    for i, word := range val {
        if !bytes.Equal(word, val_expected[i]) {
            t.Errorf("val != val_expected: value at index %d differs: " +
                "%q vs %q", i, word, val_expected[i])
            return
        }
    }
}

func TestTabColsKRWriter(t *testing.T) {
    w := new(bytes.Buffer)
    writer := libutils.NewTabColsKRWriter(w)

    bytes_expected := []byte("foo\tbar\ted\tdead\tbeef\n")
    test_val := [][]byte{[]byte("bar"), []byte("ed"), []byte("dead"),
        []byte("beef")}
    kr := libutils.NewKeyedRecordFromKeyVal([]byte("foo"), test_val)

    _, err := writer.Write(kr)
    if err != nil {
        t.Errorf("couldn't write keyed record: %s", err)
        return
    }

    if !bytes.Equal(w.Bytes(), bytes_expected) {
        t.Errorf("incorrect serialization: expected %q, got %q",
            bytes_expected, w.Bytes())
    }
}

func TestTabColsKRCodec(t *testing.T) {
    codec := libutils.NewTabColsKRCodec()

    test_val := [][]byte{[]byte("bar"), []byte("ed"), []byte("dead"),
        []byte("beef")}
    data_bytes := []byte("foo\tbar\ted\tdead\tbeef")
    kr := libutils.NewKeyedRecordFromBytes(data_bytes, codec)

    key, err := kr.Key()
    if err != nil {
        t.Errorf("couldn't get key: %s", err)
        return
    }
    key_expected := []byte("foo")
    if !bytes.Equal(key, key_expected) {
        t.Errorf("incorrect key: expected %q, got %q", key_expected, key)
        return
    }

    val_int, err := kr.Val()
    if err != nil {
        t.Errorf("couldn't get val: %s", err)
        return
    }

    val, ok := val_int.([][]byte)
    if !ok {
        t.Errorf("val not a [][]byte")
        return
    }

    if len(val) != len(test_val) {
        t.Errorf("val and test_val not the same length: %d vs %d",
            len(val), len(test_val))
        return
    }

    for i, word := range val {
        if !bytes.Equal(word, test_val[i]) {
            t.Errorf("val != test_val: value at index %d differs: " +
                "%q vs %q", i, word, test_val[i])
            return
        }
    }
}

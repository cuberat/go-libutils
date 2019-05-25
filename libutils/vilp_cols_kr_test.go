package libutils_test

import (
    "bytes"
    libutils "github.com/cuberat/go-libutils/libutils"
    "testing"
)

func TestVILPColsKRWriter(t *testing.T) {
    vilp_cols_krw_bytes := []byte{0x13, 0x3, 0x66, 0x6f, 0x6f, 0x4, 0x66, 0x6f,
        0x6f, 0x31, 0x4, 0x66, 0x6f, 0x6f, 0x32, 0x4, 0x66, 0x6f, 0x6f, 0x33,
        0x13, 0x3, 0x62, 0x61, 0x72, 0x4, 0x62, 0x61, 0x72, 0x31, 0x4, 0x62,
        0x61, 0x72, 0x32, 0x4, 0x62, 0x61, 0x72, 0x33}

    // records := make([]*KeyedRecord, 2)

    w := new(bytes.Buffer)
    cols_kr_w := libutils.NewVILPColsKRWriter(w)
    key := []byte("foo")
    val := make([][]byte, 0, 5)
    val = append(val, []byte("foo1"))
    val = append(val, []byte("foo2"))
    val = append(val, []byte("foo3"))
    rec := libutils.NewKeyedRecordFromKeyVal(key, val)
    cols_kr_w.Write(rec)
    // records = append(records, rec)

    key = []byte("bar")
    val = make([][]byte, 0, 5)
    val = append(val, []byte("bar1"))
    val = append(val, []byte("bar2"))
    val = append(val, []byte("bar3"))
    rec = libutils.NewKeyedRecordFromKeyVal(key, val)
    cols_kr_w.Write(rec)
    // records = append(records, rec)

    out_bytes := w.Bytes()
    if ! bytes.Equal(out_bytes, vilp_cols_krw_bytes) {
        t.Errorf("failed to write keyed records with VILPColsKRWriter: " +
            "got %#v, expected %#v", out_bytes, vilp_cols_krw_bytes)
    }
}

func TestVILPColsKRScanner(t *testing.T) {
    vilp_cols_krw_bytes := []byte{0x13, 0x3, 0x66, 0x6f, 0x6f, 0x4, 0x66, 0x6f,
        0x6f, 0x31, 0x4, 0x66, 0x6f, 0x6f, 0x32, 0x4, 0x66, 0x6f, 0x6f, 0x33,
        0x13, 0x3, 0x62, 0x61, 0x72, 0x4, 0x62, 0x61, 0x72, 0x31, 0x4, 0x62,
        0x61, 0x72, 0x32, 0x4, 0x62, 0x61, 0x72, 0x33}

    vilp_cols_krw_keys := [][]uint8{
        []uint8{0x66, 0x6f, 0x6f},
        []uint8{0x62, 0x61, 0x72}}
    vilp_cols_krw_vals := [][][]uint8{
        [][]uint8{
            []uint8{0x66, 0x6f, 0x6f, 0x31},
            []uint8{0x66, 0x6f, 0x6f, 0x32},
            []uint8{0x66, 0x6f, 0x6f, 0x33}},
        [][]uint8{
            []uint8{0x62, 0x61, 0x72, 0x31},
            []uint8{0x62, 0x61, 0x72, 0x32},
            []uint8{0x62, 0x61, 0x72, 0x33}}}

    r := bytes.NewBuffer(vilp_cols_krw_bytes)
    vilp_krs := libutils.NewVILPColsKRScanner(r)

    i := -1
    for vilp_krs.Scan() {
        i++
        rec := vilp_krs.Record()
        key, err := rec.Key()
        if err != nil {
            t.Errorf("couldn't get key from rec: %s", err)
            continue
        }

        if ! bytes.Equal(key, vilp_cols_krw_keys[i]) {
            t.Errorf("Error parsing key %d: got %#v, expected %#v", i, key,
                vilp_cols_krw_keys[i])
        }

        val_int, err := rec.Val()
        if err != nil {
            t.Errorf("couldn't get val from rec: %s", err)
            continue
        }
        byte_slices, ok := val_int.([][]byte)
        if !ok {
            t.Error("val return from rec is not a slice of byte slices")
        }

        for j, slice := range byte_slices {
            if ! bytes.Equal(slice, vilp_cols_krw_vals[i][j]) {
                t.Errorf("Error parsing value %d/%d: got %#v, expected %#v",
                i, j, slice, vilp_cols_krw_vals[i][j])
            }
        }
    }
}

package libutils_test

import (
    "fmt"
    libutils "github.com/cuberat/go-libutils/libutils"
    "strings"
    "testing"
)

func TestErrorf(t *testing.T) {
    err := libutils.Errorf("testing Errorf")
    err_str := fmt.Sprintf("%s", err)
    if strings.Contains(err_str, "/") {
        t.Errorf("err message should not contain a `/`")
    }
}

func TestErrorfLong(t *testing.T) {
    err := libutils.ErrorfLong("testing Errorf")
    err_str := fmt.Sprintf("%s", err)
    if !strings.Contains(err_str, "/") {
        t.Errorf("err message should contain a `/`")
    }
}

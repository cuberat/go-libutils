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

// The libutils package provides various utilities for working in Go.
//
// Installation
//
//   go get github.com/cuberat/go-libutils/libutils
package libutils

import (
    "compress/bzip2"
    "compress/gzip"
    "errors"
    "fmt"
    "io"
    "os"
    "os/exec"
    "path"
    "runtime"
    "strings"
)

var (
    UnknownSuffix error = errors.New("Unknown suffix")
    VarintNotEnoughBytes error = errors.New("Not enough bytes in varint")
)

const (
    Version = "1.02"
)

// Like fmt.Errorf(), except adds the (base) file name and line number to the
// beginning of the error message in the format `[%s:%d] `.
func Errorf(fmt_str string, args ...interface{}) error {
    _, file_name, line, ok := runtime.Caller(1)

    err_str := fmt.Sprintf(fmt_str, args...)

    if ok {
        file_name = path.Base(file_name)
        err_str = fmt.Sprintf("[%s:%d] ", file_name, line) + err_str

    }

    return errors.New(err_str)
}

// Like fmt.Errorf(), except adds the full file name and line number to the
// beginning of the error message in the format `[%s:%d] `.
func ErrorfLong(fmt_str string, args ...interface{}) error {
    _, file_name, line, ok := runtime.Caller(1)

    err_str := fmt.Sprintf(fmt_str, args...)

    if ok {
        err_str = fmt.Sprintf("[%s:%d] ", file_name, line) + err_str

    }

    return errors.New(err_str)
}

// Signature for Close() function return from OpenFileW and
// OpenFileRO. When ready to close the file, call the function to
// close and clean up.
type CloseFunc func() ()

// Open a file for writing. If the file name dends in a supported
// compression suffix, output will be compressed in that format.
//
// Supported compression:
//    gzip  (.gz)
//    bzip2 (.bz2) -- calls external program
//    xz    (.xz)  -- calls external program
//
// The CloseFunc object is a function that you must call to close the file
// properly.
func OpenFileW(outfile string) (io.Writer, CloseFunc, error) {
    out_fh, err := os.Create(outfile)
    if err != nil {
        return nil, nil, fmt.Errorf("couldn't open output file %s: %s",
            outfile, err)
    }

    orig_defer_func := func() {
        out_fh.Close()
    }

    idx := strings.LastIndex(outfile, ".")
    if idx <= -1 || idx >= len(outfile) - 1 {
        return out_fh, orig_defer_func, nil
    }

    suffix := outfile[idx+1:len(outfile)]

    w, compress_defer_func, err := AddCompressionLayer(out_fh, suffix)
    if err != nil {
        if err == UnknownSuffix {
            return out_fh, orig_defer_func, nil
        } else {
            out_fh.Close()
            return nil, nil, fmt.Errorf("couldn't add compression layer: %s", err)
        }
    }

    defer_func := func() {
        compress_defer_func()
        orig_defer_func()
    }

    return w, defer_func, nil
}

// Opens a file in read-only mode. If the file name ends in a supported
// compression suffix, input with be decompressed.
//
// Supported decompression:
//    gzip  (.gz)
//    bzip2 (.bz2)
//    xz    (.xz) -- calls external program
//
// The CloseFunc object is a function that you must call to close the file
// properly.
func OpenFileRO(infile string) (io.Reader, CloseFunc, error) {
    in_fh, err := os.Open(infile)
    if err != nil {
        return nil, nil, err
    }

    orig_defer_func := func() {
        in_fh.Close()
    }

    idx := strings.LastIndex(infile, ".")
    if idx <= -1 || idx >= len(infile) - 1 {
        return in_fh, orig_defer_func, nil
    }

    suffix := infile[idx+1:len(infile)]

    r, compress_defer_func, err := AddDecompressionLayer(in_fh, suffix)
    if err != nil {
        if err == UnknownSuffix {
            return in_fh, orig_defer_func, nil
        } else {
            in_fh.Close()
            return nil, nil, fmt.Errorf("couldn't add decompression layer: %s", err)
        }
    }

    defer_func := func() {
        compress_defer_func()
        orig_defer_func()
    }

    return r, defer_func, nil
}

// Adds compression to output written to writer w, if the suffix is supported.
//
// Supported compression:
//    gzip  (gz)
//    bzip2 (bz2) -- calls external program
//    xz    (xz)  -- calls external program
//
// The CloseFunc object is a function that you must call to shutdown the
// compression layer properly.
func AddCompressionLayer(w io.Writer, suffix string) (io.Writer,
    CloseFunc, error) {

    switch suffix {
    case "gz", "gzip":
        gzip_writer, err := gzip.NewWriterLevel(w, gzip.BestCompression)
        if err != nil {
            return nil, nil, fmt.Errorf("couldn't create gzip writer: %s", err)
        }

        defer_func := func() {
            gzip_writer.Flush()
            gzip_writer.Close()
        }

        return gzip_writer, defer_func, nil

    case "bz2", "bzip2":
        return new_bz2_writer(w)

    case "xz":
        return new_xz_writer(w)
    }

    return nil, nil, UnknownSuffix
}

// Adds decompression to input read from reader r, if the suffix is supported.
//
// Supported decompression:
//    gzip  (gz)
//    bzip2 (bz2)
//    xz    (xz) -- calls external program
//
// The CloseFunc object is a function that you must call to shutdown the
// decompression layer properly.
func AddDecompressionLayer(r io.Reader, suffix string) (io.Reader,
    CloseFunc, error) {

    switch suffix {
    case "gz", "gzip":
        new_reader, err := gzip.NewReader(r)
        if err != nil {
            return nil, nil, fmt.Errorf("couldn't create gzip reader: %s", err)
        }

        defer_func := func() {
            new_reader.Close()
        }

        return new_reader, defer_func, nil

    case "bz2", "bzip2":
        new_reader := bzip2.NewReader(r)
        defer_func := func() { }
        return new_reader, defer_func, nil

    case "xz":
        return new_xz_reader(r)
    }

    return nil, nil, UnknownSuffix
}

// Runs the list of commands, piping the output of each one to the next. The
// output of the last command is sent to the final_writer passed in.
// Each command is represented as a slice of strings. The first element of the
// slice should be the full path to the program to run. The remaining elements
// of the slice should be the arguments to the program.
//
// The writer returned writes to the standard input of the first program
// in the list. The CloseFunc should be called as a function when writing
// has been completed (and before final_writer has been closed).
func OpenPipesToWriter(final_writer io.Writer,
    progs [][]string) (io.Writer, CloseFunc, error) {

    overall_close_func := func() { }
    writer := final_writer

    last := len(progs) - 1
    for i := range progs {
        close_func := overall_close_func
        prog := progs[last - i]
        new_writer, new_close_func, err :=
            get_writer_pipe_from_exec_with_writer(writer, prog...)
        if err != nil {
            overall_close_func()
            return nil, nil, err
        }

        overall_close_func = func() {
            new_close_func()
            close_func()
        }

        writer = new_writer
    }

    return writer, overall_close_func, nil
}

func get_writer_pipe_from_exec_with_writer(prog_stdout io.Writer,
    prog ...string) (io.Writer, CloseFunc, error) {

    name := prog[0]
    args := prog[1:]
    cmd := exec.Command(name, args...)
    cmd.Stdout = prog_stdout

    writer_closer, err := cmd.StdinPipe()
    if err != nil {
        return nil, nil,
        fmt.Errorf("couldn't get stdout pipe in prog writer (%s): %s",
            strings.Join(prog, " "), err)
    }

    err = cmd.Start()
    if err != nil {
        writer_closer.Close()
        return nil, nil, fmt.Errorf("couldn't start process %s: %s",
            strings.Join(prog, " "), err)
    }

    defer_func := func() {
        writer_closer.Close()
        _ = cmd.Wait()
    }

    return writer_closer, defer_func, nil
}

func get_reader_pipe_from_exec_with_reader(prog_stdin io.Reader,
    prog ...string) (io.Reader, CloseFunc, error) {

    name := prog[0]
    args := prog[1:]
    cmd := exec.Command(name, args...)
    cmd.Stdin = prog_stdin
    reader_closer, err := cmd.StdoutPipe()
    if err != nil {
        return nil, nil, fmt.Errorf("couldn't get stdout pipe in prog reader (%s): %s",
            strings.Join(prog, " "), err)
    }

    err = cmd.Start()
    if err != nil {
        reader_closer.Close()
        return nil, nil, fmt.Errorf("couldn't start process %s: %s",
            strings.Join(prog, " "), err)
    }

    defer_func := func() {
        reader_closer.Close()
        _ = cmd.Wait()
    }

    return reader_closer, defer_func, nil

}

func new_bz2_writer(w io.Writer) (io.Writer, CloseFunc, error) {
    path, err := find_exec("bzip2")
    if err !=  nil {
        return nil, nil, err
    }

    return get_writer_pipe_from_exec_with_writer(w, path, "-z", "-c")
}

func new_xz_writer(w io.Writer) (io.Writer, CloseFunc, error) {
    xz_path, err := find_exec("xz")
    if err !=  nil {
        return nil, nil, err
    }

    return get_writer_pipe_from_exec_with_writer(w, xz_path, "-z", "-e", "-c")
}

func new_xz_reader(r io.Reader) (io.Reader, CloseFunc, error) {
    xz_path, err := find_exec("xz")
    if err !=  nil {
        return nil, nil, err
    }

    return get_reader_pipe_from_exec_with_reader(r, xz_path, "-d", "-c")
}

func find_exec(file string) (string, error) {
    dirs := []string{"/bin", "/usr/bin", "/usr/local/bin"}

    for _, dir := range dirs {
        path := fmt.Sprintf("%s/%s", dir, file)
        _, err := os.Stat(path)
        if err == nil {
            return path, nil
        }
    }

    return "", fmt.Errorf("couldn't find executable %s", file)
}

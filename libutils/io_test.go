package libutils_test

import (
    "bufio"
    "bytes"
    "crypto/md5"
    exec "os/exec"
    "fmt"
    // hex "encoding/hex"
    "io"
    ioutil "io/ioutil"
    libutils "github.com/cuberat/go-libutils/libutils"
    "path"
    "os"
    "runtime"
    "strings"
    "testing"
)

func TestPlainFileBuffered(t *testing.T) {
    _, file, _, ok := runtime.Caller(0)
    if !ok {
        t.Error("couldn't get relative path to find data files")
        return
    }

    cur_dir := path.Dir(file)
    data_dir := cur_dir + "/data"

    data_in_file := data_dir + "/rand_plain.txt"

    tmp_fh, err := ioutil.TempFile("", "libutils_io_test_plain_*.txt")
    if err != nil {
        t.Errorf("couldn't create temp file for plain file")
        return
    }
    plain_out_file := tmp_fh.Name()
    tmp_fh.Close()

    md5_data_in_file, err := file_md5(data_in_file)

    in_fh, err := os.Open(data_in_file)
    if err != nil {
        t.Errorf("couldn't open data file %q: %s", data_in_file, err)
        return
    }

    scanner := bufio.NewScanner(in_fh)

    plain_out_fh, err := libutils.CreateFileBuffered(plain_out_file, 0)
    if err != nil {
        t.Errorf("couldn't open plain out file %q for buffered IO: %s",
            plain_out_file, err)
        return
    }
    defer os.Remove(plain_out_file)

    for scanner.Scan() {
        line := scanner.Text()
        fmt.Fprintf(plain_out_fh, "%s\n", line)
    }
    in_fh.Close()

    plain_out_fh.Close()

    md5_plain_out_file, err := file_md5(plain_out_file)
    if !bytes.Equal(md5_plain_out_file, md5_data_in_file) {
        t.Errorf("plain out file incorrect, expected md5 %q, got %q",
            md5_data_in_file, md5_plain_out_file)
    }
}

func TestPlainFile(t *testing.T) {
    _, file, _, ok := runtime.Caller(0)
    if !ok {
        t.Error("couldn't get relative path to find data files")
        return
    }

    cur_dir := path.Dir(file)
    data_dir := cur_dir + "/data"

    data_in_file := data_dir + "/rand_plain.txt"

    tmp_fh, err := ioutil.TempFile("", "libutils_io_test_plain_*.txt")
    if err != nil {
        t.Errorf("couldn't create temp file for plain file")
        return
    }
    plain_out_file := tmp_fh.Name()
    tmp_fh.Close()

    md5_data_in_file, err := file_md5(data_in_file)

    in_fh, err := os.Open(data_in_file)
    if err != nil {
        t.Errorf("couldn't open data file %q: %s", data_in_file, err)
        return
    }

    scanner := bufio.NewScanner(in_fh)

    plain_out_fh, err := libutils.CreateFile(plain_out_file)
    if err != nil {
        t.Errorf("couldn't open plain out file %q for buffered IO: %s",
            plain_out_file, err)
        return
    }
    defer os.Remove(plain_out_file)

    for scanner.Scan() {
        line := scanner.Text()
        fmt.Fprintf(plain_out_fh, "%s\n", line)
    }
    in_fh.Close()

    plain_out_fh.Close()

    md5_plain_out_file, err := file_md5(plain_out_file)
    if !bytes.Equal(md5_plain_out_file, md5_data_in_file) {
        t.Errorf("plain out file incorrect, expected md5 %q, got %q",
            md5_data_in_file, md5_plain_out_file)
    }
}


func TestPlainFileNonBuffered(t *testing.T) {
    _, file, _, ok := runtime.Caller(0)
    if !ok {
        t.Error("couldn't get relative path to find data files")
        return
    }

    cur_dir := path.Dir(file)
    data_dir := cur_dir + "/data"

    data_in_file := data_dir + "/rand_plain.txt"

    tmp_fh, err := ioutil.TempFile("", "libutils_io_test_plain_*.txt")
    if err != nil {
        t.Errorf("couldn't create temp file for plain file")
        return
    }
    plain_out_file := tmp_fh.Name()
    tmp_fh.Close()

    md5_data_in_file, err := file_md5(data_in_file)

    // md5_hex := hex.EncodeToString(md5_data_in_file)
    // t.Errorf("data_in_file: %s", md5_hex)

    in_fh, err := os.Open(data_in_file)
    if err != nil {
        t.Errorf("couldn't open data file %q: %s", data_in_file, err)
        return
    }

    scanner := bufio.NewScanner(in_fh)

    plain_out_fh, err := libutils.CreateFileSync(plain_out_file)
    if err != nil {
        t.Errorf("couldn't open plain out file %q for non-buffered IO: %s",
            plain_out_file, err)
        return
    }
    defer os.Remove(plain_out_file)

    for scanner.Scan() {
        line := scanner.Text()
        fmt.Fprintf(plain_out_fh, "%s\n", line)
    }
    in_fh.Close()

    plain_out_fh.Close()

    md5_plain_out_file, err := file_md5(plain_out_file)
    if !bytes.Equal(md5_plain_out_file, md5_data_in_file) {
        t.Errorf("plain out file incorrect, expected md5 %q, got %q",
            md5_data_in_file, md5_plain_out_file)
    }
}

func TestFileGzipReader(t *testing.T) {
    skip, err := test_compressed_reader("gzip", "gz")
    if skip {
        t.Logf("%s", err)
        t.SkipNow()
        return
    }

    if err != nil {
        t.Errorf("%s", err)
    }
}

func TestFileBzip2Reader(t *testing.T) {
    skip, err := test_compressed_reader("bzip2", "bz2")
    if skip {
        t.Logf("%s", err)
        t.SkipNow()
        return
    }

    if err != nil {
        t.Errorf("%s", err)
    }
}

func TestFileXZReader(t *testing.T) {
    skip, err := test_compressed_reader("xz", "xz")
    if skip {
        t.Logf("%s", err)
        t.SkipNow()
        return
    }

    if err != nil {
        t.Errorf("%s", err)
    }
}

func test_compressed_reader(compress_name, file_ext string) (bool, error) {
    _, err := find_exec(compress_name)
    if err != nil {
        return true, fmt.Errorf("Skipping because `%s` not present",
            compress_name)
    }

    _, file, _, ok := runtime.Caller(0)
    if !ok {
        return false, fmt.Errorf("couldn't get relative path to find data " +
            "files")
    }

    cur_dir := path.Dir(file)
    data_dir := cur_dir + "/data"

    data_in_file := data_dir + "/rand_plain.txt"
    md5_data_in_file, err := file_md5(data_in_file)

    compressed_data_in_file := fmt.Sprintf("%s/rand_plain.txt.%s", data_dir,
        file_ext)

    comp_r, err := libutils.OpenFile(compressed_data_in_file)
    if err != nil {
        return false, fmt.Errorf("couldn't open compressed file %q: %s",
            compressed_data_in_file, err)
    }

    md5_compressed_sum, err := file_md5_reader(comp_r)

    if !bytes.Equal(md5_compressed_sum, md5_data_in_file) {
        return false,
            fmt.Errorf("out file incorrect for %s, expected md5 %q, got %q",
                compressed_data_in_file, md5_data_in_file, md5_compressed_sum)
    }

    return false, nil
}

func test_compressed_writer(compress_name, file_ext string) (bool, error) {
    compress_prog, err := find_exec(compress_name)
    if err != nil {
        return true, fmt.Errorf("Skipping because `%s` not present",
            compress_name)
    }

    _, file, _, ok := runtime.Caller(0)
    if !ok {
        return false, fmt.Errorf("couldn't get relative path to find data " +
            "files")
    }

    cur_dir := path.Dir(file)
    data_dir := cur_dir + "/data"

    data_in_file := data_dir + "/rand_plain.txt"

    tmpl := fmt.Sprintf("libutils_io_test_%s_*.txt.%s", compress_name, file_ext)
    tmp_fh, err := ioutil.TempFile("", tmpl)
    if err != nil {
        return false, fmt.Errorf("couldn't create temp file for `%s` file: %s",
            compress_name, err)
    }
    out_file := tmp_fh.Name()
    tmp_fh.Close()

    if !strings.HasSuffix(out_file, "." + file_ext) {
        out_file += "." + file_ext
    }

    md5_data_in_file, err := file_md5(data_in_file)

    in_fh, err := os.Open(data_in_file)
    if err != nil {
        return false, fmt.Errorf("couldn't open data file %q: %s",
            data_in_file, err)
    }

    out_fh, err := libutils.CreateFile(out_file)
    if err != nil {
        return false,
            fmt.Errorf("couldn't open out file %q for buffered IO: %s",
                out_file, err)
    }
    defer os.Remove(out_file)

    _, err = io.Copy(out_fh, in_fh)

    in_fh.Close()

    out_fh.Close()

    md5_out_file, err := file_md5_uncompress(out_file, compress_prog)

    if !bytes.Equal(md5_out_file, md5_data_in_file) {
        return false,
            fmt.Errorf("out file incorrect for %s, expected md5 %q, got %q",
                out_file, md5_data_in_file, md5_out_file)
    }

    return false, nil
}

func TestFileGzipWriter(t *testing.T) {
    skip, err := test_compressed_writer("gzip", "gz")
    if skip {
        t.Logf("%s", err)
        t.SkipNow()
        return
    }

    if err != nil {
        t.Errorf("%s", err)
    }
}

func TestFileBzip2Writer(t *testing.T) {
    skip, err := test_compressed_writer("bzip2", "bz2")
    if skip {
        t.Logf("%s", err)
        t.SkipNow()
        return
    }

    if err != nil {
        t.Errorf("%s", err)
    }
}

func TestFileXZWrite(t *testing.T) {
    skip, err := test_compressed_writer("xz", "xz")
    if skip {
        t.Logf("%s", err)
        t.SkipNow()
        return
    }

    if err != nil {
        t.Errorf("%s", err)
    }
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

func file_md5_uncompress(file, compress_prog string) ([]byte, error) {
    cmd := exec.Command(compress_prog, "-d", "-c", file)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return nil, fmt.Errorf("couldn't create stdout pipe: %s", err)
    }
    defer stdout.Close()

    if err := cmd.Start(); err != nil {
        return nil, fmt.Errorf("couldn't start cmd: %s", err)
    }

    md5_sum, err := file_md5_reader(stdout)

    if err != nil {
        return nil, fmt.Errorf("couldn't compute md5 sum: %s", err)
    }

    return md5_sum, nil
}

func file_md5_reader(r io.Reader) ([]byte, error) {
    h := md5.New()
    if _, err := io.Copy(h, r); err != nil {
        return nil, fmt.Errorf("couldn't generate md5 hash: %s", err)
    }

    return h.Sum(nil), nil
}

func file_md5(file string) ([]byte, error) {
    in_fh, err := os.Open(file)
    if err != nil {
        return nil, fmt.Errorf("couldn't open file %q for input: %s", file, err)
    }
    defer in_fh.Close()

    return file_md5_reader(in_fh)
}

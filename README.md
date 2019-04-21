

# libutils
`import "github.com/cuberat/go-libutils/libutils"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
The libutils package provides various utilities for working in Go.

Installation


	go get github.com/cuberat/go-libutils/libutils




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func AddCompressionLayer(w io.Writer, suffix string) (io.Writer, CloseFunc, error)](#AddCompressionLayer)
* [func AddDecompressionLayer(r io.Reader, suffix string) (io.Reader, CloseFunc, error)](#AddDecompressionLayer)
* [func DecodeVarint(data_in []byte) (uint64, int, error)](#DecodeVarint)
* [func EncodeVarint(int_in uint64) []byte](#EncodeVarint)
* [func NewPrefixLenScanner(r io.Reader) *bufio.Scanner](#NewPrefixLenScanner)
* [func NewVILPScanner(r io.Reader) *bufio.Scanner](#NewVILPScanner)
* [func NewVILPScannerF(file_path string) (*bufio.Scanner, CloseFunc, error)](#NewVILPScannerF)
* [func OpenFileRO(infile string) (io.Reader, CloseFunc, error)](#OpenFileRO)
* [func OpenFileW(outfile string) (io.Writer, CloseFunc, error)](#OpenFileW)
* [func OpenPipesToWriter(final_writer io.Writer, progs [][]string) (io.Writer, CloseFunc, error)](#OpenPipesToWriter)
* [func ScannerPrefixLenScan(data []byte, at_eof bool) (int, []byte, error)](#ScannerPrefixLenScan)
* [func ScannerVILPScan(data []byte, at_eof bool) (int, []byte, error)](#ScannerVILPScan)
* [type CloseFunc](#CloseFunc)
* [type PrefixLenWriter](#PrefixLenWriter)
  * [func NewPrefixLenWriter(w io.Writer) *PrefixLenWriter](#NewPrefixLenWriter)
  * [func (plw *PrefixLenWriter) Write(p []byte) (int, error)](#PrefixLenWriter.Write)
  * [func (plw *PrefixLenWriter) WriteString(s string) (int, error)](#PrefixLenWriter.WriteString)
* [type VILPWriter](#VILPWriter)
  * [func NewVILPWriter(w io.Writer) *VILPWriter](#NewVILPWriter)
  * [func NewVILPWriterF(file_path string) (*VILPWriter, CloseFunc, error)](#NewVILPWriterF)
  * [func (plw *VILPWriter) Write(p []byte) (int, error)](#VILPWriter.Write)
  * [func (plw *VILPWriter) WriteString(s string) (int, error)](#VILPWriter.WriteString)


#### <a name="pkg-files">Package files</a>
[libutils.go](/src/github.com/cuberat/go-libutils/libutils/libutils.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    UnknownSuffix        error = errors.New("Unknown suffix")
    VarintNotEnoughBytes error = errors.New("Not enough bytes in varint")
)
```


## <a name="AddCompressionLayer">func</a> [AddCompressionLayer](/src/target/libutils.go?s=10904:10990#L368)
``` go
func AddCompressionLayer(w io.Writer, suffix string) (io.Writer,
    CloseFunc, error)
```
Adds compression to output written to writer w, if the suffix is supported.

Supported compression:


	gzip  (gz)
	bzip2 (bz2) -- calls external program
	xz    (xz)  -- calls external program

The CloseFunc object is a function that you must call to shutdown the
compression layer properly.



## <a name="AddDecompressionLayer">func</a> [AddDecompressionLayer](/src/target/libutils.go?s=11824:11912#L404)
``` go
func AddDecompressionLayer(r io.Reader, suffix string) (io.Reader,
    CloseFunc, error)
```
Adds decompression to input read from reader r, if the suffix is supported.

Supported decompression:


	gzip  (gz)
	bzip2 (bz2)
	xz    (xz) -- calls external program

The CloseFunc object is a function that you must call to shutdown the
decompression layer properly.



## <a name="DecodeVarint">func</a> [DecodeVarint](/src/target/libutils.go?s=2004:2058#L43)
``` go
func DecodeVarint(data_in []byte) (uint64, int, error)
```
Decodes a varint (as used in protobuffers) into a uint64.
See <a href="https://developers.google.com/protocol-buffers/docs/encoding#varints">https://developers.google.com/protocol-buffers/docs/encoding#varints</a>
for the specification.



## <a name="EncodeVarint">func</a> [EncodeVarint](/src/target/libutils.go?s=2599:2640#L67)
``` go
func EncodeVarint(int_in uint64) []byte
```
Encodes a uint64 as a varint (as used in protobuffers).
See <a href="https://developers.google.com/protocol-buffers/docs/encoding#varints">https://developers.google.com/protocol-buffers/docs/encoding#varints</a>
for the specification.



## <a name="NewPrefixLenScanner">func</a> [NewPrefixLenScanner](/src/target/libutils.go?s=7157:7211#L232)
``` go
func NewPrefixLenScanner(r io.Reader) *bufio.Scanner
```
Returns a bufio.Scanner that scans length-prefixed strings from the
provided io.Reader.

Deprecated: use NewVILPScanner and varint length-prefixed files.



## <a name="NewVILPScanner">func</a> [NewVILPScanner](/src/target/libutils.go?s=4333:4382#L133)
``` go
func NewVILPScanner(r io.Reader) *bufio.Scanner
```
Returns a bufio.Scanner that scans varint length-prefixed strings from the
provided io.Reader.



## <a name="NewVILPScannerF">func</a> [NewVILPScannerF](/src/target/libutils.go?s=4629:4706#L142)
``` go
func NewVILPScannerF(file_path string) (*bufio.Scanner,
    CloseFunc, error)
```
Returns a bufio.Scanner that scans varint length-prefixed strings from the
provided file. Call close_func() to close the underlying file handle.



## <a name="OpenFileRO">func</a> [OpenFileRO](/src/target/libutils.go?s=9751:9811#L324)
``` go
func OpenFileRO(infile string) (io.Reader, CloseFunc, error)
```
Opens a file in read-only mode. If the file name ends in a supported
compression suffix, input with be decompressed.

Supported decompression:


	gzip  (.gz)
	bzip2 (.bz2)
	xz    (.xz) -- calls external program

The CloseFunc object is a function that you must call to close the file
properly.



## <a name="OpenFileW">func</a> [OpenFileW](/src/target/libutils.go?s=8514:8574#L278)
``` go
func OpenFileW(outfile string) (io.Writer, CloseFunc, error)
```
Open a file for writing. If the file name dends in a supported
compression suffix, output will be compressed in that format.

Supported compression:


	gzip  (.gz)
	bzip2 (.bz2) -- calls external program
	xz    (.xz)  -- calls external program

The CloseFunc object is a function that you must call to close the file
properly.



## <a name="OpenPipesToWriter">func</a> [OpenPipesToWriter](/src/target/libutils.go?s=13046:13144#L441)
``` go
func OpenPipesToWriter(final_writer io.Writer,
    progs [][]string) (io.Writer, CloseFunc, error)
```
Runs the list of commands, piping the output of each one to the next. The
output of the last command is sent to the final_writer passed in.
Each command is represented as a slice of strings. The first element of the
slice should be the full path to the program to run. The remaining elements
of the slice should be the arguments to the program.

The writer returned writes to the standard input of the first program
in the list. The CloseFunc should be called as a function when writing
has been completed (and before final_writer has been closed).



## <a name="ScannerPrefixLenScan">func</a> [ScannerPrefixLenScan](/src/target/libutils.go?s=7453:7525#L242)
``` go
func ScannerPrefixLenScan(data []byte, at_eof bool) (int, []byte, error)
```
A bufio.SplitFunc that reads length-prefixed strings from a reader

Deprecated: use NewVILPScanner and varint length-prefixed files.



## <a name="ScannerVILPScan">func</a> [ScannerVILPScan](/src/target/libutils.go?s=4958:5025#L156)
``` go
func ScannerVILPScan(data []byte, at_eof bool) (int, []byte, error)
```
A bufio.SplitFunc that reads length-prefixed strings from a reader.




## <a name="CloseFunc">type</a> [CloseFunc](/src/target/libutils.go?s=8129:8153#L266)
``` go
type CloseFunc func()
```
Signature for Close() function return from OpenFileW and
OpenFileRO. When ready to close the file, call the function to
close and clean up.










## <a name="PrefixLenWriter">type</a> [PrefixLenWriter](/src/target/libutils.go?s=5658:5705#L182)
``` go
type PrefixLenWriter struct {
    // contains filtered or unexported fields
}
```
PrefixLenWriter is used to write length-prefixed strings to an io.Writer

Deprecated: use VILPWriter and its corresponding methods.







### <a name="NewPrefixLenWriter">func</a> [NewPrefixLenWriter](/src/target/libutils.go?s=6869:6924#L221)
``` go
func NewPrefixLenWriter(w io.Writer) *PrefixLenWriter
```
Returns a new PrefixLenWriter. PrefixLenWriter implements the
io.Writer interface, in addition to the WriteString method.

Deprecated: use VILPWriter and its corresponding methods.





### <a name="PrefixLenWriter.Write">func</a> (\*PrefixLenWriter) [Write](/src/target/libutils.go?s=6215:6270#L198)
``` go
func (plw *PrefixLenWriter) Write(p []byte) (int, error)
```
Writes the provided bytes as a length-prefixed string to the
underlying io.Writer. This uses 32-bit integers for the length prefix.

Deprecated: use VILPWriter and its corresponding methods.




### <a name="PrefixLenWriter.WriteString">func</a> (\*PrefixLenWriter) [WriteString](/src/target/libutils.go?s=5910:5971#L190)
``` go
func (plw *PrefixLenWriter) WriteString(s string) (int, error)
```
Writes the provided string as a length-prefixed string to the
underlying io.Writer. This uses 32-bit integers for the length prefix.

Deprecated: use VILPWriter and its corresponding methods.




## <a name="VILPWriter">type</a> [VILPWriter](/src/target/libutils.go?s=3036:3078#L84)
``` go
type VILPWriter struct {
    // contains filtered or unexported fields
}
```
VILPWriter is used to write length-prefixed strings to an io.Writer







### <a name="NewVILPWriter">func</a> [NewVILPWriter](/src/target/libutils.go?s=3745:3790#L112)
``` go
func NewVILPWriter(w io.Writer) *VILPWriter
```
Returns a new VILPWriter. VILPWriter implements the
io.Writer interface, in addition to the WriteString method.


### <a name="NewVILPWriterF">func</a> [NewVILPWriterF](/src/target/libutils.go?s=4005:4078#L121)
``` go
func NewVILPWriterF(file_path string) (*VILPWriter, CloseFunc,
    error)
```
Opens the provided file and returns a *VILPWriter created using the
resulting file handle. Call close_func() to close the underlying file handle.





### <a name="VILPWriter.Write">func</a> (\*VILPWriter) [Write](/src/target/libutils.go?s=3355:3405#L96)
``` go
func (plw *VILPWriter) Write(p []byte) (int, error)
```
Writes the provided bytes as a length-prefixed string to the
underlying io.Writer




### <a name="VILPWriter.WriteString">func</a> (\*VILPWriter) [WriteString](/src/target/libutils.go?s=3169:3225#L90)
``` go
func (plw *VILPWriter) WriteString(s string) (int, error)
```
Writes the provided string as a length-prefixed string to the
underlying io.Writer








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)

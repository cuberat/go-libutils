

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
* [func NewVILenPrefixScanner(r io.Reader) *bufio.Scanner](#NewVILenPrefixScanner)
* [func NewVILenPrefixScannerFromFile(file_path string) (*bufio.Scanner, CloseFunc, error)](#NewVILenPrefixScannerFromFile)
* [func OpenFileRO(infile string) (io.Reader, CloseFunc, error)](#OpenFileRO)
* [func OpenFileW(outfile string) (io.Writer, CloseFunc, error)](#OpenFileW)
* [func OpenPipesToWriter(final_writer io.Writer, progs [][]string) (io.Writer, CloseFunc, error)](#OpenPipesToWriter)
* [func ScannerPrefixLenScan(data []byte, at_eof bool) (int, []byte, error)](#ScannerPrefixLenScan)
* [func ScannerVILenPrefixScan(data []byte, at_eof bool) (int, []byte, error)](#ScannerVILenPrefixScan)
* [type CloseFunc](#CloseFunc)
* [type PrefixLenWriter](#PrefixLenWriter)
  * [func NewPrefixLenWriter(w io.Writer) *PrefixLenWriter](#NewPrefixLenWriter)
  * [func (plw *PrefixLenWriter) Write(p []byte) (int, error)](#PrefixLenWriter.Write)
  * [func (plw *PrefixLenWriter) WriteString(s string) (int, error)](#PrefixLenWriter.WriteString)
* [type VILenPrefixWriter](#VILenPrefixWriter)
  * [func NewVILenPrefixWriter(w io.Writer) *VILenPrefixWriter](#NewVILenPrefixWriter)
  * [func (plw *VILenPrefixWriter) Write(p []byte) (int, error)](#VILenPrefixWriter.Write)
  * [func (plw *VILenPrefixWriter) WriteString(s string) (int, error)](#VILenPrefixWriter.WriteString)


#### <a name="pkg-files">Package files</a>
[libutils.go](/src/github.com/cuberat/go-libutils/libutils/libutils.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    UnknownSuffix        error = errors.New("Unknown suffix")
    VarintNotEnoughBytes error = errors.New("Not enough bytes in varint")
)
```


## <a name="AddCompressionLayer">func</a> [AddCompressionLayer](/src/target/libutils.go?s=10076:10162#L344)
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



## <a name="AddDecompressionLayer">func</a> [AddDecompressionLayer](/src/target/libutils.go?s=10996:11084#L380)
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



## <a name="NewPrefixLenScanner">func</a> [NewPrefixLenScanner](/src/target/libutils.go?s=6400:6454#L210)
``` go
func NewPrefixLenScanner(r io.Reader) *bufio.Scanner
```
Returns a bufio.Scanner that scans length-prefixed strings from the
provided io.Reader



## <a name="NewVILenPrefixScanner">func</a> [NewVILenPrefixScanner](/src/target/libutils.go?s=4017:4073#L121)
``` go
func NewVILenPrefixScanner(r io.Reader) *bufio.Scanner
```
Returns a bufio.Scanner that scans varint length-prefixed strings from the
provided io.Reader.



## <a name="NewVILenPrefixScannerFromFile">func</a> [NewVILenPrefixScannerFromFile](/src/target/libutils.go?s=4272:4363#L130)
``` go
func NewVILenPrefixScannerFromFile(file_path string) (*bufio.Scanner,
    CloseFunc, error)
```
Returns a bufio.Scanner that scans varint length-prefixed strings from the
provided file.



## <a name="OpenFileRO">func</a> [OpenFileRO](/src/target/libutils.go?s=8923:8983#L300)
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



## <a name="OpenFileW">func</a> [OpenFileW](/src/target/libutils.go?s=7686:7746#L254)
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



## <a name="OpenPipesToWriter">func</a> [OpenPipesToWriter](/src/target/libutils.go?s=12218:12316#L417)
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



## <a name="ScannerPrefixLenScan">func</a> [ScannerPrefixLenScan](/src/target/libutils.go?s=6625:6697#L218)
``` go
func ScannerPrefixLenScan(data []byte, at_eof bool) (int, []byte, error)
```
A bufio.SplitFunc that reads length-prefixed strings from a reader



## <a name="ScannerVILenPrefixScan">func</a> [ScannerVILenPrefixScan](/src/target/libutils.go?s=4622:4696#L144)
``` go
func ScannerVILenPrefixScan(data []byte, at_eof bool) (int, []byte, error)
```
A bufio.SplitFunc that reads length-prefixed strings from a reader.




## <a name="CloseFunc">type</a> [CloseFunc](/src/target/libutils.go?s=7301:7325#L242)
``` go
type CloseFunc func()
```
Signature for Close() function return from OpenFileW and
OpenFileRO. When ready to close the file, call the function to
close and clean up.










## <a name="PrefixLenWriter">type</a> [PrefixLenWriter](/src/target/libutils.go?s=5265:5312#L168)
``` go
type PrefixLenWriter struct {
    // contains filtered or unexported fields
}
```
PrefixLenWriter is used to write length-prefixed strings to an io.Writer







### <a name="NewPrefixLenWriter">func</a> [NewPrefixLenWriter](/src/target/libutils.go?s=6184:6239#L201)
``` go
func NewPrefixLenWriter(w io.Writer) *PrefixLenWriter
```
Returns a new PrefixLenWriter. PrefixLenWriter implements the
io.Writer interface, in addition to the WriteString method.





### <a name="PrefixLenWriter.Write">func</a> (\*PrefixLenWriter) [Write](/src/target/libutils.go?s=5594:5649#L180)
``` go
func (plw *PrefixLenWriter) Write(p []byte) (int, error)
```
Writes the provided bytes as a length-prefixed string to the
underlying io.Writer




### <a name="PrefixLenWriter.WriteString">func</a> (\*PrefixLenWriter) [WriteString](/src/target/libutils.go?s=5403:5464#L174)
``` go
func (plw *PrefixLenWriter) WriteString(s string) (int, error)
```
Writes the provided string as a length-prefixed string to the
underlying io.Writer




## <a name="VILenPrefixWriter">type</a> [VILenPrefixWriter](/src/target/libutils.go?s=3043:3092#L84)
``` go
type VILenPrefixWriter struct {
    // contains filtered or unexported fields
}
```
VILenPrefixWriter is used to write length-prefixed strings to an io.Writer







### <a name="NewVILenPrefixWriter">func</a> [NewVILenPrefixWriter](/src/target/libutils.go?s=3787:3846#L112)
``` go
func NewVILenPrefixWriter(w io.Writer) *VILenPrefixWriter
```
Returns a new VILenPrefixWriter. VILenPrefixWriter implements the
io.Writer interface, in addition to the WriteString method.





### <a name="VILenPrefixWriter.Write">func</a> (\*VILenPrefixWriter) [Write](/src/target/libutils.go?s=3376:3433#L96)
``` go
func (plw *VILenPrefixWriter) Write(p []byte) (int, error)
```
Writes the provided bytes as a length-prefixed string to the
underlying io.Writer




### <a name="VILenPrefixWriter.WriteString">func</a> (\*VILenPrefixWriter) [WriteString](/src/target/libutils.go?s=3183:3246#L90)
``` go
func (plw *VILenPrefixWriter) WriteString(s string) (int, error)
```
Writes the provided string as a length-prefixed string to the
underlying io.Writer








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)

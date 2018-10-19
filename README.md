

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
* [func NewPrefixLenScanner(r io.Reader) *bufio.Scanner](#NewPrefixLenScanner)
* [func OpenFileRO(infile string) (io.Reader, CloseFunc, error)](#OpenFileRO)
* [func OpenFileW(outfile string) (io.Writer, CloseFunc, error)](#OpenFileW)
* [func ScannerPrefixLenScan(data []byte, at_eof bool) (int, []byte, error)](#ScannerPrefixLenScan)
* [type CloseFunc](#CloseFunc)
* [type PrefixLenWriter](#PrefixLenWriter)
  * [func NewPrefixLenWriter(w io.Writer) *PrefixLenWriter](#NewPrefixLenWriter)
  * [func (plw *PrefixLenWriter) Write(p []byte) (int, error)](#PrefixLenWriter.Write)
  * [func (plw *PrefixLenWriter) WriteString(s string) (int, error)](#PrefixLenWriter.WriteString)


#### <a name="pkg-files">Package files</a>
[libutils.go](/src/github.com/cuberat/go-libutils/libutils/libutils.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    UnknownSuffix error = errors.New("Unknown suffix")
)
```


## <a name="AddCompressionLayer">func</a> [AddCompressionLayer](/src/target/libutils.go?s=6650:6736#L216)
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



## <a name="AddDecompressionLayer">func</a> [AddDecompressionLayer](/src/target/libutils.go?s=7570:7658#L252)
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



## <a name="NewPrefixLenScanner">func</a> [NewPrefixLenScanner](/src/target/libutils.go?s=2973:3027#L82)
``` go
func NewPrefixLenScanner(r io.Reader) *bufio.Scanner
```
Returns a bufio.Scanner that scans length-prefixed strings from the
provided io.Reader



## <a name="OpenFileRO">func</a> [OpenFileRO](/src/target/libutils.go?s=5496:5556#L172)
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



## <a name="OpenFileW">func</a> [OpenFileW](/src/target/libutils.go?s=4259:4319#L126)
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



## <a name="ScannerPrefixLenScan">func</a> [ScannerPrefixLenScan](/src/target/libutils.go?s=3198:3270#L90)
``` go
func ScannerPrefixLenScan(data []byte, at_eof bool) (int, []byte, error)
```
A bufio.SplitFunc that reads length-prefixed strings from a reader




## <a name="CloseFunc">type</a> [CloseFunc](/src/target/libutils.go?s=3874:3898#L114)
``` go
type CloseFunc func()
```
Signature for Close() function return from OpenFileW and
OpenFileRO. When ready to close the file, call the function to
close and clean up.










## <a name="PrefixLenWriter">type</a> [PrefixLenWriter](/src/target/libutils.go?s=1838:1885#L40)
``` go
type PrefixLenWriter struct {
    // contains filtered or unexported fields
}
```
PrefixLenWriter is used to write length-prefixed strings to an io.Writer







### <a name="NewPrefixLenWriter">func</a> [NewPrefixLenWriter](/src/target/libutils.go?s=2757:2812#L73)
``` go
func NewPrefixLenWriter(w io.Writer) *PrefixLenWriter
```
Returns a new PrefixLenWriter. PrefixLenWriter implements the
io.Writer interface, in addition to the WriteString method.





### <a name="PrefixLenWriter.Write">func</a> (\*PrefixLenWriter) [Write](/src/target/libutils.go?s=2167:2222#L52)
``` go
func (plw *PrefixLenWriter) Write(p []byte) (int, error)
```
Writes the provided bytes as a length-prefixed string to the
underlying io.Writer




### <a name="PrefixLenWriter.WriteString">func</a> (\*PrefixLenWriter) [WriteString](/src/target/libutils.go?s=1976:2037#L46)
``` go
func (plw *PrefixLenWriter) WriteString(s string) (int, error)
```
Writes the provided string as a length-prefixed string to the
underlying io.Writer


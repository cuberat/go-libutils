

# libutils
`import "github.com/cuberat/go-libutils/libutils"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
The libutils package provides various utilities for working in Go.

Installation


	go get github.com/cuberat/go-libutils/libutils




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
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
* [type KeyedRecord](#KeyedRecord)
  * [func NewKeyedRecordFromBytes(raw_rec_bytes []byte, decoder KeyedRecordDecoder) *KeyedRecord](#NewKeyedRecordFromBytes)
  * [func NewKeyedRecordFromKeyVal(key []byte, val interface{}) *KeyedRecord](#NewKeyedRecordFromKeyVal)
  * [func (kr *KeyedRecord) Key() ([]byte, error)](#KeyedRecord.Key)
  * [func (kr *KeyedRecord) KeyString() (string, error)](#KeyedRecord.KeyString)
  * [func (kr *KeyedRecord) KeyVal() ([]byte, interface{}, error)](#KeyedRecord.KeyVal)
  * [func (kr *KeyedRecord) RecordBytesOut(encoder KeyedRecordEncoder) ([]byte, error)](#KeyedRecord.RecordBytesOut)
  * [func (kr *KeyedRecord) SetDecoder(dec KeyedRecordDecoder)](#KeyedRecord.SetDecoder)
  * [func (kr *KeyedRecord) SetEncoder(enc KeyedRecordEncoder)](#KeyedRecord.SetEncoder)
  * [func (kr *KeyedRecord) Val() (interface{}, error)](#KeyedRecord.Val)
* [type KeyedRecordDecoder](#KeyedRecordDecoder)
* [type KeyedRecordEncoder](#KeyedRecordEncoder)
* [type KeyedRecordScanner](#KeyedRecordScanner)
* [type KeyedRecordWriter](#KeyedRecordWriter)
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


## <a name="pkg-constants">Constants</a>
``` go
const (
    Version = "0.9"
)
```

## <a name="pkg-variables">Variables</a>
``` go
var (
    UnknownSuffix        error = errors.New("Unknown suffix")
    VarintNotEnoughBytes error = errors.New("Not enough bytes in varint")
)
```


## <a name="AddCompressionLayer">func</a> [AddCompressionLayer](/src/target/libutils.go?s=17867:17953#L601)
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



## <a name="AddDecompressionLayer">func</a> [AddDecompressionLayer](/src/target/libutils.go?s=18787:18875#L637)
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



## <a name="DecodeVarint">func</a> [DecodeVarint](/src/target/libutils.go?s=8967:9021#L276)
``` go
func DecodeVarint(data_in []byte) (uint64, int, error)
```
Decodes a varint (as used in protobuffers) into a uint64.
See <a href="https://developers.google.com/protocol-buffers/docs/encoding#varints">https://developers.google.com/protocol-buffers/docs/encoding#varints</a>
for the specification.



## <a name="EncodeVarint">func</a> [EncodeVarint](/src/target/libutils.go?s=9562:9603#L300)
``` go
func EncodeVarint(int_in uint64) []byte
```
Encodes a uint64 as a varint (as used in protobuffers).
See <a href="https://developers.google.com/protocol-buffers/docs/encoding#varints">https://developers.google.com/protocol-buffers/docs/encoding#varints</a>
for the specification.



## <a name="NewPrefixLenScanner">func</a> [NewPrefixLenScanner](/src/target/libutils.go?s=14120:14174#L465)
``` go
func NewPrefixLenScanner(r io.Reader) *bufio.Scanner
```
Returns a bufio.Scanner that scans length-prefixed strings from the
provided io.Reader.

Deprecated: use NewVILPScanner and varint length-prefixed files.



## <a name="NewVILPScanner">func</a> [NewVILPScanner](/src/target/libutils.go?s=11296:11345#L366)
``` go
func NewVILPScanner(r io.Reader) *bufio.Scanner
```
Returns a bufio.Scanner that scans varint length-prefixed strings from the
provided io.Reader.



## <a name="NewVILPScannerF">func</a> [NewVILPScannerF](/src/target/libutils.go?s=11592:11669#L375)
``` go
func NewVILPScannerF(file_path string) (*bufio.Scanner,
    CloseFunc, error)
```
Returns a bufio.Scanner that scans varint length-prefixed strings from the
provided file. Call close_func() to close the underlying file handle.



## <a name="OpenFileRO">func</a> [OpenFileRO](/src/target/libutils.go?s=16714:16774#L557)
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



## <a name="OpenFileW">func</a> [OpenFileW](/src/target/libutils.go?s=15477:15537#L511)
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



## <a name="OpenPipesToWriter">func</a> [OpenPipesToWriter](/src/target/libutils.go?s=20009:20107#L674)
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



## <a name="ScannerPrefixLenScan">func</a> [ScannerPrefixLenScan](/src/target/libutils.go?s=14416:14488#L475)
``` go
func ScannerPrefixLenScan(data []byte, at_eof bool) (int, []byte, error)
```
A bufio.SplitFunc that reads length-prefixed strings from a reader

Deprecated: use NewVILPScanner and varint length-prefixed files.



## <a name="ScannerVILPScan">func</a> [ScannerVILPScan](/src/target/libutils.go?s=11921:11988#L389)
``` go
func ScannerVILPScan(data []byte, at_eof bool) (int, []byte, error)
```
A bufio.SplitFunc that reads length-prefixed strings from a reader.




## <a name="CloseFunc">type</a> [CloseFunc](/src/target/libutils.go?s=15092:15116#L499)
``` go
type CloseFunc func()
```
Signature for Close() function return from OpenFileW and
OpenFileRO. When ready to close the file, call the function to
close and clean up.










## <a name="KeyedRecord">type</a> [KeyedRecord](/src/target/libutils.go?s=4330:4603#L106)
``` go
type KeyedRecord struct {
    // contains filtered or unexported fields
}
```
KeyedRecord in this package supports the concept of records that consist of a
(string) key and a (possible complex) value, along with lazy marshaling and
unmarshaling of data.

Extracting records from files is considered separate from the records
themselves. E.g., records might live in a tab-delimited file where each
record is stored as a key and a JSON object separated by a tab character,
with a newline character delimiting records, like so:


	foo[tab]{"on": true}[newline]
	bar[tab]{"on": false}[newline]

In this case, the newline is not part of the record. The KeyedRecordScanner
and KeyedRecordWriter interfaces deal with reading and writing records. A
KeyedRecord needs an decoder (KeyedRecordDecoder) in order to parse a record
provided to it as a slice of bytes, and an encoder (KeyedRecordEncoder) to
serialize the record to be written out. This allows for readily changing the
output format, e.g., to variable integer length-prefixed records.







### <a name="NewKeyedRecordFromBytes">func</a> [NewKeyedRecordFromBytes](/src/target/libutils.go?s=4724:4821#L122)
``` go
func NewKeyedRecordFromBytes(raw_rec_bytes []byte,
    decoder KeyedRecordDecoder) *KeyedRecord
```
Parse the raw record from wire data, using the provided decoder. The decoder
is stored internally for later use.


### <a name="NewKeyedRecordFromKeyVal">func</a> [NewKeyedRecordFromKeyVal](/src/target/libutils.go?s=5018:5091#L133)
``` go
func NewKeyedRecordFromKeyVal(key []byte, val interface{}) *KeyedRecord
```
Create a new KeyedRecord object from a key and value.





### <a name="KeyedRecord.Key">func</a> (\*KeyedRecord) [Key](/src/target/libutils.go?s=5594:5638#L154)
``` go
func (kr *KeyedRecord) Key() ([]byte, error)
```
Parse out the key from the record (if necessary) and return it.




### <a name="KeyedRecord.KeyString">func</a> (\*KeyedRecord) [KeyString](/src/target/libutils.go?s=6461:6511#L187)
``` go
func (kr *KeyedRecord) KeyString() (string, error)
```
Parse out the key from the record (if necessary) and return it as a string.




### <a name="KeyedRecord.KeyVal">func</a> (\*KeyedRecord) [KeyVal](/src/target/libutils.go?s=7357:7417#L219)
``` go
func (kr *KeyedRecord) KeyVal() ([]byte, interface{}, error)
```
Parse out (if necessary) the key and value, returning both.




### <a name="KeyedRecord.RecordBytesOut">func</a> (\*KeyedRecord) [RecordBytesOut](/src/target/libutils.go?s=7848:7929#L237)
``` go
func (kr *KeyedRecord) RecordBytesOut(encoder KeyedRecordEncoder) ([]byte, error)
```
Serialize the record into a slice of bytes using the provided encoder.




### <a name="KeyedRecord.SetDecoder">func</a> (\*KeyedRecord) [SetDecoder](/src/target/libutils.go?s=5295:5352#L144)
``` go
func (kr *KeyedRecord) SetDecoder(dec KeyedRecordDecoder)
```
Set the decoder object within the KeyedRecord for later use.




### <a name="KeyedRecord.SetEncoder">func</a> (\*KeyedRecord) [SetEncoder](/src/target/libutils.go?s=5443:5500#L149)
``` go
func (kr *KeyedRecord) SetEncoder(enc KeyedRecordEncoder)
```
Set the encoder object within the KeyedRecord for later use.




### <a name="KeyedRecord.Val">func</a> (\*KeyedRecord) [Val](/src/target/libutils.go?s=6781:6830#L197)
``` go
func (kr *KeyedRecord) Val() (interface{}, error)
```
Return the value of from the record as an interface{}. If you know what type
the value should have, you can use an assertion to get to the underlying
type, e.g.,


	val, ok := kr.Val().(*MyStruct)




## <a name="KeyedRecordDecoder">type</a> [KeyedRecordDecoder](/src/target/libutils.go?s=2545:2796#L61)
``` go
type KeyedRecordDecoder interface {
    // Splits the record, returning the key and the serialized value data
    // structure.
    SplitKV([]byte) ([]byte, []byte, error)

    // Deserializes the value.
    UnmarshalVal([]byte) (interface{}, error)
}
```









## <a name="KeyedRecordEncoder">type</a> [KeyedRecordEncoder](/src/target/libutils.go?s=1872:2543#L44)
``` go
type KeyedRecordEncoder interface {
    // Joins the key and value bytes, returning the serialized record.
    JoinKV(key []byte, val []byte) ([]byte, error)

    // Serializes the value data structure.
    MarshalVal(interface{}) ([]byte, error)

    // If this object also implements the `KeyedRecordDecoder` interface, and
    // the encoding is the same for both input and output, CodecSame() returns
    // true. Otherwise, it returns false.
    //
    // This allows for lazy encoding. That is, if the raw record bytes that were
    // read in do not need to change, they can be written back out as-is,
    // instead of actually re-encoding.
    CodecSame() bool
}
```









## <a name="KeyedRecordScanner">type</a> [KeyedRecordScanner](/src/target/libutils.go?s=2799:3200#L71)
``` go
type KeyedRecordScanner interface {
    // Advances the scanner to the next record. It returns false when the scan
    // stops, either by reaching the end of the input or an error.
    Scan() bool

    // Returns the most recent serialized record generated by a call to Scan().
    Record() *KeyedRecord

    // Returns the first non-EOF error that was encountered by the Scanner.
    Err() error
}
```









## <a name="KeyedRecordWriter">type</a> [KeyedRecordWriter](/src/target/libutils.go?s=3202:3318#L83)
``` go
type KeyedRecordWriter interface {
    // Writes the entire seralized record.
    Write(*KeyedRecord) (int, error)
}
```









## <a name="PrefixLenWriter">type</a> [PrefixLenWriter](/src/target/libutils.go?s=12621:12668#L415)
``` go
type PrefixLenWriter struct {
    // contains filtered or unexported fields
}
```
PrefixLenWriter is used to write length-prefixed strings to an io.Writer

Deprecated: use VILPWriter and its corresponding methods.







### <a name="NewPrefixLenWriter">func</a> [NewPrefixLenWriter](/src/target/libutils.go?s=13832:13887#L454)
``` go
func NewPrefixLenWriter(w io.Writer) *PrefixLenWriter
```
Returns a new PrefixLenWriter. PrefixLenWriter implements the
io.Writer interface, in addition to the WriteString method.

Deprecated: use VILPWriter and its corresponding methods.





### <a name="PrefixLenWriter.Write">func</a> (\*PrefixLenWriter) [Write](/src/target/libutils.go?s=13178:13233#L431)
``` go
func (plw *PrefixLenWriter) Write(p []byte) (int, error)
```
Writes the provided bytes as a length-prefixed string to the
underlying io.Writer. This uses 32-bit integers for the length prefix.

Deprecated: use VILPWriter and its corresponding methods.




### <a name="PrefixLenWriter.WriteString">func</a> (\*PrefixLenWriter) [WriteString](/src/target/libutils.go?s=12873:12934#L423)
``` go
func (plw *PrefixLenWriter) WriteString(s string) (int, error)
```
Writes the provided string as a length-prefixed string to the
underlying io.Writer. This uses 32-bit integers for the length prefix.

Deprecated: use VILPWriter and its corresponding methods.




## <a name="VILPWriter">type</a> [VILPWriter](/src/target/libutils.go?s=9999:10041#L317)
``` go
type VILPWriter struct {
    // contains filtered or unexported fields
}
```
VILPWriter is used to write length-prefixed strings to an io.Writer







### <a name="NewVILPWriter">func</a> [NewVILPWriter](/src/target/libutils.go?s=10708:10753#L345)
``` go
func NewVILPWriter(w io.Writer) *VILPWriter
```
Returns a new VILPWriter. VILPWriter implements the
io.Writer interface, in addition to the WriteString method.


### <a name="NewVILPWriterF">func</a> [NewVILPWriterF](/src/target/libutils.go?s=10968:11041#L354)
``` go
func NewVILPWriterF(file_path string) (*VILPWriter, CloseFunc,
    error)
```
Opens the provided file and returns a *VILPWriter created using the
resulting file handle. Call close_func() to close the underlying file handle.





### <a name="VILPWriter.Write">func</a> (\*VILPWriter) [Write](/src/target/libutils.go?s=10318:10368#L329)
``` go
func (plw *VILPWriter) Write(p []byte) (int, error)
```
Writes the provided bytes as a length-prefixed string to the
underlying io.Writer




### <a name="VILPWriter.WriteString">func</a> (\*VILPWriter) [WriteString](/src/target/libutils.go?s=10132:10188#L323)
``` go
func (plw *VILPWriter) WriteString(s string) (int, error)
```
Writes the provided string as a length-prefixed string to the
underlying io.Writer








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)



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
* [func BytesToVILP(data []byte) []byte](#BytesToVILP)
* [func DecodeVarint(data_in []byte) (uint64, int, error)](#DecodeVarint)
* [func EncodeVarint(int_in uint64) []byte](#EncodeVarint)
* [func Errorf(fmt_str string, args ...interface{}) error](#Errorf)
* [func ErrorfLong(fmt_str string, args ...interface{}) error](#ErrorfLong)
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
  * [func (kr *KeyedRecord) MarkDirty()](#KeyedRecord.MarkDirty)
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
* [type TabColsKRCodec](#TabColsKRCodec)
  * [func NewTabColsKRCodec() *TabColsKRCodec](#NewTabColsKRCodec)
  * [func (krc *TabColsKRCodec) CodecSame() bool](#TabColsKRCodec.CodecSame)
  * [func (krc *TabColsKRCodec) JoinKV(key, val []byte) ([]byte, error)](#TabColsKRCodec.JoinKV)
  * [func (krc *TabColsKRCodec) MarshalVal(data interface{}) ([]byte, error)](#TabColsKRCodec.MarshalVal)
  * [func (krc *TabColsKRCodec) SplitKV(wire_data []byte) ([]byte, []byte, error)](#TabColsKRCodec.SplitKV)
  * [func (krc *TabColsKRCodec) UnmarshalVal(val_bytes []byte) (interface{}, error)](#TabColsKRCodec.UnmarshalVal)
* [type TabColsKRScanner](#TabColsKRScanner)
  * [func NewTabColsKRScanner(r io.Reader) *TabColsKRScanner](#NewTabColsKRScanner)
  * [func (krs *TabColsKRScanner) Err() error](#TabColsKRScanner.Err)
  * [func (krs *TabColsKRScanner) Record() *KeyedRecord](#TabColsKRScanner.Record)
  * [func (krs *TabColsKRScanner) Scan() bool](#TabColsKRScanner.Scan)
* [type TabColsKRWriter](#TabColsKRWriter)
  * [func NewTabColsKRWriter(w io.Writer) *TabColsKRWriter](#NewTabColsKRWriter)
  * [func (krw *TabColsKRWriter) Write(rec *KeyedRecord) (int, error)](#TabColsKRWriter.Write)
* [type VILPColsKRCodec](#VILPColsKRCodec)
  * [func NewVILPColsKRCodec() *VILPColsKRCodec](#NewVILPColsKRCodec)
  * [func (krc *VILPColsKRCodec) CodecSame() bool](#VILPColsKRCodec.CodecSame)
  * [func (krc *VILPColsKRCodec) JoinKV(key, val []byte) ([]byte, error)](#VILPColsKRCodec.JoinKV)
  * [func (krc *VILPColsKRCodec) MarshalVal(data interface{}) ([]byte, error)](#VILPColsKRCodec.MarshalVal)
  * [func (krc *VILPColsKRCodec) SplitKV(wire_data []byte) ([]byte, []byte, error)](#VILPColsKRCodec.SplitKV)
  * [func (krc *VILPColsKRCodec) UnmarshalVal(val_bytes []byte) (interface{}, error)](#VILPColsKRCodec.UnmarshalVal)
* [type VILPColsKRScanner](#VILPColsKRScanner)
  * [func NewVILPColsKRScanner(r io.Reader) *VILPColsKRScanner](#NewVILPColsKRScanner)
  * [func (krs *VILPColsKRScanner) Err() error](#VILPColsKRScanner.Err)
  * [func (krs *VILPColsKRScanner) Record() *KeyedRecord](#VILPColsKRScanner.Record)
  * [func (krs *VILPColsKRScanner) Scan() bool](#VILPColsKRScanner.Scan)
* [type VILPColsKRWriter](#VILPColsKRWriter)
  * [func NewVILPColsKRWriter(w io.Writer) *VILPColsKRWriter](#NewVILPColsKRWriter)
  * [func (krw *VILPColsKRWriter) Write(rec *KeyedRecord) (int, error)](#VILPColsKRWriter.Write)
* [type VILPWriter](#VILPWriter)
  * [func NewVILPWriter(w io.Writer) *VILPWriter](#NewVILPWriter)
  * [func NewVILPWriterF(file_path string) (*VILPWriter, CloseFunc, error)](#NewVILPWriterF)
  * [func (plw *VILPWriter) Write(p []byte) (int, error)](#VILPWriter.Write)
  * [func (plw *VILPWriter) WriteString(s string) (int, error)](#VILPWriter.WriteString)


#### <a name="pkg-files">Package files</a>
[keyed_record.go](/src/github.com/cuberat/go-libutils/libutils/keyed_record.go) [libutils.go](/src/github.com/cuberat/go-libutils/libutils/libutils.go) [prefix_len.go](/src/github.com/cuberat/go-libutils/libutils/prefix_len.go) [tab_cols_kr.go](/src/github.com/cuberat/go-libutils/libutils/tab_cols_kr.go) [varint.go](/src/github.com/cuberat/go-libutils/libutils/varint.go) [vilp.go](/src/github.com/cuberat/go-libutils/libutils/vilp.go) [vilp_cols_kr.go](/src/github.com/cuberat/go-libutils/libutils/vilp_cols_kr.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    Version = "1.02"
)
```

## <a name="pkg-variables">Variables</a>
``` go
var (
    UnknownSuffix        error = errors.New("Unknown suffix")
    VarintNotEnoughBytes error = errors.New("Not enough bytes in varint")
)
```


## <a name="AddCompressionLayer">func</a> [AddCompressionLayer](/src/target/libutils.go?s=5679:5765#L181)
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



## <a name="AddDecompressionLayer">func</a> [AddDecompressionLayer](/src/target/libutils.go?s=6599:6687#L217)
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



## <a name="BytesToVILP">func</a> [BytesToVILP](/src/target/vilp.go?s=1560:1598#L26)
``` go
func BytesToVILP(data []byte) []byte
```
Returns a byte slice with a varint length prefix followed by the provided
byte slice.



## <a name="DecodeVarint">func</a> [DecodeVarint](/src/target/varint.go?s=1610:1664#L25)
``` go
func DecodeVarint(data_in []byte) (uint64, int, error)
```
Decodes a varint (as used in protobuffers) into a uint64.
See <a href="https://developers.google.com/protocol-buffers/docs/encoding#varints">https://developers.google.com/protocol-buffers/docs/encoding#varints</a>
for the specification.



## <a name="EncodeVarint">func</a> [EncodeVarint](/src/target/varint.go?s=2205:2246#L49)
``` go
func EncodeVarint(int_in uint64) []byte
```
Encodes a uint64 as a varint (as used in protobuffers).
See <a href="https://developers.google.com/protocol-buffers/docs/encoding#varints">https://developers.google.com/protocol-buffers/docs/encoding#varints</a>
for the specification.



## <a name="Errorf">func</a> [Errorf](/src/target/libutils.go?s=2024:2078#L47)
``` go
func Errorf(fmt_str string, args ...interface{}) error
```
Like fmt.Errorf(), except adds the (base) file name and line number to the
beginning of the error message in the format `[%s:%d] `.



## <a name="ErrorfLong">func</a> [ErrorfLong](/src/target/libutils.go?s=2476:2534#L63)
``` go
func ErrorfLong(fmt_str string, args ...interface{}) error
```
Like fmt.Errorf(), except adds the full file name and line number to the
beginning of the error message in the format `[%s:%d] `.



## <a name="NewPrefixLenScanner">func</a> [NewPrefixLenScanner](/src/target/prefix_len.go?s=3107:3161#L77)
``` go
func NewPrefixLenScanner(r io.Reader) *bufio.Scanner
```
Returns a bufio.Scanner that scans length-prefixed strings from the
provided io.Reader.

Deprecated: use NewVILPScanner and varint length-prefixed files.



## <a name="NewVILPScanner">func</a> [NewVILPScanner](/src/target/vilp.go?s=3033:3082#L76)
``` go
func NewVILPScanner(r io.Reader) *bufio.Scanner
```
Returns a bufio.Scanner that scans varint length-prefixed strings from the
provided io.Reader.



## <a name="NewVILPScannerF">func</a> [NewVILPScannerF](/src/target/vilp.go?s=3329:3406#L85)
``` go
func NewVILPScannerF(file_path string) (*bufio.Scanner,
    CloseFunc, error)
```
Returns a bufio.Scanner that scans varint length-prefixed strings from the
provided file. Call close_func() to close the underlying file handle.



## <a name="OpenFileRO">func</a> [OpenFileRO](/src/target/libutils.go?s=4526:4586#L137)
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



## <a name="OpenFileW">func</a> [OpenFileW](/src/target/libutils.go?s=3289:3349#L91)
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



## <a name="OpenPipesToWriter">func</a> [OpenPipesToWriter](/src/target/libutils.go?s=7821:7919#L254)
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



## <a name="ScannerPrefixLenScan">func</a> [ScannerPrefixLenScan](/src/target/prefix_len.go?s=3403:3475#L87)
``` go
func ScannerPrefixLenScan(data []byte, at_eof bool) (int, []byte, error)
```
A bufio.SplitFunc that reads length-prefixed strings from a reader

Deprecated: use NewVILPScanner and varint length-prefixed files.



## <a name="ScannerVILPScan">func</a> [ScannerVILPScan](/src/target/vilp.go?s=3658:3725#L99)
``` go
func ScannerVILPScan(data []byte, at_eof bool) (int, []byte, error)
```
A bufio.SplitFunc that reads length-prefixed strings from a reader.




## <a name="CloseFunc">type</a> [CloseFunc](/src/target/libutils.go?s=2904:2928#L79)
``` go
type CloseFunc func()
```
Signature for Close() function return from OpenFileW and
OpenFileRO. When ready to close the file, call the function to
close and clean up.










## <a name="KeyedRecord">type</a> [KeyedRecord](/src/target/keyed_record.go?s=3904:4177#L83)
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







### <a name="NewKeyedRecordFromBytes">func</a> [NewKeyedRecordFromBytes](/src/target/keyed_record.go?s=4574:4671#L108)
``` go
func NewKeyedRecordFromBytes(raw_rec_bytes []byte,
    decoder KeyedRecordDecoder) *KeyedRecord
```
Parse the raw record from wire data, using the provided decoder. The decoder
is stored internally for later use.


### <a name="NewKeyedRecordFromKeyVal">func</a> [NewKeyedRecordFromKeyVal](/src/target/keyed_record.go?s=4868:4941#L119)
``` go
func NewKeyedRecordFromKeyVal(key []byte, val interface{}) *KeyedRecord
```
Create a new KeyedRecord object from a key and value.





### <a name="KeyedRecord.Key">func</a> (\*KeyedRecord) [Key](/src/target/keyed_record.go?s=5444:5488#L140)
``` go
func (kr *KeyedRecord) Key() ([]byte, error)
```
Parse out the key from the record (if necessary) and return it.




### <a name="KeyedRecord.KeyString">func</a> (\*KeyedRecord) [KeyString](/src/target/keyed_record.go?s=6311:6361#L173)
``` go
func (kr *KeyedRecord) KeyString() (string, error)
```
Parse out the key from the record (if necessary) and return it as a string.




### <a name="KeyedRecord.KeyVal">func</a> (\*KeyedRecord) [KeyVal](/src/target/keyed_record.go?s=7207:7267#L205)
``` go
func (kr *KeyedRecord) KeyVal() ([]byte, interface{}, error)
```
Parse out (if necessary) the key and value, returning both.




### <a name="KeyedRecord.MarkDirty">func</a> (\*KeyedRecord) [MarkDirty](/src/target/keyed_record.go?s=4351:4385#L101)
``` go
func (kr *KeyedRecord) MarkDirty()
```
Mark the KeyedRecord dirty such that will reserialize the value data
structure even if the encoder and encoder are the same and
encoder.CodecSame() returns true.




### <a name="KeyedRecord.RecordBytesOut">func</a> (\*KeyedRecord) [RecordBytesOut](/src/target/keyed_record.go?s=7698:7779#L223)
``` go
func (kr *KeyedRecord) RecordBytesOut(encoder KeyedRecordEncoder) ([]byte, error)
```
Serialize the record into a slice of bytes using the provided encoder.




### <a name="KeyedRecord.SetDecoder">func</a> (\*KeyedRecord) [SetDecoder](/src/target/keyed_record.go?s=5145:5202#L130)
``` go
func (kr *KeyedRecord) SetDecoder(dec KeyedRecordDecoder)
```
Set the decoder object within the KeyedRecord for later use.




### <a name="KeyedRecord.SetEncoder">func</a> (\*KeyedRecord) [SetEncoder](/src/target/keyed_record.go?s=5293:5350#L135)
``` go
func (kr *KeyedRecord) SetEncoder(enc KeyedRecordEncoder)
```
Set the encoder object within the KeyedRecord for later use.




### <a name="KeyedRecord.Val">func</a> (\*KeyedRecord) [Val](/src/target/keyed_record.go?s=6631:6680#L183)
``` go
func (kr *KeyedRecord) Val() (interface{}, error)
```
Return the value of from the record as an interface{}. If you know what type
the value should have, you can use an assertion to get to the underlying
type, e.g.,


	val, ok := kr.Val().(*MyStruct)




## <a name="KeyedRecordDecoder">type</a> [KeyedRecordDecoder](/src/target/keyed_record.go?s=2120:2371#L39)
``` go
type KeyedRecordDecoder interface {
    // Splits the record, returning the key and the serialized value data
    // structure.
    SplitKV([]byte) ([]byte, []byte, error)

    // Deserializes the value.
    UnmarshalVal([]byte) (interface{}, error)
}
```









## <a name="KeyedRecordEncoder">type</a> [KeyedRecordEncoder](/src/target/keyed_record.go?s=1447:2118#L22)
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









## <a name="KeyedRecordScanner">type</a> [KeyedRecordScanner](/src/target/keyed_record.go?s=2373:2774#L48)
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









## <a name="KeyedRecordWriter">type</a> [KeyedRecordWriter](/src/target/keyed_record.go?s=2776:2892#L60)
``` go
type KeyedRecordWriter interface {
    // Writes the entire seralized record.
    Write(*KeyedRecord) (int, error)
}
```









## <a name="PrefixLenWriter">type</a> [PrefixLenWriter](/src/target/prefix_len.go?s=1608:1655#L27)
``` go
type PrefixLenWriter struct {
    // contains filtered or unexported fields
}
```
PrefixLenWriter is used to write length-prefixed strings to an io.Writer

Deprecated: use VILPWriter and its corresponding methods.







### <a name="NewPrefixLenWriter">func</a> [NewPrefixLenWriter](/src/target/prefix_len.go?s=2819:2874#L66)
``` go
func NewPrefixLenWriter(w io.Writer) *PrefixLenWriter
```
Returns a new PrefixLenWriter. PrefixLenWriter implements the
io.Writer interface, in addition to the WriteString method.

Deprecated: use VILPWriter and its corresponding methods.





### <a name="PrefixLenWriter.Write">func</a> (\*PrefixLenWriter) [Write](/src/target/prefix_len.go?s=2165:2220#L43)
``` go
func (plw *PrefixLenWriter) Write(p []byte) (int, error)
```
Writes the provided bytes as a length-prefixed string to the
underlying io.Writer. This uses 32-bit integers for the length prefix.

Deprecated: use VILPWriter and its corresponding methods.




### <a name="PrefixLenWriter.WriteString">func</a> (\*PrefixLenWriter) [WriteString](/src/target/prefix_len.go?s=1860:1921#L35)
``` go
func (plw *PrefixLenWriter) WriteString(s string) (int, error)
```
Writes the provided string as a length-prefixed string to the
underlying io.Writer. This uses 32-bit integers for the length prefix.

Deprecated: use VILPWriter and its corresponding methods.




## <a name="TabColsKRCodec">type</a> [TabColsKRCodec](/src/target/tab_cols_kr.go?s=1942:1973#L35)
``` go
type TabColsKRCodec struct {
}
```
Implements the KeyedRecordEncoder and KeyedRecordDecoder interfaces specified
by `github.com/cuberat/go-libutils/libutils`.

This codec serializes and deserializes keyed records where the key is the
first tab-separated column, and the value is a slice of byte slices
containing the remaining columns. This is not TSV, where tab bytes can be
escaped, etc. It is a simple split on tab bytes. That is,


	<key>[tab]<val1>[tab]<val2>...







### <a name="NewTabColsKRCodec">func</a> [NewTabColsKRCodec](/src/target/tab_cols_kr.go?s=2007:2050#L40)
``` go
func NewTabColsKRCodec() *TabColsKRCodec
```
Returns a new TabColsKRCodec





### <a name="TabColsKRCodec.CodecSame">func</a> (\*TabColsKRCodec) [CodecSame](/src/target/tab_cols_kr.go?s=3411:3454#L87)
``` go
func (krc *TabColsKRCodec) CodecSame() bool
```
Returns true so that if this codec is used for both encoder and decoder,
unnecessary re-serialization can be avoided.

This allows for lazy encoding. That is, if the raw record bytes that were
read in do not need to change, they can be written back out as-is, instead of
actually re-encoding.




### <a name="TabColsKRCodec.JoinKV">func</a> (\*TabColsKRCodec) [JoinKV](/src/target/tab_cols_kr.go?s=2662:2728#L67)
``` go
func (krc *TabColsKRCodec) JoinKV(key, val []byte) ([]byte, error)
```
Joins the key and value bytes, returning the serialized record.




### <a name="TabColsKRCodec.MarshalVal">func</a> (\*TabColsKRCodec) [MarshalVal](/src/target/tab_cols_kr.go?s=2835:2906#L72)
``` go
func (krc *TabColsKRCodec) MarshalVal(data interface{}) ([]byte, error)
```
Serializes the value data structure.




### <a name="TabColsKRCodec.SplitKV">func</a> (\*TabColsKRCodec) [SplitKV](/src/target/tab_cols_kr.go?s=2171:2251#L46)
``` go
func (krc *TabColsKRCodec) SplitKV(wire_data []byte) ([]byte, []byte,
    error)
```
Splits the record, returning the key and the serialized value data
structure.




### <a name="TabColsKRCodec.UnmarshalVal">func</a> (\*TabColsKRCodec) [UnmarshalVal](/src/target/tab_cols_kr.go?s=2453:2535#L60)
``` go
func (krc *TabColsKRCodec) UnmarshalVal(val_bytes []byte) (interface{},
    error)
```
Deserializes the value.




## <a name="TabColsKRScanner">type</a> [TabColsKRScanner](/src/target/tab_cols_kr.go?s=4502:4592#L125)
``` go
type TabColsKRScanner struct {
    // contains filtered or unexported fields
}
```
Implements the `libutils.KeyedRecordScanner` interface from
`github.com/cuberat/go-libutils/libutils`.
This is a scanner for tab-delimited keyed records with column-based values.
Records are expected to look like


	<key>[tab]<val1>[tab]<val2>...[newline]







### <a name="NewTabColsKRScanner">func</a> [NewTabColsKRScanner](/src/target/tab_cols_kr.go?s=4627:4684#L131)
``` go
func NewTabColsKRScanner(r io.Reader) *TabColsKRScanner
```
Return a new TabColsKRScanner





### <a name="TabColsKRScanner.Err">func</a> (\*TabColsKRScanner) [Err](/src/target/tab_cols_kr.go?s=5424:5464#L155)
``` go
func (krs *TabColsKRScanner) Err() error
```
Returns the first non-EOF error that was encountered by the Scanner.




### <a name="TabColsKRScanner.Record">func</a> (\*TabColsKRScanner) [Record](/src/target/tab_cols_kr.go?s=5105:5157#L146)
``` go
func (krs *TabColsKRScanner) Record() *KeyedRecord
```
Returns the most recent serialized record generated by a call to Scan().




### <a name="TabColsKRScanner.Scan">func</a> (\*TabColsKRScanner) [Scan](/src/target/tab_cols_kr.go?s=4953:4993#L141)
``` go
func (krs *TabColsKRScanner) Scan() bool
```
Advances the scanner to the next record. It returns false when the scan
stops, either by reaching the end of the input or an error.




## <a name="TabColsKRWriter">type</a> [TabColsKRWriter](/src/target/tab_cols_kr.go?s=3651:3734#L95)
``` go
type TabColsKRWriter struct {
    // contains filtered or unexported fields
}
```
Implements the `libutils.KeyedRecordWriter` interface from
`github.com/cuberat/go-libutils/libutils`. Records look like


	<key>[tab]<val1>[tab]<val2>...[newline]







### <a name="NewTabColsKRWriter">func</a> [NewTabColsKRWriter](/src/target/tab_cols_kr.go?s=3770:3826#L101)
``` go
func NewTabColsKRWriter(w io.Writer) *TabColsKRWriter
```
Returns a new TabColsKRWriter.





### <a name="TabColsKRWriter.Write">func</a> (\*TabColsKRWriter) [Write](/src/target/tab_cols_kr.go?s=3993:4057#L110)
``` go
func (krw *TabColsKRWriter) Write(rec *KeyedRecord) (int, error)
```
The Write method for the KeyedRecordWriter interface




## <a name="VILPColsKRCodec">type</a> [VILPColsKRCodec](/src/target/vilp_cols_kr.go?s=1888:1920#L32)
``` go
type VILPColsKRCodec struct {
}
```
Implements the KeyedRecordEncoder and KeyedRecordDecoder interfaces specified
by `github.com/cuberat/go-libutils/libutils`.

This codec serializes and deserializes keyed records where the key is
length-prefixed using a varint. The value immediately follows and consists of
length-prefixed strings represented an array/list of strings. That is,


	<key_len><key><val1_len><val1><val2_len><val2>...







### <a name="NewVILPColsKRCodec">func</a> [NewVILPColsKRCodec](/src/target/vilp_cols_kr.go?s=1955:2000#L37)
``` go
func NewVILPColsKRCodec() *VILPColsKRCodec
```
Returns a new VILPColsKRCodec





### <a name="VILPColsKRCodec.CodecSame">func</a> (\*VILPColsKRCodec) [CodecSame](/src/target/vilp_cols_kr.go?s=4524:4568#L121)
``` go
func (krc *VILPColsKRCodec) CodecSame() bool
```
Returns true so that if this codec is used for both encoder and decoder,
unnecessary re-serialization can be avoided.

This allows for lazy encoding. That is, if the raw record bytes that were
read in do not need to change, they can be written back out as-is, instead of
actually re-encoding.




### <a name="VILPColsKRCodec.JoinKV">func</a> (\*VILPColsKRCodec) [JoinKV](/src/target/vilp_cols_kr.go?s=3398:3465#L88)
``` go
func (krc *VILPColsKRCodec) JoinKV(key, val []byte) ([]byte, error)
```
Joins the key and value bytes, returning the serialized record.




### <a name="VILPColsKRCodec.MarshalVal">func</a> (\*VILPColsKRCodec) [MarshalVal](/src/target/vilp_cols_kr.go?s=3753:3825#L99)
``` go
func (krc *VILPColsKRCodec) MarshalVal(data interface{}) ([]byte, error)
```
Serializes the value data structure.




### <a name="VILPColsKRCodec.SplitKV">func</a> (\*VILPColsKRCodec) [SplitKV](/src/target/vilp_cols_kr.go?s=2122:2203#L43)
``` go
func (krc *VILPColsKRCodec) SplitKV(wire_data []byte) ([]byte, []byte,
    error)
```
Splits the record, returning the key and the serialized value data
structure.




### <a name="VILPColsKRCodec.UnmarshalVal">func</a> (\*VILPColsKRCodec) [UnmarshalVal](/src/target/vilp_cols_kr.go?s=2677:2760#L65)
``` go
func (krc *VILPColsKRCodec) UnmarshalVal(val_bytes []byte) (interface{},
    error)
```
Deserializes the value.




## <a name="VILPColsKRScanner">type</a> [VILPColsKRScanner](/src/target/vilp_cols_kr.go?s=5464:5555#L154)
``` go
type VILPColsKRScanner struct {
    // contains filtered or unexported fields
}
```
Scanner for varint length-prefixed keyed records with column-based values.
That is, the value is an array/list of strings, each one prefixed with a
length.







### <a name="NewVILPColsKRScanner">func</a> [NewVILPColsKRScanner](/src/target/vilp_cols_kr.go?s=5591:5650#L160)
``` go
func NewVILPColsKRScanner(r io.Reader) *VILPColsKRScanner
```
Return a new VILPColsKRScanner





### <a name="VILPColsKRScanner.Err">func</a> (\*VILPColsKRScanner) [Err](/src/target/vilp_cols_kr.go?s=6392:6433#L184)
``` go
func (krs *VILPColsKRScanner) Err() error
```
Returns the first non-EOF error that was encountered by the Scanner.




### <a name="VILPColsKRScanner.Record">func</a> (\*VILPColsKRScanner) [Record](/src/target/vilp_cols_kr.go?s=6072:6125#L175)
``` go
func (krs *VILPColsKRScanner) Record() *KeyedRecord
```
Returns the most recent serialized record generated by a call to Scan().




### <a name="VILPColsKRScanner.Scan">func</a> (\*VILPColsKRScanner) [Scan](/src/target/vilp_cols_kr.go?s=5919:5960#L170)
``` go
func (krs *VILPColsKRScanner) Scan() bool
```
Advances the scanner to the next record. It returns false when the scan
stops, either by reaching the end of the input or an error.




## <a name="VILPColsKRWriter">type</a> [VILPColsKRWriter](/src/target/vilp_cols_kr.go?s=4698:4789#L127)
``` go
type VILPColsKRWriter struct {
    // contains filtered or unexported fields
}
```
Implements the `libutils.KeyedRecordWriter` interface from
`github.com/cuberat/go-libutils/libutils`.







### <a name="NewVILPColsKRWriter">func</a> [NewVILPColsKRWriter](/src/target/vilp_cols_kr.go?s=4826:4884#L133)
``` go
func NewVILPColsKRWriter(w io.Writer) *VILPColsKRWriter
```
Returns a new VILPColsKRWriter.





### <a name="VILPColsKRWriter.Write">func</a> (\*VILPColsKRWriter) [Write](/src/target/vilp_cols_kr.go?s=5073:5138#L142)
``` go
func (krw *VILPColsKRWriter) Write(rec *KeyedRecord) (int, error)
```
The Write method for the KeyedRecordWriter interface




## <a name="VILPWriter">type</a> [VILPWriter](/src/target/vilp.go?s=1853:1895#L36)
``` go
type VILPWriter struct {
    // contains filtered or unexported fields
}
```
VILPWriter is used to write length-prefixed strings to an io.Writer







### <a name="NewVILPWriter">func</a> [NewVILPWriter](/src/target/vilp.go?s=2445:2490#L55)
``` go
func NewVILPWriter(w io.Writer) *VILPWriter
```
Returns a new VILPWriter. VILPWriter implements the
io.Writer interface, in addition to the WriteString method.


### <a name="NewVILPWriterF">func</a> [NewVILPWriterF](/src/target/vilp.go?s=2705:2778#L64)
``` go
func NewVILPWriterF(file_path string) (*VILPWriter, CloseFunc,
    error)
```
Opens the provided file and returns a *VILPWriter created using the
resulting file handle. Call close_func() to close the underlying file handle.





### <a name="VILPWriter.Write">func</a> (\*VILPWriter) [Write](/src/target/vilp.go?s=2232:2282#L49)
``` go
func (plw *VILPWriter) Write(p []byte) (int, error)
```
Writes the provided bytes as a length-prefixed string to the
underlying io.Writer




### <a name="VILPWriter.WriteString">func</a> (\*VILPWriter) [WriteString](/src/target/vilp.go?s=1986:2042#L42)
``` go
func (plw *VILPWriter) WriteString(s string) (int, error)
```
Writes the provided string as a length-prefixed string to the
underlying io.Writer








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)

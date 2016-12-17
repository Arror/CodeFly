package global

import "github.com/samuel/go-thrift/parser"

var (
	// Lang language
	Lang string

	// Input input thrift file path
	Input string

	// Output output files path
	Output string

	// ThriftMapping thrift mapping
	ThriftMapping map[string]*parser.Thrift
)

const (
	// Swift Swift language
	Swift = "swift"
)

const (
	// ThriftI16 i16
	ThriftI16 = "i16"
	// ThriftI32 i32
	ThriftI32 = "i32"
	// ThriftI64 i64
	ThriftI64 = "i64"
	// ThriftString string
	ThriftString = "string"
	// ThriftBool bool
	ThriftBool = "bool"
	// ThriftDouble double
	ThriftDouble = "double"
	// ThriftByte byte
	ThriftByte = "byte"
	// ThriftBinary binary
	ThriftBinary = "binary"
)

const (
	// Set set
	Set = "set"
	// Map map
	Map = "map"
	// List list
	List = "list"
)

const (
	// Unsupported unsupported
	Unsupported = "unsupported"
)

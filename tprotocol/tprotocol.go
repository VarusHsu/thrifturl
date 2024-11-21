package tprotocol

const (
	STOP   = 0
	VOID   = 1
	BOOL   = 2
	BYTE   = 3
	I08    = 3
	DOUBLE = 4
	I16    = 6
	I32    = 8
	I64    = 10
	STRING = 11
	UTF7   = 11
	STRUCT = 12
	MAP    = 13
	SET    = 14
	LIST   = 15
	UTF8   = 16
	UTF16  = 17
	BINARY = 18
	FLOAT  = 19
)

type serializedProtocol struct {
	type_ byte   // constants in the Thrift protocol
	id    int16  // field num in Thrift protocol defined
	name  string // field name

	value any // value storage
	/*
		type_ = VOID, value is nil
		type_ = BOOL, value type of golang bool
		type_ = BYTE, I08, value type of byte
		type_ = DOUBLE, value type of float64
		type_ = I16, value type of int16
		type_ = I32, value type of int32
		type_ = I64, value type of int64
		type_ = STRING, UTF7, UTF8, UTF16 value type of golang string
		type_ = STRUCT，MAP, value type of []serializedProtocol
		type_ = SET, LIST type of []any
		type_ = FLOAT type of float32
	*/
}

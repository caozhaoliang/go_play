package geeRPC

import "io"

type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         string
}
type NewCodecFunc func(closer io.ReadWriteCloser) Codec
type Type string

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCode
}

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

type Codec interface {
	io.Closer
	ReadHeader(header *Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

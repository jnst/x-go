package grpcapp

import (
	"bytes"
	"encoding/gob"
	bytefmt "github.com/labstack/gommon/bytes"
	"google.golang.org/protobuf/proto"
)

func binarySize(val interface{}) string {
	if message, ok := val.(proto.Message); ok {
		if b, err := proto.Marshal(message); err == nil {
			return bytefmt.Format(int64(len(b)))
		}
	}

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(val)
	if err != nil {
		return "0B"
	}
	return bytefmt.Format(int64(len(buff.Bytes())))
}

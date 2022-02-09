package grpcapp

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"

	bytefmt "github.com/labstack/gommon/bytes"
)

func binarySize(val interface{}) string {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(val)
	if err != nil {
		return "0B"
	}
	size := binary.Size(buff.Bytes())
	return bytefmt.Format(int64(size))
}

package p2p

import (
	"encoding/gob"
	"io"
)

// interface definition
type Decoder interface {
	Decode(io.Reader, *Message) error
}

// concrete type
type GOBDecoder struct{}

func (d GOBDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct{}

func (d *DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 2000)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}
	msg.Payload = buf[:n]
	return nil
}

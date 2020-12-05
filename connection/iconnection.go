package connection

import "github.com/gmsorg/gms/protocol"

type IConnection interface {
	Send(reqData []byte) error
	// Read(response interface{}) error
	Read() (protocol.Imessage, error)
}

package packet

import (
	"Learning2/cocoabeans/protocol"
	"bytes"
)

type HandshakingHandshakeInPacket struct {
	protocolVersion uint8
	address         string
	port            uint16
	nextState       protocol.ConnectionState
}

func (packet HandshakingHandshakeInPacket) WriteTo(buffer bytes.Buffer) {
	buffer.WriteByte(0)
}

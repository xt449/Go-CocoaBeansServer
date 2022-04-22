package packet

type Listener interface {
	handleHandshakingHandshake(packet HandshakingHandshakeInPacket)
}

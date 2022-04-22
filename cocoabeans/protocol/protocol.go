package protocol

type ConnectionState uint8

const (
	HANDSHAKING ConnectionState = iota
	STATUS      ConnectionState = iota
	LOGIN       ConnectionState = iota
	PLAY        ConnectionState = iota
)

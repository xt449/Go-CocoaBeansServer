package packet_data

type Vibration struct {
	Origin        Position
	PositionType  String
	BlockPosition Position
	EntityId      VarInt
	Ticks         VarInt
}

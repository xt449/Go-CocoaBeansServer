package particle_data

import "Learning2/cocoabeans/protocol/packet/packet_data"

type Vibration struct {
	Origin        packet_data.Position
	PositionType  packet_data.String
	BlockPosition packet_data.Position
	EntityId      packet_data.VarInt
	Ticks         packet_data.VarInt
}

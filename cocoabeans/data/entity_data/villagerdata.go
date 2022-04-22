package entity_data

import "Learning2/cocoabeans/protocol/packet/packet_data"

type VillagerData struct {
	Type       packet_data.VarInt
	Profession packet_data.VarInt
	Level      packet_data.VarInt
}

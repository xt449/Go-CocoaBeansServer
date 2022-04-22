package particle_data

import (
	"Learning2/cocoabeans/data/entity_data"
	"Learning2/cocoabeans/protocol/packet/packet_data"
)

type Particle[D ParticleData] struct {
	Id   packet_data.VarInt
	Data D
}

type ParticleData interface {
	interface{} | entity_data.OptBlockID | Dust | DustTransition | packet_data.Slot | Vibration
}

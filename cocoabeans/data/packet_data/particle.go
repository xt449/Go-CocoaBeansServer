package packet_data

import (
	"Learning2/cocoabeans/data/entity_data"
)

type Particle[D ParticleData] struct {
	Id   VarInt
	Data D
}

type ParticleData interface {
	interface{} | entity_data.OptBlockID | Dust | DustTransition | Slot | Vibration
}

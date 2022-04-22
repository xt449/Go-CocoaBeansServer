package packet_data

type Particle[D ParticleData] struct {
	Id   VarInt
	Data D
}

type ParticleData interface {
	interface{} | OptBlockID | Dust | DustTransition | Slot | Vibration
}

package particle_data

import "Learning2/cocoabeans/protocol/packet/packet_data"

type Dust struct {
	Red   packet_data.Float
	Green packet_data.Float
	Blue  packet_data.Float
	Scale packet_data.Float // Clamped [0.01..4]
}

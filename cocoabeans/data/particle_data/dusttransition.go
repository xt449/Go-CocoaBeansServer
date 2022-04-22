package particle_data

import "Learning2/cocoabeans/protocol/packet/packet_data"

type DustTransition struct {
	FromRed   packet_data.Float
	FromGreen packet_data.Float
	FromBlue  packet_data.Float
	Scale     packet_data.Float // Clamped [0.01..4]
	ToRed     packet_data.Float
	ToGreen   packet_data.Float
	ToBlue    packet_data.Float
}

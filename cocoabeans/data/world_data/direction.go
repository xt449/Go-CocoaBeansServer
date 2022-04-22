package world_data

import "Learning2/cocoabeans/protocol/packet/packet_data"

type Direction packet_data.Enum[Direction]

const (
	Down  Direction = iota
	Up    Direction = iota
	North Direction = iota
	South Direction = iota
	West  Direction = iota
	East  Direction = iota
)

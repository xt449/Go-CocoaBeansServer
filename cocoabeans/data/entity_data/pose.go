package entity_data

import "Learning2/cocoabeans/protocol/packet/packet_data"

type Pose packet_data.Enum[Pose]

const (
	STANDING     Pose = iota
	FALL_FLYING  Pose = iota
	SLEEPING     Pose = iota
	SWIMMING     Pose = iota
	SPIN_ATTACK  Pose = iota
	SNEAKING     Pose = iota
	LONG_JUMPING Pose = iota
	DYING        Pose = iota
)

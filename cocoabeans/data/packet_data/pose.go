package packet_data

type Pose Enum[Pose]

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

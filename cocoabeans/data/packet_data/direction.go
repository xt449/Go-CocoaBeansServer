package packet_data

type Direction Enum[Direction]

const (
	Down  Direction = iota
	Up    Direction = iota
	North Direction = iota
	South Direction = iota
	West  Direction = iota
	East  Direction = iota
)

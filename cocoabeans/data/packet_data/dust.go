package packet_data

type Dust struct {
	Red   Float
	Green Float
	Blue  Float
	Scale Float // Clamped [0.01..4]
}

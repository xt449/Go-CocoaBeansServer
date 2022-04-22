package packet_data

type DustTransition struct {
	FromRed   Float
	FromGreen Float
	FromBlue  Float
	Scale     Float // Clamped [0.01..4]
	ToRed     Float
	ToGreen   Float
	ToBlue    Float
}

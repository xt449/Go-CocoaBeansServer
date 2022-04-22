package math_extra

func MinOf[N int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64](vars ...N) N {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOf[N int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64](vars ...N) N {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

// Clamping

func Clamp[N int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64](value N, min N, max N) N {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func ClampToInt8[N int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64](value N, min int8, max int8) int8 {
	if int8(value) < min {
		return min
	}
	if int8(value) > max {
		return max
	}
	return int8(value)
}

func ClampToUInt8[N int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64](value N, min uint8, max uint8) uint8 {
	if uint8(value) < min {
		return min
	}
	if uint8(value) > max {
		return max
	}
	return uint8(value)
}

func ClampToInt16[N int32 | uint32 | int64 | uint64 | float32 | float64](value N, min int16, max int16) int16 {
	if int16(value) < min {
		return min
	}
	if int16(value) > max {
		return max
	}
	return int16(value)
}

func ClampToUInt16[N int32 | uint32 | int64 | uint64 | float32 | float64](value N, min uint16, max uint16) uint16 {
	if uint16(value) < min {
		return min
	}
	if uint16(value) > max {
		return max
	}
	return uint16(value)
}

func LoopClamp[N int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](value N, min N, max N) N {
	return (value % (max + 1)) + min
}

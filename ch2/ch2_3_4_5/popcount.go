package ch2_3_4_5

var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 2.3
func PopCountCycle(x uint64) int {
	var result int
	for i := uint(0); i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

// 2.4
func PopCountRight(x uint64) int {
	var result int
	for i := uint(0); i < 8; i++ {
		if x&1 != 0 {
			result++
		}
		x >>= 1
	}
	return result
}

// 2.5
func PopCountClearing(x uint64) int {
	var result int
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return result
}

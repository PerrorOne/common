// 获取int/int8/int32/int64/int32/float32/float64数组的最大值
package utils

func MaxInt64(d []int64) (max int64, index int) {
	for index1, v := range d {
		if v > max {
			max = v
			index = index1
		}
	}
	return
}

func MaxInt(d []int) (max int, index int) {
	for index1, v := range d {
		if v > max {
			max = v
			index = index1
		}
	}
	return
}

func MaxInt32(d []int32) (max int32, index int) {
	for index1, v := range d {
		if v > max {
			max = v
			index = index1
		}
	}
	return
}

func MaxInt8(d []int8) (max int8, index int) {
	for index1, v := range d {
		if v > max {
			max = v
			index = index1
		}
	}
	return
}

func MaxFloat64(d []float64) (max float64, index int) {
	for index1, v := range d {
		if v > max {
			max = v
			index = index1
		}
	}
	return
}

func MaxFloat32(d []float32) (max float32, index int) {
	for index1, v := range d {
		if v > max {
			max = v
			index = index1
		}
	}
	return
}

func MinInt64(d []int64) (min int64, index int) {
	if len(d) != 0 {
		min = d[0]
	}
	for index1, v := range d {
		if v < min {
			min = v
			index = index1
		}
	}
	return
}

func MinInt(d []int) (min int, index int) {
	if len(d) != 0 {
		min = d[0]
	}
	for index1, v := range d {
		if v < min {
			min = v
			index = index1
		}
	}
	return
}

func MinInt32(d []int32) (min int32, index int) {
	if len(d) != 0 {
		min = d[0]
	}
	for index1, v := range d {
		if v < min {
			min = v
			index = index1
		}
	}
	return
}

func MinInt8(d []int8) (min int8, index int) {
	if len(d) != 0 {
		min = d[0]
	}
	for index1, v := range d {
		if v < min {
			min = v
			index = index1
		}
	}
	return
}

func MinFloat64(d []float64) (min float64, index int) {
	if len(d) != 0 {
		min = d[0]
	}
	for index1, v := range d {
		if v < min {
			min = v
			index = index1
		}
	}
	return
}

func MinFloat32(d []float32) (min float32, index int) {
	if len(d) != 0 {
		min = d[0]
	}
	for index1, v := range d {
		if v < min {
			min = v
			index = index1
		}
	}
	return
}

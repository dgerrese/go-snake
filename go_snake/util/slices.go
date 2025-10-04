package util

func IntSliceMin(nums []int) int {
	if len(nums) == 0 {
		panic("empty slice")
	}

	intMin := nums[0]
	for _, num := range nums {
		if num < intMin {
			intMin = num
		}
	}

	return intMin
}

func IntSliceMax(nums []int) int {
	if len(nums) == 0 {
		panic("empty slice")
	}

	intMax := nums[0]
	for _, num := range nums {
		if num > intMax {
			intMax = num
		}
	}

	return intMax
}

func MapSlice[T any, R any](slice []T, mapper func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}

	return result
}

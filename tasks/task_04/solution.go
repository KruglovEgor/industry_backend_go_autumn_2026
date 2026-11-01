package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}
	first := nums[1] - nums[0]
	stats := Stats{Count: len(nums) - 1, Sum: first, Min: first, Max: first}
	for i := 2; i < len(nums); i++ {
		delta := nums[i] - nums[i-1]
		stats.Sum += delta
		stats.Min = min(stats.Min, delta)
		stats.Max = max(stats.Max, delta)
	}
	return stats
}

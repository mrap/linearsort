package linearsort

func Sort(a *[]int) {
	if a == nil || len(*a) == 0 {
		return
	}

	_, max := minMax(a)
	counts := make([]int, max+1)
	for i := 0; i < len(*a); i++ {
		counts[(*a)[i]]++
	}

	for i, insertAt := 0, 0; insertAt < len(*a); i++ {
		for ; counts[i] > 0; counts[i]-- {
			(*a)[insertAt] = i
			insertAt++
		}
	}
}

func minMax(a *[]int) (int, int) {
	min := (*a)[0]
	max := (*a)[0]
	for i := 1; i < len(*a); i++ {
		v := (*a)[i]
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}

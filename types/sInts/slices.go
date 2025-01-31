package sInts

func CompareSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func SumSlice(s []int) int {
	var sum int
	for i := range s {
		sum += s[i]
	}
	return sum
}

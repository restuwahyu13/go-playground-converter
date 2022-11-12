package gpc

func mergeSlice(input ...[]string) (res []string) {
	if len(input) > 0 {
		for i, _ := range input {
			for _, j := range input[i] {
				res = append(res, j)
			}
		}
	} else {
		return res
	}
	return res
}

package gpc

func mergeSlice(input ...[]string) []string {
	if len(input) == 0 {
		return []string{}
	}

	totalLen := 0
	for _, in := range input {
		totalLen += len(in)
	}

	res := make([]string, 0, totalLen)

	for _, in := range input {
		res = append(res, in...)
	}

	return res
}

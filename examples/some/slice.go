package some

func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]struct{}, len(slice1))
	result := make([]string, 0)
	repeated := make(map[string]struct{}, len(slice2))
	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	for _, v := range slice2 {
		_, r := repeated[v]
		if _, ok := m[v]; ok && !r {
			repeated[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Intersect2 最优版本
func Intersect2(slice1, slice2 []string) []string {
	m := make(map[string]struct{}, len(slice1))
	result := make([]string, 0)
	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			delete(m, v)
			result = append(result, v)
		}
	}
	return result
}

func Intersect1(slice1, slice2 []string) []string {
	m := make(map[string]struct{}, len(slice1))
	result := make([]string, 0)
	repeated := make(map[string]struct{}, len(slice2))
	for i := 0; i < len(slice1); i++ {
		m[slice1[i]] = struct{}{}
	}
	for i := 0; i < len(slice2); i++ {
		v := slice2[i]
		_, r := repeated[v]
		if _, ok := m[v]; ok && !r {
			repeated[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

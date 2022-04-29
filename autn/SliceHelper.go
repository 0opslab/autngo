package autn

type SliceHelper struct {
}

//求并集
func (tt *SliceHelper) Union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

//求交集
func (tt *SliceHelper) Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func (tt *SliceHelper) Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := tt.Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

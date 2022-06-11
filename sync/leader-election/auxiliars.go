package les

func getMaxMapIntStruct(u map[int]struct{}) int {
	maxuid := -1
	for k := range u {
		if k > maxuid {
			maxuid = k
		}
	}
	return maxuid
}
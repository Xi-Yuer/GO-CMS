package unique

// RemoveDuplicatesAndEmpty 移除数组中重复元素
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	ret = make([]string, 0)
	for i := 0; i < len(a); i++ {
		repeat := false
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			ret = append(ret, a[i])
		}
	}
	return
}

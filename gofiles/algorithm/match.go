package algorithm

// 滑动窗口精确匹配, target 为待查询字符串
func Match(str, target string) int {
	targetLength := len(target)
	strLength := len(str)
	// 要查询的标题长度超过此标题长度直接返回
	if targetLength > strLength {
		return -1
	}
	for i := 0; i < strLength-targetLength+1; i++ {
		if str[i:i+targetLength] == target {
			return 1
		}
	}

	return -1
}

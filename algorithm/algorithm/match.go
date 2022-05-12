package algorithm

import "strings"

// 滑动窗口模糊匹配, target 为待查询字符串
func Match(str, target string) bool {
	// 要匹配字符串
	targetLength := len(target)
	// 被匹配字符串
	strLength := len(str)
	// 不分大小写搜索
	target = strings.ToLower(target)
	str = strings.ToLower(str)
	// 要查询的标题长度超过此标题长度直接返回
	if targetLength > strLength {
		return false
	}
	for i := 0; i < strLength-targetLength+1; i++ {
		if str[i:i+targetLength] == target {
			return true
		}
	}

	return false
}

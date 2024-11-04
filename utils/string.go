package utils

func IsAnyBlank(strings ...string) bool {
	// 遍历可变参数列表
	for _, arg := range strings {
		if arg == "" {
			return true
		}
	}
	return false
}

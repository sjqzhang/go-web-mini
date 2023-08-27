package util

//转换为驼峰命名
func ToCamelCase(name string) string {
	if name == "" {
		return ""
	}
	var result string
	for i, v := range name {
		if i == 0 {
			result += string(v)
		} else {
			if name[i-1] == '_' {
				result += string(v - 32)
			} else {
				result += string(v)
			}
		}
	}
	return result
}

//转换为下划线命名
func ToUnderlineCase(name string) string {
	if name == "" {
		return ""
	}
	var result string
	for i, v := range name {
		if v >= 65 && v <= 90 {
			if i == 0 {
				result += string(v + 32)
			} else {
				result += "_" + string(v+32)
			}
		} else {
			result += string(v)
		}
	}
	return result
}

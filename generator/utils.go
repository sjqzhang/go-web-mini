package generator

// TransToCamel 字符串：下划线转驼峰
func TransToCamel(s string, firstLowerCase bool) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		// 首字母小写
		if firstLowerCase && i == 0 && d >= 'A' && d <= 'Z' {
			d = d + 32
		}
		data = append(data, d)
	}
	return string(data[:])
}

// TransToUnderline 字符串：驼峰转下划线
func TransToUnderline(s string) string {
	data := make([]byte, 0, len(s)*2)
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' {
			data = append(data, '_')
		}
		data = append(data, d)
	}
	return string(data[:])
}

// GetTypeMap 获取类型转换map
func GetTypeMap() map[string]string {
	return map[string]string{
		"varchar":    "string",
		"char":       "string",
		"text":       "string",
		"longtext":   "string",
		"tinytext":   "string",
		"mediumtext": "string",
		"blob":       "string",
		"longblob":   "string",
		"tinyblob":   "string",
		"mediumblob": "string",
		"enum":       "string",
		"set":        "string",
		"bit":        "string",
		"binary":     "string",
		"varbinary":  "string",
		"json":       "string",
		"date":       "time.Time",
		"int":        "int32",
		"smallint":   "int32",
		"mediumint":  "int32",
		"bigint":     "int64",
		"tinyint":    "int32",
		"datetime":   "time.Time",
		"year":       "time.Time",
		"time":       "time.Time",
		"timestamp":  "time.Time",
		"float":      "float64",
		"double":     "float64",
		"decimal":    "float64",
	}
}

package service

import "go-web-mini/apps/workflow/vo"

func VarJsonToMap(vars []vo.Variables) map[string]interface{} {
	res := make(map[string]interface{}, 0)
	for _, v := range vars {
		res[v.Name] = v.Value
	}
	return res
}

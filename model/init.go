package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

)

type Model struct {
	ID        int64      `gorm:"primarykey" json:"id"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index;softDelete:flag" json:"deleted_at" `
}

type PagerModel struct {
	Total    int64                  `json:"total"`
	List     interface{}            `json:"list"`
	PageNum  int                    `json:"pageNum" form:"pageNum"`
	PageSize int                    `json:"pageSize" form:"pageSize"`
	Extra    map[string]interface{} `json:"extra"`
}

type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull
)

func BuildWhere(obj interface{}) (whereSQL string, valList []interface{}, err error) {
	bsObj, err := json.Marshal(obj)
	if err != nil {
		return "", nil, err
	}
	var where map[string]interface{}
	err = json.Unmarshal(bsObj, &where)
	if err != nil {
		return "", nil, err
	}
	return BuildWhereSelectFromValue(where, false)

}

// sql build where
func BuildWhereSelectFromKey(where map[string]interface{}, or bool) (whereSQL string, valList []interface{}, err error) {
	for k, v := range where {
		//check v is empty
		if v == nil || fmt.Sprintf("%v", v) == "" || fmt.Sprintf("%v", v) == "0" {
			continue
		}
		ks := strings.Split(k, " ")
		if len(ks) > 2 {
			return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		}

		if whereSQL != "" {
			if or {
				whereSQL += " OR "
			} else {
				whereSQL += " AND "
			}

		}
		strings.Join(ks, ",")
		switch len(ks) {
		case 1:
			//fmt.Println(reflect.TypeOf(v))
			switch v := v.(type) {
			case NullType:
				if v == IsNotNull {
					whereSQL += fmt.Sprint(k, " IS NOT NULL")
				} else {
					whereSQL += fmt.Sprint(k, " IS NULL")
				}
			default:
				whereSQL += fmt.Sprint(k, "=?")
				valList = append(valList, v)
			}
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				whereSQL += fmt.Sprint(k, "=?")
				valList = append(valList, v)
			case ">":
				whereSQL += fmt.Sprint(k, ">?")
				valList = append(valList, v)
			case ">=":
				whereSQL += fmt.Sprint(k, ">=?")
				valList = append(valList, v)
			case "<":
				whereSQL += fmt.Sprint(k, "<?")
				valList = append(valList, v)
			case "<=":
				whereSQL += fmt.Sprint(k, "<=?")
				valList = append(valList, v)
			case "!=":
				whereSQL += fmt.Sprint(k, "!=?")
				valList = append(valList, v)
			case "<>":
				whereSQL += fmt.Sprint(k, "!=?")
				valList = append(valList, v)
			case "in":
				whereSQL += fmt.Sprint(k, " in (?)")
				valList = append(valList, v)
			case "like":
				whereSQL += fmt.Sprint(k, " like ?")
				valList = append(valList, v)
			}
		}
	}
	return
}

// sql build where
func BuildWhereSelectFromValue(where map[string]interface{}, or bool) (whereSQL string, valList []interface{}, err error) {
	for k, v := range where {
		if v == nil || fmt.Sprintf("%v", v) == "" || fmt.Sprintf("%v", v) == "0" {
			continue
		}
		if whereSQL != "" {
			if or {
				whereSQL += " OR "
			} else {
				whereSQL += " AND "
			}
		}

		// 获取字段名和操作符
		v, operator, err := parseKey(fmt.Sprintf("%v", v))
		field:=k
		if err != nil {
			return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		}

		switch operator {
		case "=":
			whereSQL += fmt.Sprint(field, "=?")
			valList = append(valList, v)
		case ">":
			whereSQL += fmt.Sprint(field, ">?")
			valList = append(valList, v)
		case ">=":
			whereSQL += fmt.Sprint(field, ">=?")
			valList = append(valList, v)
		case "<":
			whereSQL += fmt.Sprint(field, "<?")
			valList = append(valList, v)
		case "<=":
			whereSQL += fmt.Sprint(field, "<=?")
			valList = append(valList, v)
		case "!=":
			whereSQL += fmt.Sprint(field, "!=?")
			valList = append(valList, v)
		case "<>":
			whereSQL += fmt.Sprint(field, "!=?")
			valList = append(valList, v)
		case "in":
			whereSQL += fmt.Sprint(field, " in (?)")
			valList = append(valList, v)
		case "like":
			whereSQL += fmt.Sprint(field, " like ?")
			valList = append(valList, v)
		default:
			switch v := v.(type) {
			case NullType:
				if v == IsNotNull {
					whereSQL += fmt.Sprint(field, " IS NOT NULL")
				} else {
					whereSQL += fmt.Sprint(field, " IS NULL")
				}
			default:
				whereSQL += fmt.Sprint(field, "=?")
				valList = append(valList, v)
			}
		}
	}
	return
}

// 解析 key，返回字段名和操作符
func parseKey(val string) (value interface{}, operator string, err error) {
	// 获取操作符
	parts := strings.Fields(val)
	if len(parts) == 1 {
		return parts[0], "=", nil
	} else if len(parts) == 2 {
		return parts[1], parts[0], nil
	} else if len(parts) > 2 {
		return strings.Join(parts[1:]," "), parts[0], nil
	}
	return parts[0], parts[1], nil
}

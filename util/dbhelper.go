package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull
)

func GetDB(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return nil
	}
	return ctx.Value("db").(*gorm.DB)
}

func GetAllData(ctx context.Context, md interface{}, result interface{}) error {
	d := GetDB(ctx).Model(md).Find(result)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func GetPartitionData(ctx context.Context, tableName string, result interface{}) error {
	d := GetDB(ctx).Table(tableName).Find(result)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func SearchAllDataByPage(ctx context.Context, md interface{}, result interface{}, page uint32, count uint32, searchData map[string]interface{}) (uint32, error) {
	query := GetDB(ctx).Model(md).Where(searchData)

	var total int64
	d := query.Count(&total)

	if d.Error != nil {
		return 0, d.Error
	}

	d = query.Offset(int((page - 1) * count)).Limit(int(count)).Find(result)
	if d.Error != nil {
		return 0, d.Error
	}

	return uint32(total), nil
}

func SearchAllPartitionDataByPage(ctx context.Context, tableName string, result interface{}, page uint32, count uint32, searchData map[string]interface{}) (uint32, error) {
	query := GetDB(ctx).Table(tableName).Where(searchData)

	var total int64
	d := query.Count(&total)

	if d.Error != nil {
		return 0, d.Error
	}

	d = query.Offset(int((page - 1) * count)).Limit(int(count)).Find(result)
	if d.Error != nil {
		return 0, d.Error
	}

	return uint32(total), nil
}

func SearchAllData(ctx context.Context, md interface{}, result interface{}, searchData map[string]interface{}) error {
	d := GetDB(ctx).Model(md).Where(searchData).Find(result)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func FuzzySearchAllData(ctx context.Context, md interface{}, result interface{}, searchMap map[string]interface{}, orderBy string) error {
	query := GetDB(ctx).Model(md)
	if len(searchMap) > 0 {
		cond, valList, err := WhereBuild(searchMap)
		if err != nil {
			return err
		}
		query.Where(cond, valList...)
	}
	if len(orderBy) > 0 {
		query = query.Order(orderBy) // 如 "mtime desc"
	}
	db := query.Find(result)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func SearchAllPartitionData(ctx context.Context, tableName string, result interface{}, searchData map[string]interface{}) error {
	d := GetDB(ctx).Table(tableName).Where(searchData).Find(result)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func SearchAllPartitionDataWithLimit(ctx context.Context, tableName string, result interface{}, searchData map[string]interface{}) ([]interface{}, error) {
	var allResults []interface{}
	d := GetDB(ctx).Table(tableName).Where(searchData).FindInBatches(result, 10000, func(tx *gorm.DB, batch int) error {
		// 批量处理找到的记录
		s := reflect.ValueOf(result).Elem()
		for i := 0; i < s.Len(); i++ {
			allResults = append(allResults, s.Index(i).Interface())
		}
		return nil
	})
	if d.Error != nil {
		return nil, d.Error
	}
	return allResults, nil
}

func FuzzySearchPage(ctx context.Context, md interface{}, result interface{}, offset int, count int, searchMap map[string]interface{}, orderBy string) (int64, error) {
	var total int64
	query := GetDB(ctx).Model(md)
	if len(orderBy) > 0 {
		query = query.Order(orderBy) // 如 "mtime desc"
	}
	if len(searchMap) > 0 {
		cond, valList, err := WhereBuild(searchMap)
		if err != nil {
			return total, err
		}
		query = query.Where(cond, valList...)
	}
	db := query.Count(&total)
	if db.Error != nil {
		return total, db.Error
	}

	db = query.Offset(offset).Limit(int(count)).Find(result)
	if db.Error != nil {
		return total, db.Error
	}
	return total, nil
}

func GetDataById(ctx context.Context, result interface{}, id uint64) error {
	d := GetDB(ctx).Model(result).Where("id = ?", id).Take(result)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func GetPartitionDataById(ctx context.Context, tableName string, result interface{}, id uint64) error {
	d := GetDB(ctx).Table(tableName).Where("id = ?", id).Take(result)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func BatchCreateDataWithLimit(ctx context.Context, data interface{}, limit int) error {
	d := GetDB(ctx).Model(data).CreateInBatches(data, limit)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func BatchCreatePartitionDataWithLimit(ctx context.Context, tableName string, data interface{}, limit int) error {
	d := GetDB(ctx).Table(tableName).CreateInBatches(data, limit)
	for errors.Is(d.Error, mysql.ErrInvalidConn) {
	}
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func CreateData(ctx context.Context, data interface{}) error {
	d := GetDB(ctx).Model(data).Create(data)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func Create(ctx context.Context, data interface{}) error {
	d := GetDB(ctx).Create(data)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

// 分表需要指定tablename
func CreatePartitionData(ctx context.Context, tableName string, data interface{}) error {
	d := GetDB(ctx).Table(tableName).Create(data)
	for errors.Is(d.Error, mysql.ErrInvalidConn) {
		d = GetDB(ctx).Table(tableName).Create(data)
	}
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func UpdateData(ctx context.Context, data interface{}) error {
	d := GetDB(ctx).Model(data).Save(data)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func UpdatePartitionData(ctx context.Context, tableName string, data interface{}) error {
	d := GetDB(ctx).Table(tableName).Save(data)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func DeleteDataById(ctx context.Context, md interface{}, id uint64) error {
	d := GetDB(ctx).Model(md).Where("id = ?", id).Delete(md)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

//这个方法有问题，如果searchmap为空，则会删除整张表
func DeleteData(ctx context.Context, md interface{}, searchMap map[string]interface{}) error {
	query := GetDB(ctx).Model(md)
	if len(searchMap) == 0 {
		return fmt.Errorf("searchMap is empty")
	}
	if len(searchMap) > 0 {
		cond, valList, err := WhereBuild(searchMap)
		if err != nil {
			return fmt.Errorf("build sql error, %s", err)
		}
		query.Where(cond, valList...)
	}
	db := query.Delete(md)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func DeletePartitionDataById(ctx context.Context, md interface{}, tableName string, id uint64) error {
	d := GetDB(ctx).Table(tableName).Where("id = ?", id).Delete(md)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

// sql build where
func WhereBuild(where map[string]interface{}) (whereSQL string, valList []interface{}, err error) {
	for k, v := range where {
		ks := strings.Split(k, " ")
		if len(ks) > 2 {
			return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		}

		if whereSQL != "" {
			whereSQL += " AND "
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
func WhereBuildSelect(where map[string]interface{}, or bool) (whereSQL string, valList []interface{}, err error) {
	for k, v := range where {
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

func UpdateDataWithMap(ctx context.Context, model interface{}, data map[string]interface{}) error {
	d := GetDB(ctx).Model(model).Updates(data)
	if d.Error != nil {
		return d.Error
	}
	if d.RowsAffected == 0 {
		return fmt.Errorf("update failed,id not found")
	}
	return nil
}

func UpdateDataWithParam(ctx context.Context, model interface{}, params map[string]interface{}, data map[string]interface{}) error {
	d := GetDB(ctx).Model(model).Where(params).Updates(data)
	if d.Error != nil {
		return d.Error
	}
	if d.RowsAffected == 0 {
		return fmt.Errorf("update failed, record not found")
	}
	return nil
}

func UpdateDataWithComplexParam(ctx context.Context, model interface{}, params map[string]interface{}, data map[string]interface{}) error {
	cond, valList, _ := WhereBuild(params)
	d := GetDB(ctx).Model(model).Where(cond, valList...).Updates(data)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func SearchAllDataByPageWithComplex(ctx context.Context, md interface{}, result interface{}, page uint32, count uint32, searchData map[string]interface{}) (uint32, error) {
	cond, valList, _ := WhereBuild(searchData)
	query := GetDB(ctx).Model(md).Where(cond, valList...)
	var total int64
	d := query.Count(&total)
	if d.Error != nil {
		return 0, d.Error
	}
	d = query.Offset(int((page - 1) * count)).Limit(int(count)).Order("id DESC").Find(result)
	if d.Error != nil {
		return 0, d.Error
	}

	return uint32(total), nil
}

func SearchAllDataWithComplex(ctx context.Context, md interface{}, result interface{}, searchData map[string]interface{}, or bool) error {
	cond, valList, _ := WhereBuildSelect(searchData, or)
	d := GetDB(ctx).Model(md).Where(cond, valList...).Order("id DESC").Find(result)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func SearchAllDataWithComplexAndOrder(ctx context.Context, md interface{}, result interface{}, searchData map[string]interface{}, or bool, order string) error {
	cond, valList, _ := WhereBuildSelect(searchData, or)
	d := GetDB(ctx).Model(md).Where(cond, valList...).Order(order).Find(result)
	//d := GetDB(ctx).Model(md).Where(searchData).Find(result)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func DeleteDataByParams(ctx context.Context, md interface{}, query map[string]interface{}) error {
	cond, valList, err := WhereBuild(query)
	if err != nil {
		return err
	}
	d := GetDB(ctx).Model(md).Where(cond, valList...).Delete(md)
	if d.Error != nil {
		return d.Error
	}

	if d.RowsAffected == 0 {
		return fmt.Errorf("delete failed,id not found")
	}
	return nil
}

func DeleteTableDataByParams(ctx context.Context, tableName string, md interface{}, query map[string]interface{}) error {
	cond, valList, err := WhereBuild(query)
	if err != nil {
		return err
	}
	d := GetDB(ctx).Table(tableName).Where(cond, valList...).Delete(md)
	if d.Error != nil {
		return d.Error
	}

	if d.RowsAffected == 0 {
		return fmt.Errorf("delete failed,id not found")
	}
	return nil
}

func FuzzyBatchUpdatesDataWithMap(ctx context.Context, md interface{}, data map[string]interface{}, searchData map[string]interface{}) error {
	query := GetDB(ctx).Model(md)
	if len(searchData) > 0 {
		cond, valList, err := WhereBuild(searchData)
		if err != nil {
			return fmt.Errorf("build sql error, %s", err)
		}
		query.Where(cond, valList...)
	}
	db := query.Updates(data)
	if db.Error != nil {
		return fmt.Errorf("update data error, %s", db.Error)
	}
	return nil
}

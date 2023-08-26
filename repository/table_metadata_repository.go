package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"go-web-mini/config"
	"go-web-mini/global"
	"go-web-mini/model"
	"time"
)

type TableMetadataRepository struct {
}

func (r *TableMetadataRepository) List(ctx context.Context, query *model.TableMetadataQuery) (*model.PagerModel, error) {
	r.Import(ctx)
	db := global.GetDB(ctx)
	var list []*model.TableMetadata
	var obj model.TableMetadata
	copier.CopyWithOption(&obj, &query, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	var total int64
	where, values, _ := model.BuildWhere(obj)
	err := db.Debug().Model(&obj).Where(where, values...).Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return nil, err
	}
	err = db.Model(&obj).Debug().Where(where, values...).Where("deleted_at is null").Offset((query.PageNum - 1) * query.PageSize).Limit(query.PageSize).Find(&list).Error
	if err != nil {
		return nil, err
	}
	var pagerModel model.PagerModel
	pagerModel.List = list
	pagerModel.Total = total
	pagerModel.PageNum = query.PageNum
	pagerModel.PageSize = query.PageSize
	return &pagerModel, err
}

func (r *TableMetadataRepository) Create(ctx context.Context, obj *model.TableMetadata) (*model.TableMetadata, error) {
	db := global.GetDB(ctx)
	return obj, db.Create(obj).Error
}

func (r *TableMetadataRepository) GetById(ctx context.Context, id int64) (*model.TableMetadata, error) {
	db := global.GetDB(ctx)
	var obj model.TableMetadata
	err := db.Model(obj).Where("id=?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, err
}

func (r *TableMetadataRepository) Update(ctx context.Context, obj *model.TableMetadata) (*model.TableMetadata, error) {
	db := global.GetDB(ctx)
	if obj.ID == 0 {
		return nil, fmt.Errorf("id is empty")
	}
	//当column_comment更新时，更新相应的表的comment
	if obj.ColumnComment != "" {
		//ALTER TABLE your_table_name
		//MODIFY COLUMN column_name data_type COMMENT '新的注释';
		sql := fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s %s COMMENT '%s';", obj.TableAlias, obj.ColumnName, obj.ColumnType, obj.ColumnComment)
		fmt.Println(sql)
		db.Exec(sql)

	}
	count := db.Model(obj).Where("id=?", obj.ID).Save(obj).RowsAffected
	if count > 0 {
		return obj, db.First(obj).Error
	}
	return nil, fmt.Errorf("not found")
}

func (r *TableMetadataRepository) Delete(ctx context.Context, ids []int64) (int64, error) {
	db := global.GetDB(ctx)
	//软删除
	return db.Model(model.TableMetadata{}).Where("id in (?)", ids).UpdateColumn("deleted_at", time.Now()).RowsAffected, nil
}

/*
INSERT INTO `table_metadata` (`table_alias`, `column_name`, `column_comment`, `is_nullable`, `data_type`, `character_max_length`)

SELECT

    TABLE_NAME AS 'table_alias',
    COLUMN_NAME AS 'column_name',
    COLUMN_COMMENT AS 'column_comment',
    IS_NULLABLE as 'is_nullable',
    DATA_TYPE as 'data_type',
    CHARACTER_MAXIMUM_LENGTH 'character_max_length'
FROM
    INFORMATION_SCHEMA.COLUMNS
WHERE
    TABLE_SCHEMA = 'go_web_mini'
ORDER BY
    TABLE_NAME,
    ORDINAL_POSITION;
*/

// 利用上面的sql语句，将数据库中的表结构信息导入到表中，忽略已经存在的数据
func (r *TableMetadataRepository) Import(ctx context.Context) error {
	db := global.GetDB(ctx)
	db.Exec("truncate table table_metadata")
	sql := `INSERT IGNORE INTO table_metadata (table_alias, column_name, column_comment,column_type, is_nullable, data_type, character_max_length)
SELECT

    TABLE_NAME AS 'table_alias',
    COLUMN_NAME AS 'column_name',
    COLUMN_COMMENT AS 'column_comment',
 	COLUMN_TYPE as 'column_type',
    IS_NULLABLE as 'is_nullable',
    DATA_TYPE as 'data_type',
    CHARACTER_MAXIMUM_LENGTH 'character_max_length'
FROM
    INFORMATION_SCHEMA.COLUMNS
WHERE
    TABLE_SCHEMA = '%s'
ORDER BY
    TABLE_NAME,
    ORDINAL_POSITION;`
	return db.Exec(fmt.Sprintf(sql, config.Conf.Mysql.Database)).Error
}

package test

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
	_ "gorm.io/plugin/soft_delete"
	"testing"
	"time"
)

type GormTestSimple struct {
	ID         string
	CreateTime *ReconcileLocalDateTime
	UpdateTime *ReconcileLocalDateTime
}

// NameReplacer是用来替换列名还是替换表名的。gorm:"tableName"
func TestUpsetGormTest(t *testing.T) {
	dateTime := ReconcileLocalDateTime(time.Now())
	model := GormTest{
		CreateTime:  &dateTime,
		UpdateTime:  &dateTime,
		GormContent: datatypes.JSON(`{"name": "zhangsan", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`),
	}
	// model里面就包含了自增主键
	Db.Create(&model)
	fmt.Println("开始save")
	model.Content = "111"
	Db.Save(&model)
	fmt.Println("结束save")

	id := model.ID
	// 主键
	fmt.Println(id)
	// Error 1062 (23000): Duplicate entry '33' for key 'gorm_test.PRIMARY'
	Db.Create(&model)

	// 如果想要upset，需要，这句话等于 INSERT INTO `gorm_test` *** ON DUPLICATE KEY UPDATE `id`=xxx;
	// Save方法，Save不能和Model配合使用
	Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"id": id}),
	}).Create(&model)

	simple := &GormTestSimple{}
	// 一张表里，我只想查某个字段
	// SELECT `gorm_test`.`id`,`gorm_test`.`create_time`,`gorm_test`.`update_time` FROM `gorm_test` WHERE `gorm_test`.`deleted_at` IS NULL LIMIT 10
	Db.Model(&GormTest{}).Limit(10).Find(simple)

	fmt.Println(simple)

	// 单会话模式
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	Db.WithContext(ctx).Find(&GormTest{})

}

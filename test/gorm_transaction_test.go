package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/plugin/soft_delete"
	"testing"
	"time"
)

// 事务相关
func TestTransactionGormTest(t *testing.T) {
	model := GormTest{}
	Db.Session(&gorm.Session{SkipDefaultTransaction: true}).First(&model, 1)
	fmt.Println(model)

	now := ReconcileLocalDateTime(time.Now())

	err := Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//  Column 'create_time' cannot be null
		if err := tx.Create(&GormTest{
			Content:    "事务测试1",
			CreateTime: &now,
			UpdateTime: &now,
		}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&GormTest{Content: "事务测试2"}).Error; err != nil {
			return nil
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	// savePoint。 This command is not supported in the prepared statement protocol yet

	// 开启事务
	tx := Db.Begin()

	tx.Create(&GormTest{
		Content:    "事务测试3",
		CreateTime: &now,
		UpdateTime: &now,
	})

}

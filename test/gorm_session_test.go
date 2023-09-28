package test

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	_ "gorm.io/plugin/soft_delete"
	"testing"
	"time"
)

func TestSessionGormTest(t *testing.T) {
	dateTime := ReconcileLocalDateTime(time.Now())
	model := GormTest{
		CreateTime:  &dateTime,
		UpdateTime:  &dateTime,
		GormContent: datatypes.JSON(`{"name": "zhangsan", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`),
	}
	// 新建会话模式
	stmt := Db.Session(&gorm.Session{DryRun: true}).First(&model, 1).Statement
	fmt.Println(stmt.SQL.String())

	// 预编译。执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
	// 会话模式
	tx := Db.Session(&gorm.Session{PrepareStmt: true})
	tx.First(&model, 1)

	// 可以跳过勾子函数
	Db.Session(&gorm.Session{SkipHooks: true}).Create(&model)

	// 设置session的上下文
	timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second)
	Db.Session(&gorm.Session{Context: timeoutCtx})

}

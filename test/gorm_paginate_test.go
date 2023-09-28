package test

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/plugin/soft_delete"
	"testing"
)

// 分页
func TestPaginateGormTest(t *testing.T) {
	var g []GormTest
	// Scopes 可以把类似的逻辑定义到Scopes里
	Db.Scopes(Paginate(1, 10)).Find(&g)
	marshal, _ := json.Marshal(&g)
	fmt.Println(string(marshal))

}

func Paginate(pageIndex int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageIndex <= 0 {
			pageIndex = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (pageIndex - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

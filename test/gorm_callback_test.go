package test

import (
	"fmt"
	"github.com/LinkinStars/go-scaffold/contrib/cryptor"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

var callbackDb *gorm.DB

func init() {
	callbackDb = initDB()

}

type CallbackTest struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
	ApplyTime time.Time
}

func (u *CallbackTest) BeforeSave(db *gorm.DB) (err error) {
	u.ApplyTime = db.NowFunc()
	// 在保存之前对密码字段进行加密
	encryptedContent := cryptor.AesSimpleEncrypt(u.Content, "1234")
	u.Content = encryptedContent
	return nil
}

// 钩子函数，在从数据库检索数据之后进行处理
func (u *CallbackTest) AfterFind(*gorm.DB) error {
	// 在检索之后对密码字段进行解密
	decryptedContent := cryptor.AesSimpleDecrypt(u.Content, "1234")
	u.Content = decryptedContent
	return nil
}

func TestSaveCallbackTest(t *testing.T) {

	// [error] Got error when compile callbacks, got conflicting callback gorm:test_save_case with before gorm:save_before_associations
	err := callbackDb.Callback().Create().Before("gorm:save_before_associations").After("gorm:before_create").Register("gorm:test_save_case", func(tx *gorm.DB) {
		fmt.Println("我注册了个钩子函数在上面")
	})

	if err != nil {
		return
	}

	model := CallbackTest{
		Content: "callback测试",
	}

	result := Db.Create(&model)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	fmt.Println(model)
	fmt.Println(model.Content)

	queryResult := CallbackTest{}
	Db.Table("callback_test").Where("id = ?", "29").Find(&queryResult)
	fmt.Println(queryResult.Content)

}

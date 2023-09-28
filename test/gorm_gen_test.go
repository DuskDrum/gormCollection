package test

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {
	var f1 float32 = 16777216.0
	var f2 float32 = 16777217.0
	fmt.Println(f1 == f2)
}

func TestNil(t *testing.T) {
	//var str1 string = nil
	//fmt.Println(str1)

	var str2 *string = nil
	fmt.Println(str2)
}

//
//import (
//	"database/sql/driver"
//	"errors"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gormCollection/dal/gen"
//	"log"
//	"time"
//)
//
//var Db *gorm.DB
//
//type ReconcileLocalDateTime time.Time
//
//func init() {
//	initDB()
//	//Db.Callback().Create().Before("gorm:create").After("gorm:update_time_stamp").Register("gorm:reconcile_local_date_time", updateReconcileLocalDateTimeForCreateCallback)
//
//}
//
//type GormTest struct {
//	gorm.Model
//	Content    string `gorm:"default:'gorm测试'"`
//	CreateTime ReconcileLocalDateTime
//	UpdateTime ReconcileLocalDateTime
//}
//
//
//func (s *GormTest) BeforeSave() {
//	log.Print("开始创建前")
//}
//
//
//func (rlt ReconcileLocalDateTime) Value() (driver.Value, error) {
//	tTime := time.Time(rlt)
//	return tTime.Format("2006-01-02 15:04:05"), nil
//}
//
//func (rlt *ReconcileLocalDateTime) Scan(value interface{}) error {
//	switch vt := value.(type) {
//	case []byte:
//		tTime, _ := time.Parse("2006-01-02 15:04:05", string(vt))
//		*rlt = ReconcileLocalDateTime(tTime)
//	case string:
//		tTime, _ := time.Parse("2006-01-02 15:04:05", vt)
//		*rlt = ReconcileLocalDateTime(tTime)
//	default:
//		fmt.Printf("%#v\n", vt)
//		return errors.New("类型处理错误")
//	}
//	return nil
//}
//
//func initDB() {
//	fmt.Println("数据库初始化")
//	var err error
//
//	host := "127.0.0.1"
//	user := "root"
//	password := "pass"
//	dbName := "gorm"
//
//	// 初始化DB连接
//	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
//		user, password, host, dbName)))
//
//	if err != nil {
//		log.Fatalf("connect error : %v\n", err)
//	}
//
//	gen.SetDefault(db)
//}

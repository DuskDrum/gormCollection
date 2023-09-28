package test

import (
	"database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/hints"
	_ "gorm.io/plugin/soft_delete"
	"log"
	"strings"
	"testing"
	"time"
)

var Db *gorm.DB
var datetimePrecision = 2

type ReconcileLocalDateTime time.Time

func init() {
	Db = initDB()
	//Db.Callback().Create().Before("gorm:create").After("gorm:update_time_stamp").Register("gorm:reconcile_local_date_time", updateReconcileLocalDateTimeForCreateCallback)

}

type GormTest struct {
	gorm.Model
	Content     string `gorm:"default:'gorm测试'"`
	CreateTime  *ReconcileLocalDateTime
	UpdateTime  *ReconcileLocalDateTime
	GormContent datatypes.JSON
	//Deleted     soft_delete.DeletedAt `gorm:"column:deleted;type:int unsigned;softDelete:flag" json:"deleted"` // 是否删除（软删除字段）

}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in the RFC 3339 format with sub-second precision.
// If the timestamp cannot be represented as valid RFC 3339
// (e.g., the year is out of range), then an error is reported.
func (rlt *ReconcileLocalDateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(*rlt)
	return t.MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time must be a quoted string in the RFC 3339 format.
func (rlt *ReconcileLocalDateTime) UnmarshalJSON(data []byte) error {
	t := time.Time(*rlt)
	return t.UnmarshalJSON(data)
}

// Field
func (s *GormTest) TableName() string {
	return "gorm_test"
}

// NameReplacer是用来替换列名还是替换表名的。gorm:"tableName"
func TestSaveGormTest(t *testing.T) {
	dateTime := ReconcileLocalDateTime(time.Now())
	model := GormTest{
		CreateTime:  &dateTime,
		UpdateTime:  &dateTime,
		GormContent: datatypes.JSON(`{"name": "zhangsan", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`),
	}

	result := Db.Create(&model)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	fmt.Println(model)
	fmt.Println(model.Content)

	queryResult := GormTest{}
	Db.Find(&queryResult).Where("id = ?", "1")
	fmt.Println(queryResult.CreateTime.Value())
	fmt.Println(queryResult.UpdateTime.Value())

	jsonQueryResult := GormTest{}
	// Error 3141 (22032): Invalid JSON text in argument 1 to function json_extract: "The document is empty." at position 0.
	Db.Find(&jsonQueryResult, datatypes.JSONQuery("gorm_content").HasKey("name"))
	fmt.Println(jsonQueryResult.GormContent.Value())

	// 强制使用索引
	// Key xxx doesn't exist in table 'gorm_test'
	var testList []GormTest
	Db.Clauses(hints.ForceIndex("idx_gorm_test")).Find(&testList)
	fmt.Println(testList)

}

// 需要实现 driver.Valuer, 和 sql.Scanner 两个接口
func (rlt *ReconcileLocalDateTime) Value() (driver.Value, error) {
	if rlt == nil {
		return nil, nil
	}
	return time.Time(*rlt).Format(time.DateTime), nil
}

func (rlt *ReconcileLocalDateTime) Scan(value interface{}) error {
	if value == nil {
		*rlt = ReconcileLocalDateTime(time.Time{})
		return nil
	}

	switch vt := value.(type) {
	case []byte:
		tTime, _ := time.Parse(time.DateTime, string(vt))
		*rlt = ReconcileLocalDateTime(tTime)
	case string:
		tTime, _ := time.Parse(time.DateTime, vt)
		*rlt = ReconcileLocalDateTime(tTime)
	case time.Time:
		*rlt = ReconcileLocalDateTime(vt)
	}

	return nil
}

func initDB() *gorm.DB {
	fmt.Println("数据库初始化")
	var err error

	host := "127.0.0.1"
	user := "root"
	password := "pass"
	dbName := "gorm"

	// 初始化DB连接
	Db, err := gorm.Open(mysql.New(mysql.Config{
		// MySQL 数据库的连接字符串.https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, dbName),
		// 默认的字符串长度，如果struct里没有定义字段，会取这个值为默认的
		DefaultStringSize: 256,
		// 是否禁用时间精度。默认为 false。设置为 true 时，GORM 将禁用 MySQL 中的时间精度. MySQL 5.6前不支持
		DisableDatetimePrecision: false,
		// 默认的日期时间精度。默认为 0，表示没有指定精度。用于指定数据库中日期时间类型字段的默认精度。
		// 如果设置的是2，百分之一秒，代表时间格式如下：2023-06-01 12:34:56.12
		DefaultDatetimePrecision: &datetimePrecision,
		// 是否禁用重命名索引。默认为 false，表示支持重命名索引。设置为 true 时，GORM 将禁用 MySQL 中的重命名索引功能。mysql5.7以后才支持
		DontSupportRenameIndex: true,
		// 是否禁用重命名列。默认为 false，表示支持重命名列。设置为 true 时，GORM 将禁用 MySQL 中的重命名列功能
		DontSupportRenameColumn: true,
		// 默认为 false。设置为 true 时，GORM 将跳过与数据库版本相关的初始化步骤。自动根据版本做一些初始化
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
			// 表的前缀
			TablePrefix: "",
			// 定义名称替换规则，映射结构和数据库之间
			NameReplacer: strings.NewReplacer("old", "new"),
			// 不禁用字段和表名的小写转换功能
			NoLowerCase: false,
		},
		// 设置为info就有sql了
		Logger: logger.Default.LogMode(logger.Info),
		// 定义now方法
		NowFunc: func() time.Time { return time.Now().Local() },
		//是否跳过默认事务,设置为 true 时，GORM 将跳过默认的事务包装，您需要手动控制事务的开始和提交/回滚
		SkipDefaultTransaction: false,
		// 是否开启预编译语句。设置为 true 时，GORM 将会预编译 SQL 语句以提高执行效率
		PrepareStmt: true,
		// 是否禁用自动 Ping 数据库连接。true 时，GORM 将不会自动 Ping 数据库连接以保持连接活跃
		DisableAutomaticPing: false,
		// 只输出查询语句而不会对数据库进行实际操作
		DryRun: false,
		// 是否完全保存关联对象。设置为 true 时，GORM 将会递归保存关联对象，包括它们的关联对象
		FullSaveAssociations: true,
	})

	if err != nil {
		log.Fatalf("connect error : %v\n", err)
	}

	return Db
}

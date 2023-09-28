package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gormCollection/dal/gen"
	"log"
	"strings"
	"time"
)

var datetimePrecision = 3

func Init() {
	// 加载配置
	host := "127.0.0.1:3306"
	user := "root"
	password := "pass"
	dbName := "gorm"

	// 初始化DB连接
	//db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//	user, password, host, dbName)))

	db, err := gorm.Open(mysql.New(mysql.Config{
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
			// 使用单数表名
			SingularTable: true,
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

	gen.SetDefault(db)
}

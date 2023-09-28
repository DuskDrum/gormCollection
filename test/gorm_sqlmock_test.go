package test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mock   sqlmock.Sqlmock
	err    error
	sqlDB  *sql.DB
	gormDB *gorm.DB
)

type User struct {
	ID   uint
	Name string
}

// TestMain是在当前package下，最先运行的一个函数，常用于初始化
func TestMain(m *testing.M) {
	initDb()
	// m.Run 是调用包下面各个Test函数的入口，它会运行所有的测试用例，并返回一个整数类型的退出码（exit code），0是正常的
	os.Exit(m.Run())
}

func initDb() {

	sqlDB, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      sqlDB,
	}), &gorm.Config{})

	if nil != err {
		initDbErr := fmt.Sprintf("Init DB with sqlmock failed, err %v", err)
		panic(initDbErr)
	}
}

func TestSqlMock(t *testing.T) {
	// 设置 Mock 的期望行为和返回结果
	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "John")
	mock.ExpectQuery("SELECT * FROM `users`").WillReturnRows(rows)

	// 执行被测试的代码，其中涉及数据库操作
	var users []User
	result := gormDB.Find(&users)
	if result.Error != nil {
		log.Fatalf("Failed to fetch users: %s", result.Error)
	}

	// 输出结果
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	}

	// 断言 Mock 的期望行为是否满足
	if err := mock.ExpectationsWereMet(); err != nil {
		log.Fatalf("Unfulfilled expectations: %s", err)
	}
}

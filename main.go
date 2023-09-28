package main

import (
	"gormCollection/initializer/mysql"
	"log"
	"net/http"
)

func main() {
	// 初始化数据库
	mysql.Init()
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// 注意: 请使用package运行模式, 而非单个文件运行模式, 否则会提示找不到func定义
// 参考: https://stackoverflow.com/questions/28153203/undefined-function-declared-in-another-file
func main() {
	// 加载配置
	host := "127.0.0.1:3306"
	user := "root"
	password := "pass"
	dbName := "gorm"

	// 初始化DB连接
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName)))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	// 初始化代码生成器实例
	g := gen.NewGenerator(gen.Config{
		OutPath:           "dal/gen", // 生成的DAL类文件系统路径
		ModelPkgPath:      "model",   // DataObject模型包名
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,  // generate pointer when field is nullable
		FieldCoverable:    false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values
		FieldSignable:     false, // detect integer field's unsigned type, adjust generated data type
		FieldWithIndexTag: false, // generate with gorm index tag
		FieldWithTypeTag:  true,  // generate with gorm column type tag
	})
	g.UseDB(db)

	// 通用配置
	commonOpts := commonFieldOpts()

	// 配置业务数据表
	models := addBizTable(g, commonOpts)
	g.ApplyBasic(models...)

	// 生成代码
	g.Execute()
}

func commonFieldOpts() []gen.ModelOpt {
	// 自动更新时间戳
	autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;autoUpdateTime")
	autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")

	// soft deleted: `deleted`
	softDeleteFieldTag := gen.FieldGORMTag("deleted", "column:deleted;type:int unsigned;softDelete:flag")
	softDeleteFieldType := gen.FieldType("deleted", "soft_delete.DeletedAt")

	return []gen.ModelOpt{
		autoUpdateTimeField,
		autoCreateTimeField,
		softDeleteFieldTag,
		softDeleteFieldType,
	}
}

func addBizTable(g *gen.Generator, commonOpts []gen.ModelOpt) []interface{} {
	models := make([]interface{}, 32)

	// 我方流水
	clearingRecord := g.GenerateModelAs("clearing_record", "ClearingRecord",
		append(
			commonOpts,
			// 在下面添加其他配置
		)...,
	)
	// 这里添加自定义SQL
	g.ApplyInterface(func(clearingRecordMapper) {}, clearingRecord)
	models = append(models, clearingRecord)

	// 渠道流水
	channelClearingRecord := g.GenerateModelAs("channel_clearing_record", "ChannelClearingRecord",
		append(
			commonOpts,
		)...,
	)
	models = append(models, channelClearingRecord)

	// 差异流水
	clearingDiffRecord := g.GenerateModelAs("clearing_diff_record", "ClearingDiffRecord",
		append(
			commonOpts,
		)...,
	)
	models = append(models, clearingDiffRecord)

	return models
}

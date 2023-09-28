// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gormCollection/dal/model"
)

func newClearingRecord(db *gorm.DB, opts ...gen.DOOption) clearingRecord {
	_clearingRecord := clearingRecord{}

	_clearingRecord.clearingRecordDo.UseDB(db, opts...)
	_clearingRecord.clearingRecordDo.UseModel(&model.ClearingRecord{})

	tableName := _clearingRecord.clearingRecordDo.TableName()
	_clearingRecord.ALL = field.NewAsterisk(tableName)
	_clearingRecord.ClearingRecordID = field.NewString(tableName, "clearing_record_id")
	_clearingRecord.ChannelReferenceID = field.NewString(tableName, "channel_reference_id")
	_clearingRecord.TransactionOrderID = field.NewString(tableName, "transaction_order_id")
	_clearingRecord.FundsReconciliationPipeID = field.NewString(tableName, "funds_reconciliation_pipe_id")
	_clearingRecord.ClearingReconciliationPipeID = field.NewString(tableName, "clearing_reconciliation_pipe_id")
	_clearingRecord.RecordType = field.NewString(tableName, "record_type")
	_clearingRecord.TransactionType = field.NewString(tableName, "transaction_type")
	_clearingRecord.OrderStatus = field.NewString(tableName, "order_status")
	_clearingRecord.Amount = field.NewString(tableName, "amount")
	_clearingRecord.Currency = field.NewString(tableName, "currency")
	_clearingRecord.OrderTime = field.NewTime(tableName, "order_time")
	_clearingRecord.ClearingReconciliationStatus = field.NewString(tableName, "clearing_reconciliation_status")
	_clearingRecord.ClearingReconciliationResult = field.NewString(tableName, "clearing_reconciliation_result")
	_clearingRecord.ClearingTime = field.NewTime(tableName, "clearing_time")
	_clearingRecord.StatementID = field.NewString(tableName, "statement_id")
	_clearingRecord.FundsReconciliationStatus = field.NewString(tableName, "funds_reconciliation_status")
	_clearingRecord.FundsReconciliationBatch = field.NewString(tableName, "funds_reconciliation_batch")
	_clearingRecord.Age = field.NewInt32(tableName, "age")
	_clearingRecord.AgeLevel = field.NewString(tableName, "age_level")
	_clearingRecord.AgeStatus = field.NewString(tableName, "age_status")
	_clearingRecord.ChannelID = field.NewString(tableName, "channel_id")
	_clearingRecord.ChannelClearingBatchNo = field.NewString(tableName, "channel_clearing_batch_no")
	_clearingRecord.ChannelClearingTime = field.NewString(tableName, "channel_clearing_time")
	_clearingRecord.ChannelDealID = field.NewString(tableName, "channel_deal_id")
	_clearingRecord.ChannelSettlementCurrency = field.NewString(tableName, "channel_settlement_currency")
	_clearingRecord.RecordStatus = field.NewString(tableName, "record_status")
	_clearingRecord.EstimateClearingTime = field.NewString(tableName, "estimate_clearing_time")
	_clearingRecord.CreateTime = field.NewTime(tableName, "create_time")
	_clearingRecord.UpdateTime = field.NewTime(tableName, "update_time")
	_clearingRecord.ClearingCycle = field.NewString(tableName, "clearing_cycle")
	_clearingRecord.ChannelClearingRecordID = field.NewString(tableName, "channel_clearing_record_id")
	_clearingRecord.OriginTransactionOrderID = field.NewString(tableName, "origin_transaction_order_id")
	_clearingRecord.ClearingReconcileDetailResultID = field.NewString(tableName, "clearing_reconcile_detail_result_id")
	_clearingRecord.PayTransactionOrderID = field.NewString(tableName, "pay_transaction_order_id")
	_clearingRecord.RefundTransactionOrderID = field.NewString(tableName, "refund_transaction_order_id")
	_clearingRecord.MerchantID = field.NewString(tableName, "merchant_id")
	_clearingRecord.Region = field.NewString(tableName, "region")
	_clearingRecord.PayMethod = field.NewString(tableName, "pay_method")
	_clearingRecord.Extension = field.NewString(tableName, "extension")
	_clearingRecord.SettleFinishTime = field.NewString(tableName, "settle_finish_time")
	_clearingRecord.BillID = field.NewString(tableName, "bill_id")

	_clearingRecord.fillFieldMap()

	return _clearingRecord
}

type clearingRecord struct {
	clearingRecordDo clearingRecordDo

	ALL                             field.Asterisk
	ClearingRecordID                field.String // 主键ID
	ChannelReferenceID              field.String // 我方流水号（我方送渠道订单号）
	TransactionOrderID              field.String // 代表我方的支付订单主键
	FundsReconciliationPipeID       field.String // 资金对账通道
	ClearingReconciliationPipeID    field.String // 清算对账通道
	RecordType                      field.String // 单据类型：收单/外汇
	TransactionType                 field.String // 支付类型：payment/refund/cb/cb_win/cb_lose
	OrderStatus                     field.String // 订单状态：success/fail
	Amount                          field.String // 我方清算金额
	Currency                        field.String // 我方清算币种
	OrderTime                       field.Time   // 我方支付时间
	ClearingReconciliationStatus    field.String // 清算对账状态: 已对账/未对账
	ClearingReconciliationResult    field.String // 清算对账结果
	ClearingTime                    field.Time   // 清算对账时间
	StatementID                     field.String // 我方账单编号
	FundsReconciliationStatus       field.String // 资金对账状态：已对账/未对账
	FundsReconciliationBatch        field.String // 资金对账批次号
	Age                             field.Int32  // 账龄
	AgeLevel                        field.String // 账龄等级
	AgeStatus                       field.String // 账龄处理状态
	ChannelID                       field.String // 渠道
	ChannelClearingBatchNo          field.String // 渠道清算批次
	ChannelClearingTime             field.String // 渠道清算时间
	ChannelDealID                   field.String // 渠道流水号（渠道订单号）
	ChannelSettlementCurrency       field.String // 结算币种
	RecordStatus                    field.String // 凭证处理状态: 正常/异常
	EstimateClearingTime            field.String // 预计清算时间
	CreateTime                      field.Time   // 创建时间
	UpdateTime                      field.Time   // 更新时间
	ClearingCycle                   field.String // 清算周期
	ChannelClearingRecordID         field.String // 标准渠道流水id
	OriginTransactionOrderID        field.String // 退款原支付单id
	ClearingReconcileDetailResultID field.String // 清算对账结果id
	PayTransactionOrderID           field.String // 交易层支付Id
	RefundTransactionOrderID        field.String // 交易层退款Id
	MerchantID                      field.String // 商户id
	Region                          field.String // 国家/地区
	PayMethod                       field.String // 支付方式
	Extension                       field.String // 拓展字段
	SettleFinishTime                field.String // 结账时间
	BillID                          field.String // POS结账退款对账id

	fieldMap map[string]field.Expr
}

func (c clearingRecord) Table(newTableName string) *clearingRecord {
	c.clearingRecordDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c clearingRecord) As(alias string) *clearingRecord {
	c.clearingRecordDo.DO = *(c.clearingRecordDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *clearingRecord) updateTableName(table string) *clearingRecord {
	c.ALL = field.NewAsterisk(table)
	c.ClearingRecordID = field.NewString(table, "clearing_record_id")
	c.ChannelReferenceID = field.NewString(table, "channel_reference_id")
	c.TransactionOrderID = field.NewString(table, "transaction_order_id")
	c.FundsReconciliationPipeID = field.NewString(table, "funds_reconciliation_pipe_id")
	c.ClearingReconciliationPipeID = field.NewString(table, "clearing_reconciliation_pipe_id")
	c.RecordType = field.NewString(table, "record_type")
	c.TransactionType = field.NewString(table, "transaction_type")
	c.OrderStatus = field.NewString(table, "order_status")
	c.Amount = field.NewString(table, "amount")
	c.Currency = field.NewString(table, "currency")
	c.OrderTime = field.NewTime(table, "order_time")
	c.ClearingReconciliationStatus = field.NewString(table, "clearing_reconciliation_status")
	c.ClearingReconciliationResult = field.NewString(table, "clearing_reconciliation_result")
	c.ClearingTime = field.NewTime(table, "clearing_time")
	c.StatementID = field.NewString(table, "statement_id")
	c.FundsReconciliationStatus = field.NewString(table, "funds_reconciliation_status")
	c.FundsReconciliationBatch = field.NewString(table, "funds_reconciliation_batch")
	c.Age = field.NewInt32(table, "age")
	c.AgeLevel = field.NewString(table, "age_level")
	c.AgeStatus = field.NewString(table, "age_status")
	c.ChannelID = field.NewString(table, "channel_id")
	c.ChannelClearingBatchNo = field.NewString(table, "channel_clearing_batch_no")
	c.ChannelClearingTime = field.NewString(table, "channel_clearing_time")
	c.ChannelDealID = field.NewString(table, "channel_deal_id")
	c.ChannelSettlementCurrency = field.NewString(table, "channel_settlement_currency")
	c.RecordStatus = field.NewString(table, "record_status")
	c.EstimateClearingTime = field.NewString(table, "estimate_clearing_time")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateTime = field.NewTime(table, "update_time")
	c.ClearingCycle = field.NewString(table, "clearing_cycle")
	c.ChannelClearingRecordID = field.NewString(table, "channel_clearing_record_id")
	c.OriginTransactionOrderID = field.NewString(table, "origin_transaction_order_id")
	c.ClearingReconcileDetailResultID = field.NewString(table, "clearing_reconcile_detail_result_id")
	c.PayTransactionOrderID = field.NewString(table, "pay_transaction_order_id")
	c.RefundTransactionOrderID = field.NewString(table, "refund_transaction_order_id")
	c.MerchantID = field.NewString(table, "merchant_id")
	c.Region = field.NewString(table, "region")
	c.PayMethod = field.NewString(table, "pay_method")
	c.Extension = field.NewString(table, "extension")
	c.SettleFinishTime = field.NewString(table, "settle_finish_time")
	c.BillID = field.NewString(table, "bill_id")

	c.fillFieldMap()

	return c
}

func (c *clearingRecord) WithContext(ctx context.Context) IClearingRecordDo {
	return c.clearingRecordDo.WithContext(ctx)
}

func (c clearingRecord) TableName() string { return c.clearingRecordDo.TableName() }

func (c clearingRecord) Alias() string { return c.clearingRecordDo.Alias() }

func (c *clearingRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *clearingRecord) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 41)
	c.fieldMap["clearing_record_id"] = c.ClearingRecordID
	c.fieldMap["channel_reference_id"] = c.ChannelReferenceID
	c.fieldMap["transaction_order_id"] = c.TransactionOrderID
	c.fieldMap["funds_reconciliation_pipe_id"] = c.FundsReconciliationPipeID
	c.fieldMap["clearing_reconciliation_pipe_id"] = c.ClearingReconciliationPipeID
	c.fieldMap["record_type"] = c.RecordType
	c.fieldMap["transaction_type"] = c.TransactionType
	c.fieldMap["order_status"] = c.OrderStatus
	c.fieldMap["amount"] = c.Amount
	c.fieldMap["currency"] = c.Currency
	c.fieldMap["order_time"] = c.OrderTime
	c.fieldMap["clearing_reconciliation_status"] = c.ClearingReconciliationStatus
	c.fieldMap["clearing_reconciliation_result"] = c.ClearingReconciliationResult
	c.fieldMap["clearing_time"] = c.ClearingTime
	c.fieldMap["statement_id"] = c.StatementID
	c.fieldMap["funds_reconciliation_status"] = c.FundsReconciliationStatus
	c.fieldMap["funds_reconciliation_batch"] = c.FundsReconciliationBatch
	c.fieldMap["age"] = c.Age
	c.fieldMap["age_level"] = c.AgeLevel
	c.fieldMap["age_status"] = c.AgeStatus
	c.fieldMap["channel_id"] = c.ChannelID
	c.fieldMap["channel_clearing_batch_no"] = c.ChannelClearingBatchNo
	c.fieldMap["channel_clearing_time"] = c.ChannelClearingTime
	c.fieldMap["channel_deal_id"] = c.ChannelDealID
	c.fieldMap["channel_settlement_currency"] = c.ChannelSettlementCurrency
	c.fieldMap["record_status"] = c.RecordStatus
	c.fieldMap["estimate_clearing_time"] = c.EstimateClearingTime
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_time"] = c.UpdateTime
	c.fieldMap["clearing_cycle"] = c.ClearingCycle
	c.fieldMap["channel_clearing_record_id"] = c.ChannelClearingRecordID
	c.fieldMap["origin_transaction_order_id"] = c.OriginTransactionOrderID
	c.fieldMap["clearing_reconcile_detail_result_id"] = c.ClearingReconcileDetailResultID
	c.fieldMap["pay_transaction_order_id"] = c.PayTransactionOrderID
	c.fieldMap["refund_transaction_order_id"] = c.RefundTransactionOrderID
	c.fieldMap["merchant_id"] = c.MerchantID
	c.fieldMap["region"] = c.Region
	c.fieldMap["pay_method"] = c.PayMethod
	c.fieldMap["extension"] = c.Extension
	c.fieldMap["settle_finish_time"] = c.SettleFinishTime
	c.fieldMap["bill_id"] = c.BillID
}

func (c clearingRecord) clone(db *gorm.DB) clearingRecord {
	c.clearingRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c clearingRecord) replaceDB(db *gorm.DB) clearingRecord {
	c.clearingRecordDo.ReplaceDB(db)
	return c
}

type clearingRecordDo struct{ gen.DO }

type IClearingRecordDo interface {
	gen.SubQuery
	Debug() IClearingRecordDo
	WithContext(ctx context.Context) IClearingRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IClearingRecordDo
	WriteDB() IClearingRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IClearingRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IClearingRecordDo
	Not(conds ...gen.Condition) IClearingRecordDo
	Or(conds ...gen.Condition) IClearingRecordDo
	Select(conds ...field.Expr) IClearingRecordDo
	Where(conds ...gen.Condition) IClearingRecordDo
	Order(conds ...field.Expr) IClearingRecordDo
	Distinct(cols ...field.Expr) IClearingRecordDo
	Omit(cols ...field.Expr) IClearingRecordDo
	Join(table schema.Tabler, on ...field.Expr) IClearingRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IClearingRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IClearingRecordDo
	Group(cols ...field.Expr) IClearingRecordDo
	Having(conds ...gen.Condition) IClearingRecordDo
	Limit(limit int) IClearingRecordDo
	Offset(offset int) IClearingRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IClearingRecordDo
	Unscoped() IClearingRecordDo
	Create(values ...*model.ClearingRecord) error
	CreateInBatches(values []*model.ClearingRecord, batchSize int) error
	Save(values ...*model.ClearingRecord) error
	First() (*model.ClearingRecord, error)
	Take() (*model.ClearingRecord, error)
	Last() (*model.ClearingRecord, error)
	Find() ([]*model.ClearingRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ClearingRecord, err error)
	FindInBatches(result *[]*model.ClearingRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ClearingRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IClearingRecordDo
	Assign(attrs ...field.AssignExpr) IClearingRecordDo
	Joins(fields ...field.RelationField) IClearingRecordDo
	Preload(fields ...field.RelationField) IClearingRecordDo
	FirstOrInit() (*model.ClearingRecord, error)
	FirstOrCreate() (*model.ClearingRecord, error)
	FindByPage(offset int, limit int) (result []*model.ClearingRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IClearingRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByName(channelReferenceId string) (result []*model.ClearingRecord, err error)
}

// SELECT * FROM @@table WHERE channelReferenceId = @channelReferenceId
func (c clearingRecordDo) GetByName(channelReferenceId string) (result []*model.ClearingRecord, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, channelReferenceId)
	generateSQL.WriteString("SELECT * FROM clearing_record WHERE channelReferenceId = ? ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c clearingRecordDo) Debug() IClearingRecordDo {
	return c.withDO(c.DO.Debug())
}

func (c clearingRecordDo) WithContext(ctx context.Context) IClearingRecordDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c clearingRecordDo) ReadDB() IClearingRecordDo {
	return c.Clauses(dbresolver.Read)
}

func (c clearingRecordDo) WriteDB() IClearingRecordDo {
	return c.Clauses(dbresolver.Write)
}

func (c clearingRecordDo) Session(config *gorm.Session) IClearingRecordDo {
	return c.withDO(c.DO.Session(config))
}

func (c clearingRecordDo) Clauses(conds ...clause.Expression) IClearingRecordDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c clearingRecordDo) Returning(value interface{}, columns ...string) IClearingRecordDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c clearingRecordDo) Not(conds ...gen.Condition) IClearingRecordDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c clearingRecordDo) Or(conds ...gen.Condition) IClearingRecordDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c clearingRecordDo) Select(conds ...field.Expr) IClearingRecordDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c clearingRecordDo) Where(conds ...gen.Condition) IClearingRecordDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c clearingRecordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IClearingRecordDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c clearingRecordDo) Order(conds ...field.Expr) IClearingRecordDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c clearingRecordDo) Distinct(cols ...field.Expr) IClearingRecordDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c clearingRecordDo) Omit(cols ...field.Expr) IClearingRecordDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c clearingRecordDo) Join(table schema.Tabler, on ...field.Expr) IClearingRecordDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c clearingRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IClearingRecordDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c clearingRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IClearingRecordDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c clearingRecordDo) Group(cols ...field.Expr) IClearingRecordDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c clearingRecordDo) Having(conds ...gen.Condition) IClearingRecordDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c clearingRecordDo) Limit(limit int) IClearingRecordDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c clearingRecordDo) Offset(offset int) IClearingRecordDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c clearingRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IClearingRecordDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c clearingRecordDo) Unscoped() IClearingRecordDo {
	return c.withDO(c.DO.Unscoped())
}

func (c clearingRecordDo) Create(values ...*model.ClearingRecord) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c clearingRecordDo) CreateInBatches(values []*model.ClearingRecord, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c clearingRecordDo) Save(values ...*model.ClearingRecord) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c clearingRecordDo) First() (*model.ClearingRecord, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClearingRecord), nil
	}
}

func (c clearingRecordDo) Take() (*model.ClearingRecord, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClearingRecord), nil
	}
}

func (c clearingRecordDo) Last() (*model.ClearingRecord, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClearingRecord), nil
	}
}

func (c clearingRecordDo) Find() ([]*model.ClearingRecord, error) {
	result, err := c.DO.Find()
	return result.([]*model.ClearingRecord), err
}

func (c clearingRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ClearingRecord, err error) {
	buf := make([]*model.ClearingRecord, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c clearingRecordDo) FindInBatches(result *[]*model.ClearingRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c clearingRecordDo) Attrs(attrs ...field.AssignExpr) IClearingRecordDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c clearingRecordDo) Assign(attrs ...field.AssignExpr) IClearingRecordDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c clearingRecordDo) Joins(fields ...field.RelationField) IClearingRecordDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c clearingRecordDo) Preload(fields ...field.RelationField) IClearingRecordDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c clearingRecordDo) FirstOrInit() (*model.ClearingRecord, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClearingRecord), nil
	}
}

func (c clearingRecordDo) FirstOrCreate() (*model.ClearingRecord, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClearingRecord), nil
	}
}

func (c clearingRecordDo) FindByPage(offset int, limit int) (result []*model.ClearingRecord, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c clearingRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c clearingRecordDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c clearingRecordDo) Delete(models ...*model.ClearingRecord) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *clearingRecordDo) withDO(do gen.Dao) *clearingRecordDo {
	c.DO = *do.(*gen.DO)
	return c
}
